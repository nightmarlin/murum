// Package tester is a testing package. Its contents will change regularly as new features are developed.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/nfnt/resize"
	"go.uber.org/zap"

	"github.com/nightmarlin/murum/layout"
	"github.com/nightmarlin/murum/provider"
	"github.com/nightmarlin/murum/renderer"
)

func main() {
	log, err := zap.NewDevelopment()
	if err != nil {
		_, _ = os.Stderr.WriteString("failed to initialise logger")
		return
	}
	log = log.Named("murum-tester")

	log.Info("starting murum in testing mode", zap.Strings("arguments", os.Args))
	if len(os.Args) < 2 {
		log.Error("second argument must be path to save output to", zap.Strings("args", os.Args))
		return
	}

	log.Info("parsing output path")
	path, err := filepath.Abs(os.Args[1])
	if err != nil {
		log.Error("second argument must be a valid file path", zap.String("arg", os.Args[1]), zap.Error(err))
		return
	}

	var (
		xCount, yCount = 16, 9
		base           = image.Rect(0, 0, 1920, 1080)
		randSRC        = rand.NewSource(time.Now().Unix())

		// layout.RandomPlacer(randSRC).Place(base, xCount, yCount)
		// layout.EvenPlacer().Place(base, xCount, yCount)
		placements = layout.VariedPlacer(
			image.Rectangle{Min: image.Point{X: -25, Y: -25}, Max: image.Point{X: 25, Y: 25}},
			randSRC,
		).
			Place(base, xCount, yCount)
	)

	log.Info("generating layout", zap.Int("x_count", xCount), zap.Int("y_count", yCount), zap.Any("rect", base))
	l, err := layout.Generate(base, placements)
	if err != nil {
		log.Error("failed to create layout", zap.Error(err))
		return
	}

	/*
		log.Info("generating filesystem provider")
		albumProvider, err := filesystem.New(filesystem.WithDefaults())
		if err != nil {
			log.Error("failed to init album info provider", zap.Error(err))
			return
		}
		log.Info("fetching files")
		info, err := albumProvider.Fetch(xCount * yCount)
		if err != nil {
			log.Error("failed to fetch album info", zap.Error(err))
			return
		}
	*/
	imgCount := xCount * yCount
	info := make([]provider.AlbumInfo, imgCount)
	for n := 0; n < imgCount; n++ {
		info[n] = provider.AlbumInfo{
			Art: NewBoundedUniform(
				color.RGBA{
					R: uint8(rand.Int() % 255),
					G: uint8(rand.Int() % 255),
					B: uint8(rand.Int() % 255),
					A: 255,
				},
				base,
				image.Rectangle{
					Min: image.Point{X: 480, Y: 270},
					Max: image.Point{X: 1440, Y: 810},
				},
			),
		}
	}

	log.Info("creating renderer")
	r, err := renderer.NewBasic(renderer.BasicWithInterpolationFunc(resize.Lanczos3))
	if err != nil {
		log.Error("failed to create renderer", zap.Error(err))
		return
	}

	log.Info("rendering...")
	img, err := r.Render(l, info)
	if err != nil {
		log.Error("failed to draw image", zap.Error(err))
		return
	}

	log.Info("creating file")
	f, err := os.Create(path)
	if err != nil {
		log.Error("unable to create file", zap.Error(err))
		return
	}
	defer func() { _ = f.Close() }()

	log.Info("saving file")
	err = png.Encode(f, img)
	if err != nil {
		log.Error("unable to encode image", zap.Error(err))
		return
	}

	log.Info("complete", zap.String("find your image at", path))
}

// BoundedUniform is a bounded Image of uniform color.
// It implements the color.Color, color.Model, and Image interfaces.
type BoundedUniform struct {
	C   color.Color
	B   image.Rectangle
	Not image.Rectangle
}

func (bu *BoundedUniform) RGBA() (r, g, b, a uint32)       { return bu.C.RGBA() }
func (bu *BoundedUniform) ColorModel() color.Model         { return bu }
func (bu *BoundedUniform) Convert(color.Color) color.Color { return bu.C }
func (bu *BoundedUniform) Bounds() image.Rectangle         { return bu.B }
func (bu *BoundedUniform) At(x, y int) color.Color {
	if (image.Point{X: x, Y: y}.In(bu.Not)) {
		return color.Transparent
	}
	if (image.Point{X: x, Y: y}.In(bu.B)) {
		return bu.C
	}
	return color.Transparent
}
func (bu *BoundedUniform) RGBA64At(x, y int) color.RGBA64 {
	r, g, b, a := bu.At(x, y).RGBA()
	return color.RGBA64{R: uint16(r), G: uint16(g), B: uint16(b), A: uint16(a)}
}
func (bu *BoundedUniform) Opaque() bool {
	_, _, _, a := bu.C.RGBA()
	return a == 0xffff
}
func NewBoundedUniform(c color.Color, bounds, trans image.Rectangle) *BoundedUniform {
	return &BoundedUniform{C: c, B: bounds.Canon(), Not: trans.Canon()}
}

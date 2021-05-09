// Package tester is a testing package. Its contents will change regularly as new features are developed.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/nightmarlin/murum/layout"
	"go.uber.org/zap"
)

// colors generates a set of random colors to fill a voronoi chart with
func colors(n int) []color.RGBA {
	rand.Seed(time.Now().Unix())

	var c []color.RGBA
	for i := 0; i < n; i++ {
		c = append(c, color.RGBA{
			R: uint8(rand.Intn(0xFF)),
			G: uint8(rand.Intn(0xFF)),
			B: uint8(rand.Intn(0xFF)),
			A: 0xFF,
		})
	}
	return c
}

func DrawVoronoi(log *zap.SugaredLogger, rect image.Rectangle, allPoints []image.Point) *image.RGBA {
	img := image.NewRGBA(rect)
	regions, err := layout.Generate(rect, allPoints)
	if err != nil {
		log.Errorw("failed to generate regions", zap.Error(err))
		return nil
	}

	c := colors(len(regions))
	for i := range regions {
		regionColor := c[i]

		for _, p := range regions[i].Pixels() {
			img.Set(p.X, p.Y, regionColor)
		}
	}

	return img
}

func main() {
	l, err := zap.NewDevelopment()
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, "failed to initialise logger")
		return
	}
	log := l.Sugar().Named("murum-tester")

	log.Infow("starting murum in testing mode", zap.Strings("arguments", os.Args))
	if len(os.Args) < 2 {
		log.Errorw("second argument must be path to save output to", zap.Strings("args", os.Args))
		return
	}

	path, err := filepath.Abs(os.Args[1])

	img := DrawVoronoi(
		log,
		image.Rect(0, 0, 1920, 1080),
		func() []image.Point {
			var res []image.Point
			for i := 0; i < 12; i++ {
				res = append(res, image.Point{X: rand.Intn(1920), Y: rand.Intn(1080)})
			}
			return res
		}(),
	)

	f, err := os.Create(path)
	if err != nil {
		log.Errorw("unable to create file", zap.Error(err))
		return
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		log.Errorw("unable to encode image", zap.Error(err))
		return
	}

	log.Infow("complete", "find your image at", path)
}

// Package tester is a testing package. Its contents will change regularly as new features are developed.
package main

import (
	"fmt"
	"image"
	"image/png"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/nfnt/resize"
	"go.uber.org/zap"

	"github.com/nightmarlin/murum/layout"
	"github.com/nightmarlin/murum/provider/filesystem"
	"github.com/nightmarlin/murum/renderer"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	log, err := zap.NewDevelopment()
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, "failed to initialise logger")
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
		xCount, yCount = 2, 2
		base           = image.Rect(0, 0, 1920, 1080)
	)

	log.Info("generating layout", zap.Int("x_count", xCount), zap.Int("y_count", yCount), zap.Any("rect", base))
	l, err := layout.Generate(base, layout.EvenPlacer().Place(base, xCount, yCount))
	if err != nil {
		log.Error("failed to create layout", zap.Error(err))
		return
	}

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

	log.Info("rendering...")
	r, err := renderer.NewBasic(renderer.BasicWithInterpolationFunc(resize.Lanczos3))
	if err != nil {
		log.Error("failed to create renderer", zap.Error(err))
		return
	}

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

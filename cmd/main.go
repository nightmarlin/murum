// package main is the entrypoint for the murum cli.
package main

import (
	"flag"
	"image"
	"image/png"
	"math/rand"
	"os"
	"path/filepath"

	"go.uber.org/zap"

	"github.com/Nightmarlin/murum"
)

var log *zap.SugaredLogger

func init() {
	verbose := flag.Bool("v", false, "log debug-level data")
	flag.Parse()

	var l *zap.Logger
	var err error
	if verbose != nil && *verbose {
		l, err = zap.NewDevelopment()
	} else {
		l, err = zap.NewProduction()
	}

	if err != nil {
		panic(err)
	}

	log = l.Sugar().Named("cmd")
}

func main() {
	log.Infow("starting murum in cli mode", zap.Strings("arguments", os.Args))
	if len(os.Args) < 2 {
		log.Errorw("second argument must be path to save output to", zap.Strings("args", os.Args))
		return
	}

	path, err := filepath.Abs(os.Args[1])

	img := murum.DrawVoronoi(
		image.Rect(0, 0, 1920, 1080),
		func() []image.Point {
			var res []image.Point
			for i := 0; i < 15; i++ {
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

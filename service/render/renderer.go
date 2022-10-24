package render

import (
	"image"
	"math"

	layout "github.com/nightmarlin/murum/protos/layout/v1"
	"github.com/nightmarlin/murum/provider"
)

// Renderer allows for murum image rendering
type Renderer interface {
	// Render draws an image with the bounding rectangle rect filled with the pattern defined in the
	// layout.L. Each L section will be filled with an image based on an entry from the
	// provider.AlbumInfo slice. If len(l) > len(ai) then some regions may be left blank. If len(ai) >
	// len(l) then not all Album Infos need to be used.
	Render(l *layout.Layout, ai []provider.AlbumInfo) (image.Image, error)
}

// GetBounds finds the bounding image.Rectangle for a set of points
func GetBounds(ps []image.Point) image.Rectangle {
	if len(ps) == 0 {
		return image.Rectangle{}
	}

	min, max := ps[0], ps[0]
	for _, p := range ps {
		if p.X < min.X {
			min.X = p.X
		}
		if p.Y < min.Y {
			min.Y = p.Y
		}

		if p.X > max.X {
			max.X = p.X
		}
		if p.Y > max.Y {
			max.Y = p.Y
		}
	}

	return image.Rectangle{
		Min: min,
		// Add {1, 1} as image.Rectangle max values are exclusive
		Max: max.Add(image.Point{X: 1, Y: 1}),
	}
}

// ScaleRect scales a rectangle to fill at least the area of some target rectangle while maintaining
// its original aspect ratio. If any dimension of either passed rectangle is 0 the bounding
// rectangle is returned. The rectangle will be scaled about the origin (image.Point{0, 0}).
func ScaleRect(toScale, bounding image.Rectangle) image.Rectangle {
	if toScale.Dx() == 0 || bounding.Dx() == 0 || toScale.Dy() == 0 || bounding.Dy() == 0 {
		return bounding
	}

	var (
		xSF = float64(bounding.Dx()) / float64(toScale.Dx())
		ySF = float64(bounding.Dy()) / float64(toScale.Dy())

		scaleFactor = math.Max(xSF, ySF)
		doScale     = func(i int) int { return int(math.Ceil(float64(i) * scaleFactor)) }
	)

	return image.Rectangle{
		Min: image.Point{X: doScale(toScale.Min.X), Y: doScale(toScale.Min.Y)},
		Max: image.Point{X: doScale(toScale.Max.X), Y: doScale(toScale.Max.Y)},
	}
}

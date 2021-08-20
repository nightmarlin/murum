package layout

import (
	"image"
	"math"
	"math/rand"
)

// Placer allows points to be placed on an image.Rectangle
type Placer interface {
	// Place takes an image.Rectangle and calculates where best to place xCount*yCount center points.
	// If xCount == 0 or yCount == 0, the following expression should resolve to true:
	//   len(Place(rect, xCount, yCount)) == 0
	// A Placer should support rectangles with negative regions
	Place(rect image.Rectangle, xCount int, yCount int) []image.Point
}

// RandomPlacer returns a Placer that uses a rand.Source to randomly Place points on the passed
// rectangle.
func RandomPlacer(src rand.Source) *randomPlacer {
	return &randomPlacer{src: src}
}

// EvenPlacer returns a Placer that places xCount*yCount points evenly across a rectangle. If xCount
// or yCount <= 0, nil is returned.
func EvenPlacer() *evenPlacer {
	return &evenPlacer{}
}

// VariedPlacer generates a new Placer. It's based on EvenPlacer, but then moves the centered point
// a random amount within the varyRect (with the rectangle overlaid such that vr(0, 0) is aligned to
// the centered point).
//   vr.RMin.X <= randomXVariance < vr.RMax.X
//   vr.RMin.Y <= randomYVariance < vr.RMax.Y
func VariedPlacer(varyRect image.Rectangle, src rand.Source) *variedPlacer {
	return &variedPlacer{varyRect: varyRect, src: src}
}

type randomPlacer struct {
	src rand.Source
}

func (rp *randomPlacer) Place(rect image.Rectangle, xCount int, yCount int) []image.Point {
	if xCount <= 0 || yCount <= 0 {
		return nil
	}

	var (
		total = xCount * yCount
		res   = make([]image.Point, total)
	)

	for i := 0; i < total; i++ {
		res[i] = image.Point{
			X: int(rp.src.Int63()) % rect.Dx(),
			Y: int(rp.src.Int63()) % rect.Dy(),
		}
	}

	return res
}

type evenPlacer struct{}

func (ep *evenPlacer) Place(r image.Rectangle, xCount int, yCount int) []image.Point {
	if xCount <= 0 || yCount <= 0 {
		return nil
	}

	var (
		n    = 0
		res  = make([]image.Point, xCount*yCount)
		xGap = int(math.Floor(float64(r.Dx()) / float64(xCount) / 2.0))
		yGap = int(math.Floor(float64(r.Dy()) / float64(yCount) / 2.0))
	)

	for x := 0; x < xCount; x++ {
		for y := 0; y < yCount; y++ {
			res[n] = image.Point{X: xGap + (2*xGap)*x, Y: yGap + (2*yGap)*y}
			n += 1
		}
	}

	return res
}

type variedPlacer struct {
	evenPlacer

	varyRect image.Rectangle
	src      rand.Source
}

func (vp *variedPlacer) Place(r image.Rectangle, xCount int, yCount int) []image.Point {
	points := vp.evenPlacer.Place(r, xCount, yCount)

	for idx := range points {
		points[idx].X += vp.varyRect.Min.X + int(vp.src.Int63())%vp.varyRect.Dx()
		points[idx].Y += vp.varyRect.Min.Y + int(vp.src.Int63())%vp.varyRect.Dy()
	}

	return points
}

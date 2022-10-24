package layout

import (
	"image"
	"math"
	"math/rand"
)

// A Placer takes an image.Rectangle and calculates where best to place xCount*
// yCount nearestCenter points. If xCount <= 0 or yCount <= 0, the following
// expression should resolve to true:
//
//	len(Place(rect, xCount, yCount)) == 0
//
// A Placer should support rectangles with negative regions
type Placer func(rect image.Rectangle, xCount int, yCount int) []image.Point

// RandomPlacer returns a Placer that uses a rand.Source to randomly Place
// points on the passed rectangle.
func RandomPlacer(src rand.Source) Placer {
	return func(rect image.Rectangle, xCount int, yCount int) []image.Point {
		if xCount <= 0 || yCount <= 0 {
			return nil
		}

		var (
			total = xCount * yCount
			res   = make([]image.Point, total)
		)

		for i := 0; i < total; i++ {
			res[i] = image.Point{
				X: int(src.Int63()) % rect.Dx(),
				Y: int(src.Int63()) % rect.Dy(),
			}
		}

		return res
	}
}

// EvenPlacer returns a Placer that places xCount*yCount points evenly across a
// rectangle. If xCount or yCount <= 0, nil is returned.
func EvenPlacer() Placer {
	return func(r image.Rectangle, xCount int, yCount int) []image.Point {
		if xCount <= 0 || yCount <= 0 {
			return nil
		}

		var (
			n    = 0
			res  = make([]image.Point, xCount*yCount)
			xGap = int(math.Floor(float64(r.Dx()) / float64(xCount) / 2))
			yGap = int(math.Floor(float64(r.Dy()) / float64(yCount) / 2))
		)

		for x := 0; x < xCount; x++ {
			for y := 0; y < yCount; y++ {
				res[n] = image.Point{X: xGap + (2*xGap)*x, Y: yGap + (2*yGap)*y}
				n += 1
			}
		}

		return res
	}
}

// VariedPlacer generates a new Placer. It's based on EvenPlacer, but then moves
// the centered point a random amount within the varyRect (with the rectangle
// overlaid such that vr(0, 0) is aligned to the centered point).
//
//	varyRect.Min.X <= randomXVariance < varyRect.Max.X
//	varyRect.Min.Y <= randomYVariance < varyRect.Max.Y
func VariedPlacer(varyRect image.Rectangle, src rand.Source) Placer {
	return func(r image.Rectangle, xCount int, yCount int) []image.Point {
		points := EvenPlacer()(r, xCount, yCount)

		for idx := range points {
			points[idx].X += varyRect.Min.X + int(src.Int63())%varyRect.Dx()
			points[idx].Y += varyRect.Min.Y + int(src.Int63())%varyRect.Dy()
		}

		return points
	}
}

// TriangularPlacer creates and returns a Placer that evenly distributes points
// across a rectangle in a triangular pattern. When rendered with Generate, the
// resulting L will have hexagon-shaped sections
// TODO: Triangular (to create hexagonal voronoi cells)
func TriangularPlacer() Placer {
	return func(rect image.Rectangle, xCount int, yCount int) []image.Point {
		return make([]image.Point, xCount*yCount)
	}
}

// WeightedPlacer creates and returns a Placer that pushes/pulls points to/from
// a given center point. It is based on the EvenPlacer, and then moves points to
// create the gravitational effect.
//
//	center : Where to apply the gravity from.
//	gravity: How strong the effect is.
//	          +ve: Points are pulled towards center;
//	          -ve: Points are pushed away from center
//
// TODO: Weighted (place more points closer to a given point, and fewer further away. Or vice versa)
func WeightedPlacer(center image.Point, gravity float64) Placer {
	return func(r image.Rectangle, xCount int, yCount int) []image.Point {
		return make([]image.Point, xCount*yCount)
	}
}

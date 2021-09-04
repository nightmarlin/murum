// Package layout provides voronoi generation, the layout mechanic behind murum's images. Each region is defined as the
// set of points closest to a reference point. It exposes the Generate function, which creates a set of sets of
// image.Point s.
package layout

import (
	"errors"
	"image"
	"math"
)

type L struct {
	Bounds image.Rectangle
	Points map[image.Point][]image.Point
}

// Generate creates an L map, where each key is a single value from centerPoints and its
// corresponding value is the set of all points closest to the key. If two centerPoints have equal
// distance, the first point in centerPoints will be used.
func Generate(rect image.Rectangle, centerPoints []image.Point) (*L, error) {
	if len(centerPoints) == 0 {
		return nil, errors.New("centerPoints must have at least one element")
	}

	l := &L{Bounds: rect, Points: make(map[image.Point][]image.Point)}
	for x := rect.Min.X; x < rect.Max.X; x++ {
		for y := rect.Min.Y; y < rect.Max.Y; y++ {
			p := image.Point{X: x, Y: y}

			nearest := centerPoints[0]
			for i := 1; i < len(centerPoints); i++ {
				if distance(p, centerPoints[i]) < distance(p, nearest) {
					nearest = centerPoints[i]
				}
			}

			l.Points[nearest] = append(l.Points[nearest], p)
		}
	}

	return l, nil
}

// distance performs a simple pythagorean calculation:
//   a^2 + b^2 = c^2
//   ∴ c = |sqrt(a^2 + b^2)|
// a here is the x-difference and b is the y-difference
func distance(p1, p2 image.Point) float64 {
	// math.Sqrt is always > 0. Or NaN. Probably won't be NaN (no complex numbers here!)
	return math.Sqrt(math.Pow(float64(p2.X-p1.X), 2) + math.Pow(float64(p2.Y-p1.Y), 2))
}

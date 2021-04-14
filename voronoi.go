package murum

import (
	"image"
	"image/color"
	"math"
	"math/rand"
	"time"
)

// distance performs a simple pythagorean calculation: a^2 + b^2 = c^2 therefore c = |sqrt(a^2 + b^2)|.
// a here is the x-difference and b is the y-difference
func distance(p1, p2 image.Point) float64 {
	xDiff := float64(p2.X - p1.X)
	yDiff := float64(p2.Y - p1.Y)
	square := math.Pow(xDiff, 2) + math.Pow(yDiff, 2)
	return math.Sqrt(square) // math.Sqrt is always > 0. Or NaN. Probably won't be NaN.
}

// getNearestPoint retrieves the nearest point in allPoints to p (and its index). It panics if allPoints is nil or
// 0-length. If the distance between two other points is the same, the earliest point in allPoints will be selected
func getNearestPoint(p image.Point, allPoints []image.Point) (image.Point, int) {
	nearestIdx := 0

	for i := 1; i < len(allPoints); i++ {
		if distance(allPoints[nearestIdx], p) > distance(allPoints[i], p) {
			nearestIdx = i
		}
	}

	return allPoints[nearestIdx], nearestIdx
}

// generateRegions creates a slice of slices of points in Rect. The slice at regions[i] is the set of points closest to
// allPoints[i].
func generateRegions(rect image.Rectangle, allPoints []image.Point) [][]image.Point {
	regions := make([][]image.Point, len(allPoints))

	// TODO: parallelize in chunks / create a worker pool
	for x := rect.Min.X; x < rect.Max.X; x++ {
		for y := rect.Min.Y; y < rect.Max.Y; y++ {
			p := image.Point{X: x, Y: y}
			_, idx := getNearestPoint(p, allPoints)
			regions[idx] = append(regions[idx], p)
		}
	}

	return regions
}

// colors contains a set of random colors generated at runtime
var colors = func() []color.RGBA {
	rand.Seed(time.Now().Unix())

	var c []color.RGBA
	for i := 0; i < 16; i++ {
		c = append(c, color.RGBA{
			R: uint8(rand.Intn(0xFF)),
			G: uint8(rand.Intn(0xFF)),
			B: uint8(rand.Intn(0xFF)),
			A: 0xFF,
		})
	}
	return c
}()

func DrawVoronoi(rect image.Rectangle, allPoints []image.Point) *image.RGBA {
	img := image.NewRGBA(rect)
	regions := generateRegions(rect, allPoints)

	for i, r := range regions {
		regionColor := colors[int(math.Mod(float64(i), float64(len(colors))))]

		for _, p := range r {
			img.Set(p.X, p.Y, regionColor)
		}
	}

	return img
}

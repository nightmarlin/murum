package layout

import (
	"image"
	"math"
)

// distance performs a simple pythagorean calculation: a^2 + b^2 = c^2 therefore c = |sqrt(a^2 + b^2)|.
// a here is the x-difference and b is the y-difference
func distance(p1, p2 image.Point) float64 {
	xDiff := float64(p2.X - p1.X)
	yDiff := float64(p2.Y - p1.Y)
	square := math.Pow(xDiff, 2) + math.Pow(yDiff, 2)
	return math.Sqrt(square) // math.Sqrt is always > 0. Or NaN. Probably won't be NaN.
}

// getNearestPoint retrieves the nearest reference-point in layout to p (and its index). It panics if layout is nil or
// 0-length. If the distance between two other points is the same, the earliest point in layout will be selected
func getNearestPoint(p image.Point, layout L) image.Point {
	nearestIdx := 0

	for i := 1; i < len(layout); i++ {
		if distance(layout[nearestIdx].center, p) > distance(layout[i].center, p) {
			nearestIdx = i
		}
	}

	return layout[nearestIdx].center
}

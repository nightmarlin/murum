// Package layout provides voronoi generation, the layout mechanic behind murum's images. Each
// region is defined as the set of points closest to a reference point. It exposes the Generate
// function, which creates a set of sets of image.Point s.
package layout

import (
	"errors"
	"image"
	"math"
	"runtime"
)

type L struct {
	Bounds image.Rectangle
	Points map[image.Point][]image.Point
}

func runCount() int {
	var (
		maxProcs  = runtime.GOMAXPROCS(0)
		numGR     = runtime.NumGoroutine()
		available = maxProcs - numGR - 1 // save 1 for the reader goroutine
	)

	if available < 1 {
		return 1
	}
	return available
}

type pointPair struct {
	refPoint      image.Point
	nearestCenter image.Point
}

func goGenerate(centerPoints []image.Point, in <-chan image.Point, out chan<- pointPair) {
	for p := range in {
		nearest := centerPoints[0]
		for i := 1; i < len(centerPoints); i++ {
			if distance(p, centerPoints[i]) < distance(p, nearest) {
				nearest = centerPoints[i]
			}
		}

		out <- pointPair{nearestCenter: nearest, refPoint: p}
	}
}

// Generate creates an L map, where each key is a single value from centerPoints and its
// corresponding value is the set of all points closest to the key. If two centerPoints have equal
// distance, the first point in centerPoints will be used.
func Generate(rect image.Rectangle, centerPoints []image.Point) (L, error) {
	if len(centerPoints) == 0 {
		return L{}, errors.New("centerPoints must have at least one element")
	}

	var (
		gc         = runCount()
		pointsChan = make(chan image.Point)
		outChan    = make(chan pointPair)

		ready = make(chan struct{}, 1)
		l     = L{Bounds: rect.Canon(), Points: make(map[image.Point][]image.Point)}
	)

	// Spin up GoRoutines
	for i := 0; i < gc; i++ {
		go goGenerate(centerPoints, pointsChan, outChan)
	}

	// Read from outChan and apply pointPairs
	go func(limit int) {
		processed := 0
		for pp := range outChan {
			l.Points[pp.nearestCenter] = append(l.Points[pp.nearestCenter], pp.refPoint)

			processed += 1
			if processed == limit {
				ready <- struct{}{}
			}
		}
	}(rect.Dx() * rect.Dy())

	// Send all points for processing
	for x := rect.Min.X; x < rect.Max.X; x++ {
		for y := rect.Min.Y; y < rect.Max.Y; y++ {
			pointsChan <- image.Point{X: x, Y: y}
		}
	}
	close(pointsChan)

	_ = <-ready
	close(outChan)
	return l, nil
}

// distance performs a simple pythagorean calculation:
//
//	a^2 + b^2 = c^2
//	∴ c = |sqrt(a^2 + b^2)|
//
// a here is the x-difference and b is the y-difference
func distance(p1, p2 image.Point) float64 {
	// math.Sqrt is always > 0. Or NaN. Probably won't be NaN (no complex numbers here!)
	return math.Sqrt(math.Pow(float64(p2.X-p1.X), 2) + math.Pow(float64(p2.Y-p1.Y), 2))
	// TODO: Fast inverse square root.
}

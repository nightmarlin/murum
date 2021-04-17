// Package layout provides voronoi generation, the layout mechanic behind murum's images. Each region is defined as the
// set of points closest to a reference point. It exposes the Generate function, which creates a set of sets of
// image.Point s.
package layout

import (
	"errors"
	"image"
	"sync"
)

// generatorWorker allows for parallelization of pixel processing. Each worker calculates the nearest point in allPoints
// and then adds the current pixel
func generatorWorker(
	wg *sync.WaitGroup,
	layout *L,
	pixels <-chan image.Point,
) {
	defer wg.Done()
	if layout == nil {
		return
	}

	for p := range pixels {
		ref := getNearestPoint(p, *layout)
		layout.addClosestPixel(ref, p)
	}
}

// Generate creates a layout.L slice. Each Region in the returned L will have a center in allPoints and contain the
// points in rect closest to that centre.
func Generate(rect image.Rectangle, allPoints []image.Point) (L, error) {
	if len(allPoints) == 0 {
		return nil, errors.New("allPoints cannot be nil or 0-length")
	}

	var layout L
	for _, ref := range allPoints {
		layout = append(layout, Region{center: ref})
	}

	wg := &sync.WaitGroup{}
	pixels := make(chan image.Point)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go generatorWorker(wg, &layout, pixels)
	}

	for x := rect.Min.X; x < rect.Max.X; x++ {
		for y := rect.Min.Y; y < rect.Max.Y; y++ {
			pixels <- image.Point{X: x, Y: y}
		}
	}

	close(pixels)
	wg.Wait()

	return layout, nil
}

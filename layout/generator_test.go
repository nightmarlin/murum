package layout

import (
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generateRegions(t *testing.T) {
	rect := image.Rect(0, 0, 3, 3)
	allPoints := []image.Point{
		{X: 0, Y: 0},
		{X: 2, Y: 2},
	}

	expect := L{
		{
			center: image.Point{X: 0, Y: 0},
			pixels: []image.Point{ // Selects first in list for closeness where distance is tied!
				{X: 0, Y: 0},
				{X: 0, Y: 1},
				{X: 0, Y: 2},
				{X: 1, Y: 0},
				{X: 1, Y: 1},
				{X: 2, Y: 0},
			},
		},
		{
			center: image.Point{X: 2, Y: 2},
			pixels: []image.Point{
				{X: 1, Y: 2},
				{X: 2, Y: 1},
				{X: 2, Y: 2},
			},
		},
	}

	layout, err := Generate(rect, allPoints)

	assert.NoError(t, err)
	if assert.NotNil(t, layout) && assert.Equal(t, len(expect), len(layout)) {
		// Order only matters for region centers, not the pixels associated with them!
		for i := 0; i < len(expect); i++ {
			assert.Equal(t, expect[i].center, layout[i].center)
			assert.ElementsMatch(t, expect[i].pixels, layout[i].pixels)
		}
	}
}

// The largest point will be at res[0], the smallest at res[len(res)-1]. They will get exponentially closer together the
// smaller the value of res[n] is.
func generateNPoints(n, bigN int) []image.Point {
	res := make([]image.Point, n)
	for i := 1; i <= n; i++ {
		v := bigN / i
		res = append(res, image.Point{X: v, Y: v})
	}
	return res
}

func Benchmark_generateRegions(b *testing.B) {
	for i := 0; i < b.N; i++ {
		allPoints := generateNPoints(10, b.N)
		r := image.Rect(0, 0, b.N, b.N)

		_, err := Generate(r, allPoints)
		if err != nil {
			b.Fatalf("Generate should never return an error. Got: %v", err)
		}
	}
}

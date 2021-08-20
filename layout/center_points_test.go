package layout

import (
	"image"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// testingSource contains a simple counter starting at 0 that increments each time Int63 is called.
// The result of Int63 increments once every 2 calls.
type testingSource struct {
	genCount int64
}

// Seed resets the internal counter of the testingSource.
func (ts *testingSource) Seed(int64) {
	ts.genCount = 0
}

// Int63 ticks up once every 2 calls.
func (ts *testingSource) Int63() int64 {
	g := ts.genCount / 2
	ts.genCount += 1
	return g
}

// logRect prints the passed imageRectangle to the testing.T log. Any image.Point in the slice that
// is also in the rect is rendered as a "XX" - any other point is rendered as a "[]".
func logRect(t *testing.T, r image.Rectangle, ps []image.Point) {
	var b strings.Builder
	for x := r.Min.X; x < r.Max.X; x++ {
		b.WriteString("\n")
		for y := r.Min.Y; y < r.Max.Y; y++ {
			w := "[]"
			for i := range ps {
				if ps[i] == (image.Point{X: x, Y: y}) {
					w = "XX"
					break
				}
			}
			b.WriteString(w)
		}
	}
	t.Log(b.String())
}

func TestRandomPlacer(t *testing.T) {
	testTable := []struct {
		Name string

		Rect   image.Rectangle
		XCount int
		YCount int

		Expect []image.Point
	}{
		{Name: "nil if xCount is 0", XCount: 0, YCount: 1, Expect: nil},
		{Name: "nil if yCount is 0", XCount: 1, YCount: 0, Expect: nil},
		{
			Name:   "places points in expected locations",
			Rect:   image.Rectangle{Max: image.Point{X: 3, Y: 3}},
			XCount: 2,
			YCount: 2,
			Expect: []image.Point{
				{X: 0, Y: 0},
				{X: 1, Y: 1},
				{X: 2, Y: 2},
				{X: 0, Y: 0}, // Reset back to minimum bound at rect edge
			},
		},
	}

	for _, entry := range testTable {
		entry := entry
		t.Run(entry.Name, func(t *testing.T) {
			t.Parallel()

			res := RandomPlacer(&testingSource{}).Place(entry.Rect, entry.XCount, entry.YCount)

			if !assert.Equal(t, entry.Expect, res) {
				logRect(t, entry.Rect, res)
			}
		})
	}
}

func TestEvenPlacer(t *testing.T) {
	testTable := []struct {
		Name string

		Rect   image.Rectangle
		XCount int
		YCount int

		Expect []image.Point
	}{
		{Name: "nil if xCount is 0", XCount: 0, YCount: 1, Expect: nil},
		{Name: "nil if yCount is 0", XCount: 1, YCount: 0, Expect: nil},
		{
			Name:   "20x20 split 1x2",
			Rect:   image.Rectangle{Max: image.Point{X: 20, Y: 20}},
			XCount: 1,
			YCount: 2,
			Expect: []image.Point{{X: 10, Y: 5}, {X: 10, Y: 15}},
		},
		{
			Name:   "20x20 split 2x2",
			Rect:   image.Rectangle{Max: image.Point{X: 20, Y: 20}},
			XCount: 2,
			YCount: 2,
			Expect: []image.Point{{X: 5, Y: 5}, {X: 5, Y: 15}, {X: 15, Y: 5}, {X: 15, Y: 15}},
		},
		{
			Name:   "40x40 split 2x2",
			Rect:   image.Rectangle{Max: image.Point{X: 40, Y: 40}},
			XCount: 2,
			YCount: 2,
			Expect: []image.Point{{X: 10, Y: 10}, {X: 10, Y: 30}, {X: 30, Y: 10}, {X: 30, Y: 30}},
		},
	}

	for _, entry := range testTable {
		entry := entry
		t.Run(entry.Name, func(t *testing.T) {
			t.Parallel()

			res := EvenPlacer().Place(entry.Rect, entry.XCount, entry.YCount)

			if !assert.Equal(t, entry.Expect, res) {
				logRect(t, entry.Rect, res)
			}
		})
	}
}

func TestVariedPlacer(t *testing.T) {
	testTable := []struct {
		Name string

		VaryRect image.Rectangle

		Rect   image.Rectangle
		XCount int
		YCount int

		Expect []image.Point
	}{
		{Name: "nil if xCount is 0", XCount: 0, YCount: 1, Expect: nil},
		{Name: "nil if yCount is 0", XCount: 1, YCount: 0, Expect: nil},
		{
			Name:     "20x20 split 1x2",
			VaryRect: image.Rectangle{Max: image.Point{X: 3, Y: 3}},
			Rect:     image.Rectangle{Max: image.Point{X: 20, Y: 20}},
			XCount:   1,
			YCount:   2,
			Expect: []image.Point{
				{X: 10, Y: 5},  // offset: 0
				{X: 11, Y: 16}, // offset: 1
			},
		},
		{
			Name:     "20x20 split 2x2",
			VaryRect: image.Rectangle{Max: image.Point{X: 3, Y: 3}},
			Rect:     image.Rectangle{Max: image.Point{X: 20, Y: 20}},
			XCount:   2,
			YCount:   2,
			Expect: []image.Point{
				{X: 5, Y: 5},   // offset: 0
				{X: 6, Y: 16},  // offset: 1
				{X: 17, Y: 7},  // offset: 2
				{X: 15, Y: 15}, // offset: 0 (reset at edge of varyRect)
			},
		},
		{
			Name:     "40x40 split 2x2",
			VaryRect: image.Rectangle{Max: image.Point{X: 3, Y: 3}},
			Rect:     image.Rectangle{Max: image.Point{X: 40, Y: 40}},
			XCount:   2,
			YCount:   2,
			Expect: []image.Point{
				{X: 10, Y: 10}, // offset: 0
				{X: 11, Y: 31}, // offset: 1
				{X: 32, Y: 12}, // offset: 2
				{X: 30, Y: 30}, // offset: 0 (reset at edge of varyRect)
			},
		},
	}

	for _, entry := range testTable {
		entry := entry
		t.Run(entry.Name, func(t *testing.T) {
			t.Parallel()

			res := VariedPlacer(entry.VaryRect, &testingSource{}).
				Place(entry.Rect, entry.XCount, entry.YCount)

			if !assert.Equal(t, entry.Expect, res) {
				logRect(t, entry.Rect, res)
			}
		})
	}
}

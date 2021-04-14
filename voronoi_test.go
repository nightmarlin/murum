package murum

import (
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_distance(t *testing.T) {
	testTable := []struct {
		Name     string
		P1       image.Point
		P2       image.Point
		Distance float64
	}{
		{
			Name:     "identical points have distance 0",
			P1:       image.Point{},
			P2:       image.Point{},
			Distance: 0,
		},
		{
			Name:     "pythagorean triple 3-4-5 is correct",
			P1:       image.Point{},
			P2:       image.Point{X: 3, Y: 4},
			Distance: 5,
		},
		{
			Name:     "pythagorean triple 5-12-13 is correct",
			P1:       image.Point{},
			P2:       image.Point{X: 5, Y: 12},
			Distance: 13,
		},
	}

	for _, entry := range testTable {
		entry := entry
		t.Run(entry.Name, func(t *testing.T) {
			t.Parallel()

			d := distance(entry.P1, entry.P2)
			assert.Equal(t, entry.Distance, d)
		})
	}
}

func Test_getNearestPoint_panics(t *testing.T) {
	testTable := []struct {
		Name      string
		AllPoints []image.Point
	}{
		{
			Name:      "panics with nil slice",
			AllPoints: nil,
		},
		{
			Name:      "panics with empty slice",
			AllPoints: []image.Point{},
		},
	}

	for _, entry := range testTable {
		entry := entry
		t.Run(entry.Name, func(t *testing.T) {
			t.Parallel()

			shouldPanic := func() { getNearestPoint(image.Point{}, entry.AllPoints) }
			assert.Panics(t, shouldPanic)
		})
	}
}

func Test_getNearestPoint(t *testing.T) {
	testTable := []struct {
		Name string

		RefPoint  image.Point
		AllPoints []image.Point

		NearestPoint image.Point
		PointIndex   int
	}{
		{
			Name:     "always nearest to self",
			RefPoint: image.Point{},
			AllPoints: []image.Point{
				{},
				{X: 50, Y: 50},
				{X: 100, Y: 100},
			},
			NearestPoint: image.Point{},
			PointIndex:   0,
		},
		{
			Name:     "correctly selects nearest (at idx 0)",
			RefPoint: image.Point{X: 10, Y: 10},
			AllPoints: []image.Point{
				{},
				{X: 50, Y: 50},
				{X: 100, Y: 100},
			},
			NearestPoint: image.Point{},
			PointIndex:   0,
		},
		{
			Name:     "correctly selects nearest (at idx 2)",
			RefPoint: image.Point{X: 90, Y: 90},
			AllPoints: []image.Point{
				{},
				{X: 50, Y: 50},
				{X: 100, Y: 100},
			},
			NearestPoint: image.Point{X: 100, Y: 100},
			PointIndex:   2,
		},
		{
			Name:     "points on boundary select earliest element",
			RefPoint: image.Point{X: 75, Y: 75},
			AllPoints: []image.Point{
				{},
				{X: 50, Y: 50},
				{X: 100, Y: 100},
			},
			NearestPoint: image.Point{X: 50, Y: 50},
			PointIndex:   1,
		},
	}

	for _, entry := range testTable {
		entry := entry
		t.Run(entry.Name, func(t *testing.T) {
			t.Parallel()

			p, i := getNearestPoint(entry.RefPoint, entry.AllPoints)

			assert.Equal(t, entry.NearestPoint, p)
			assert.Equal(t, entry.PointIndex, i)
		})
	}
}

func Test_generateRegions(t *testing.T) {
	rect := image.Rect(0, 0, 3, 3)
	allPoints := []image.Point{
		{X: 0, Y: 0},
		{X: 2, Y: 2},
	}

	expect := [][]image.Point{
		{ // Selects first in list for closeness where distance is tied!
			{X: 0, Y: 0},
			{X: 0, Y: 1},
			{X: 0, Y: 2},
			{X: 1, Y: 0},
			{X: 1, Y: 1},
			{X: 2, Y: 0},
		},
		{
			{X: 1, Y: 2},
			{X: 2, Y: 1},
			{X: 2, Y: 2},
		},
	}

	regions := generateRegions(rect, allPoints)
	assert.Equal(t, expect, regions)
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

		generateRegions(r, allPoints)
	}
}

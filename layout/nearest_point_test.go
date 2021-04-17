package layout

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

func layoutWithFilledCenters(allPoints []image.Point) L {
	if allPoints == nil {
		return nil
	}
	l := make(L, len(allPoints))
	for i := range allPoints {
		l[i].center = allPoints[i]
	}
	return l
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

			shouldPanic := func() {
				getNearestPoint(
					image.Point{},
					layoutWithFilledCenters(entry.AllPoints),
				)
			}
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

			p := getNearestPoint(
				entry.RefPoint,
				layoutWithFilledCenters(entry.AllPoints),
			)

			assert.Equal(t, entry.NearestPoint, p)
		})
	}
}

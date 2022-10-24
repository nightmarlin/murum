package layout

import (
	"errors"
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertEqualL(t *testing.T, expect, actual L) bool {
	if !assert.Equal(t, expect.Bounds, actual.Bounds, "bounding rectangles must match") {
		return false
	}

	if !assert.Equal(t, len(expect.Points), len(actual.Points), "map lengths must match") {
		return false
	}

	for eKey := range expect.Points {
		found := false
		for aKey := range actual.Points {
			if eKey == aKey {
				found = true
				break
			}
		}

		if !assert.Truef(
			t,
			found,
			"both maps must have the same keys, couldn't find %v in actual", eKey,
		) {
			return false
		}

		if !assert.ElementsMatchf(
			t,
			expect.Points[eKey],
			actual.Points[eKey],
			"slices under entry %v must match", eKey,
		) {
			return false
		}
	}

	return true
}

func TestGenerate(t *testing.T) {
	testTable := []struct {
		Name string

		Rect   image.Rectangle
		Points []image.Point

		ExpectRes L
		ExpectErr error
	}{
		{
			Name:      "errors if centerPoints is nil",
			Points:    nil,
			ExpectRes: L{},
			ExpectErr: errors.New("centerPoints must have at least one element"),
		},
		{
			Name:      "errors if centerPoints is empty",
			Points:    make([]image.Point, 0),
			ExpectRes: L{},
			ExpectErr: errors.New("centerPoints must have at least one element"),
		},
		{
			Name:   "succeeds under normal circumstances",
			Rect:   image.Rectangle{Max: image.Point{X: 3, Y: 3}},
			Points: []image.Point{{X: 0, Y: 0}, {X: 2, Y: 2}},
			ExpectRes: L{
				Bounds: image.Rectangle{Max: image.Point{X: 3, Y: 3}},
				Points: map[image.Point][]image.Point{
					{X: 0, Y: 0}: {
						{X: 0, Y: 0},
						{X: 0, Y: 1},
						{X: 0, Y: 2},
						{X: 1, Y: 0},
						{X: 1, Y: 1},
						{X: 2, Y: 0},
					},
					{X: 2, Y: 2}: {
						{X: 1, Y: 2},
						{X: 2, Y: 1},
						{X: 2, Y: 2},
					},
				},
			},
		},
	}

	for _, entry := range testTable {
		entry := entry
		t.Run(entry.Name, func(t *testing.T) {
			t.Parallel()

			res, err := Generate(entry.Rect, entry.Points)
			if entry.ExpectErr != nil {
				assert.Nil(t, res)
				assert.EqualError(t, err, entry.ExpectErr.Error())
			} else {
				assert.NoError(t, err)
				assertEqualL(t, entry.ExpectRes, res)
			}
		})
	}
}

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

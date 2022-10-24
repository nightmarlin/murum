package render

import (
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBounds(t *testing.T) {
	testTable := []struct {
		Name string

		Points []image.Point
		Expect image.Rectangle
	}{
		{
			Name:   "nil slice returns 0-Rect",
			Points: nil,
			Expect: image.Rectangle{},
		},
		{
			Name:   "0-length slice returns 0-Rect",
			Points: make([]image.Point, 0),
			Expect: image.Rectangle{},
		},
		{
			Name:   "bounds for one point are correct",
			Points: []image.Point{{X: 1, Y: 1}},
			Expect: image.Rectangle{
				Max: image.Point{X: 1, Y: 1},
				Min: image.Point{X: 1, Y: 1},
			},
		},
		{
			Name: "bounds for many points are correct",
			Points: []image.Point{
				{X: -2, Y: 2},
				{X: 0, Y: 0},
				{X: 2, Y: -2},
			},
			Expect: image.Rectangle{
				Min: image.Point{X: -2, Y: -2},
				Max: image.Point{X: 2, Y: 2}},
		},
	}

	for _, entry := range testTable {
		entry := entry
		t.Run(entry.Name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, entry.Expect, GetBounds(entry.Points))
		})
	}
}

func TestScaleRect(t *testing.T) {
	baseRect := image.Rectangle{Max: image.Point{X: 1920, Y: 1080}}

	testTable := []struct {
		Name string

		Bounding image.Rectangle
		ToScale  image.Rectangle
		Expect   image.Rectangle
	}{
		{
			Name:     "if bounding X is 0, result is bounding",
			Bounding: image.Rectangle{Min: image.Point{X: 1, Y: 1}, Max: image.Point{X: 1, Y: 2}},
			Expect:   image.Rectangle{Min: image.Point{X: 1, Y: 1}, Max: image.Point{X: 1, Y: 2}},
		},
		{
			Name:     "if bounding Y is 0, result is bounding",
			Bounding: image.Rectangle{Min: image.Point{X: 1, Y: 1}, Max: image.Point{X: 2, Y: 1}},
			Expect:   image.Rectangle{Min: image.Point{X: 1, Y: 1}, Max: image.Point{X: 2, Y: 1}},
		},
		{
			Name:     "if toScale X is 0, result is bounding",
			Bounding: baseRect,
			ToScale:  image.Rectangle{Min: image.Point{X: 1, Y: 1}, Max: image.Point{X: 1, Y: 2}},
			Expect:   baseRect,
		},
		{
			Name:     "if toScale Y is 0, result is bounding",
			Bounding: baseRect,
			ToScale:  image.Rectangle{Min: image.Point{X: 1, Y: 1}, Max: image.Point{X: 2, Y: 1}},
			Expect:   baseRect,
		},
		{
			Name:     "960x540 -> 1920x1080",
			Bounding: baseRect,
			ToScale:  image.Rectangle{Max: image.Point{X: 960, Y: 540}},
			Expect:   image.Rectangle{Max: image.Point{X: 1920, Y: 1080}},
		},
		{
			Name:     "960x1080 -> 1920x2160",
			Bounding: baseRect,
			ToScale:  image.Rectangle{Max: image.Point{X: 960, Y: 1080}},
			Expect:   image.Rectangle{Max: image.Point{X: 1920, Y: 2160}},
		},
		{
			Name:     "1920x540 -> 3840x1080",
			Bounding: baseRect,
			ToScale:  image.Rectangle{Max: image.Point{X: 1920, Y: 540}},
			Expect:   image.Rectangle{Max: image.Point{X: 3840, Y: 1080}},
		},
		{
			Name:     "3840x2160 -> 1920x1080",
			Bounding: baseRect,
			ToScale:  image.Rectangle{Max: image.Point{X: 3840, Y: 2160}},
			Expect:   image.Rectangle{Max: image.Point{X: 1920, Y: 1080}},
		},
		{
			Name:     "3840x4320 -> 1920x2160",
			Bounding: baseRect,
			ToScale:  image.Rectangle{Max: image.Point{X: 3840, Y: 4320}},
			Expect:   image.Rectangle{Max: image.Point{X: 1920, Y: 2160}},
		},
		{
			Name:     "7680x2160 -> 3840x1080",
			Bounding: baseRect,
			ToScale:  image.Rectangle{Max: image.Point{X: 7680, Y: 2160}},
			Expect:   image.Rectangle{Max: image.Point{X: 3840, Y: 1080}},
		},
		{
			Name:     "960x2160 -> 1920x4320",
			Bounding: baseRect,
			ToScale:  image.Rectangle{Max: image.Point{X: 960, Y: 2160}},
			Expect:   image.Rectangle{Max: image.Point{X: 1920, Y: 4320}},
		},
	}

	for _, entry := range testTable {
		entry := entry
		t.Run(entry.Name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, entry.Expect, ScaleRect(entry.ToScale, entry.Bounding))
		})
	}
}

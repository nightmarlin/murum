package layout

import (
	"image"
	"image/color"
	"testing"

	isPkg "github.com/matryer/is"
	"github.com/stretchr/testify/assert"
)

func TestL_addClosestPixel(t *testing.T) {
	// implicitly tests Region.addPixel
	testTable := []struct {
		Name string

		ToAddTo L

		Reference image.Point
		ToAdd     image.Point

		Expect L
	}{
		{
			Name:      "cannot add to L without regions",
			ToAddTo:   L{},
			Reference: image.Point{X: 1, Y: 1},
			ToAdd:     image.Point{X: 2, Y: 2},
			Expect:    L{},
		},
		{
			Name: "cannot add to L without a center that matches reference",
			ToAddTo: L{
				{center: image.Point{}},
			},
			Reference: image.Point{X: 1, Y: 1},
			ToAdd:     image.Point{X: 2, Y: 2},
			Expect: L{
				{center: image.Point{}},
			},
		},
		{
			Name: "successfully adds matching point to L",
			ToAddTo: L{
				{center: image.Point{}},
			},
			Reference: image.Point{},
			ToAdd:     image.Point{X: 2, Y: 2},
			Expect: L{
				{
					center: image.Point{},
					pixels: []image.Point{{X: 2, Y: 2}},
				},
			},
		},
	}

	for _, entry := range testTable {
		entry := entry
		t.Run(entry.Name, func(t *testing.T) {
			t.Parallel()

			entry.ToAddTo.addClosestPixel(entry.Reference, entry.ToAdd)
			assert.Equal(t, entry.Expect, entry.ToAddTo)
		})
	}
}

func TestRegion_Center(t *testing.T) {
	is := isPkg.New(t)

	r := Region{center: image.Point{X: 1, Y: 1}}
	expect := image.Point{X: 1, Y: 1}
	is.Equal(expect, r.Center())
}

func TestRegion_Pixels(t *testing.T) {
	is := isPkg.New(t)

	r := Region{pixels: []image.Point{{}, {}}}
	expect := []image.Point{{}, {}}
	is.Equal(expect, r.Pixels())
}

func TestRegion_Bounds(t *testing.T) {
	testTable := []struct {
		Name   string
		Pixels []image.Point
		Expect image.Rectangle
	}{
		{
			Name:   "nil slice returns empty rectangle",
			Pixels: nil,
			Expect: image.Rectangle{},
		},
		{
			Name:   "empty slice returns empty rectangle",
			Pixels: []image.Point{},
			Expect: image.Rectangle{},
		},
		{
			Name:   "1-unit slice returns 1-unit rectangle",
			Pixels: []image.Point{{X: 1, Y: 1}},
			Expect: image.Rectangle{
				Min: image.Point{X: 1, Y: 1},
				Max: image.Point{X: 1, Y: 1},
			},
		},
		{
			Name:   "2x2 region (in order)",
			Pixels: []image.Point{{X: 0, Y: 0}, {X: 1, Y: 1}},
			Expect: image.Rectangle{
				Min: image.Point{X: 0, Y: 0},
				Max: image.Point{X: 1, Y: 1},
			},
		},
		{
			Name:   "2x2 region (out of order)",
			Pixels: []image.Point{{X: 0, Y: 1}, {X: 1, Y: 0}},
			Expect: image.Rectangle{
				Min: image.Point{X: 0, Y: 0},
				Max: image.Point{X: 1, Y: 1},
			},
		},
		{
			Name: "3x3 region with redundant point",
			Pixels: []image.Point{
				{X: 0, Y: 0},
				{X: 1, Y: 1},
				{X: 2, Y: 2},
			},
			Expect: image.Rectangle{
				Min: image.Point{X: 0, Y: 0},
				Max: image.Point{X: 2, Y: 2},
			},
		},
	}

	for _, entry := range testTable {
		entry := entry
		t.Run(entry.Name, func(t *testing.T) {
			t.Parallel()

			r := Region{pixels: entry.Pixels}
			assert.Equal(t, entry.Expect, r.Bounds())
		})
	}
}

func TestRegion_At(t *testing.T) {
	testTable := []struct {
		Name string

		Pixels []image.Point
		Target image.Point
		Expect color.Color
	}{
		{
			Name:   "transparent for points in nil slice",
			Pixels: nil,
			Target: image.Point{X: 1, Y: 1},
			Expect: color.Transparent,
		},
		{
			Name:   "transparent for points in empty slice",
			Pixels: []image.Point{},
			Target: image.Point{X: 1, Y: 1},
			Expect: color.Transparent,
		},
		{
			Name:   "opaque for point in slice",
			Pixels: []image.Point{{X: 1, Y: 1}},
			Target: image.Point{X: 1, Y: 1},
			Expect: color.Opaque,
		},
		{
			Name:   "transparent for point not in slice",
			Pixels: []image.Point{{X: 0, Y: 0}},
			Target: image.Point{X: 1, Y: 1},
			Expect: color.Transparent,
		},
	}

	for _, entry := range testTable {
		entry := entry
		t.Run(entry.Name, func(t *testing.T) {
			t.Parallel()

			r := Region{pixels: entry.Pixels}
			c := r.At(entry.Target.X, entry.Target.Y)
			assert.Equal(t, entry.Expect, c)
		})
	}
}

func TestRegion_ColorModel(t *testing.T) {
	r := Region{}
	assert.Equal(t, color.RGBAModel, r.ColorModel())
}

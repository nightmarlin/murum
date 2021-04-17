package layout

import (
	"image"
	"image/color"
	"sync"
)

// L defines a Layout - a slice of Region s, where each region is the set of points closest to the reference point
// (which can be found using Region.Center). All pointer receivers on L will panic if called on a nil value
type L []Region

// addClosestPixel will find the
func (l *L) addClosestPixel(ref image.Point, p image.Point) {
	// While clunky, using (l *L) shows that this pointer receiver modifies the caller!

	for i := range *l {
		if (*l)[i].center.Eq(ref) {
			(*l)[i].addPixel(p)
		}
	}
}

// A Region is defined as the set of image.Point s closest to a reference central image.Point. It implements
// image.Image. All pointer receivers on Region will panic if called on a nil value
type Region struct {
	// mux provides thread safety, but is only needed internally
	mux sync.Mutex

	// center should be set at creation and left unchanged
	center image.Point

	// pixels is the slice of all pixels closest to center
	pixels []image.Point
}

// Center returns the reference point that the Pixels in this Region are closest to
func (r *Region) Center() image.Point {
	r.mux.Lock()
	defer r.mux.Unlock()

	return r.center
}

// Pixels returns the set of pixels closest to the reference point at Center
func (r *Region) Pixels() []image.Point {
	r.mux.Lock()
	defer r.mux.Unlock()
	if r.pixels == nil {
		return nil
	}

	c := make([]image.Point, len(r.pixels))
	copy(c, r.pixels)
	return c
}

// addPixel adds a pixel to this Region's pixels slice
func (r *Region) addPixel(p image.Point) {
	r.mux.Lock()
	defer r.mux.Unlock()
	r.pixels = append(r.pixels, p)
}

// Bounds implements image.Image. Bounds gets the bounding rectangle of this Region's pixels. The Bounds are not
// necessarily fully-filled
func (r *Region) Bounds() image.Rectangle {
	r.mux.Lock()
	defer r.mux.Unlock()

	if len(r.pixels) == 0 {
		return image.Rectangle{}
	}

	minX, minY := r.pixels[0].X, r.pixels[0].Y
	maxX, maxY := minX, minY

	for i := 1; i < len(r.pixels); i++ {
		p := r.pixels[i]

		if minY > p.Y {
			minY = p.Y
		} else if maxY < p.Y {
			maxY = p.Y
		}

		if minX > p.X {
			minX = p.X
		} else if maxX < p.X {
			maxX = p.X
		}
	}

	return image.Rect(minX, minY, maxX, maxY)
}

// At implements image.Image. If a pixel is within this Region, color.Opaque is returned. Otherwise color.Transparent is
// returned. This allows a Region to be used as a mask
func (r *Region) At(x, y int) color.Color {
	r.mux.Lock()
	defer r.mux.Unlock()

	p := image.Point{X: x, Y: y}
	for i := range r.pixels {
		if r.pixels[i].Eq(p) {
			return color.Opaque
		}
	}
	return color.Transparent
}

// ColorModel implements image.Image. murum uses the color.RGBAModel by default
func (r *Region) ColorModel() color.Model {
	return color.RGBAModel
}

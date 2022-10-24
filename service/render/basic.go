package render

import (
	"errors"
	"fmt"
	"image"

	"github.com/nfnt/resize"

	layout "github.com/nightmarlin/murum/protos/layout/v1"
	"github.com/nightmarlin/murum/provider"
)

// BasicRendererOption allows changes to the basic renderer.
type BasicRendererOption func(b *BasicRenderer) error

// BasicWithInterpolationFunc sets the interpolation function used by the basic renderer to resize
// images to fit the bounding rectangle of a layout.L region
func BasicWithInterpolationFunc(intFunc resize.InterpolationFunction) BasicRendererOption {
	return func(b *BasicRenderer) error {
		b.interpolationFunc = intFunc
		return nil
	}
}

// BasicWithAspectRatioIgnored allows images to be scaled to fit the bounding rectangle of a
// layout.L region without preserving their aspect ratio. This reduces the number of calculations
// done at the cost of skewing images
func BasicWithAspectRatioIgnored() BasicRendererOption {
	return func(b *BasicRenderer) error {
		b.ignoreAspectRatio = true
		return nil
	}
}

// NewBasic creates and returns a new BasicRenderer that implements the Renderer interface. The
// renderer simply fills in regions of the layout.L with images passed to it and returns the image.
// If there aren't enough images, some sections of the returned image will have the default color.
// Note: If BasicWithInterpolationFunc is not passed as an option, the BasicRenderer will use the
// resize.Lanczos3 algorithm to resize images to fit the bounding rectangles.
func NewBasic(opts ...BasicRendererOption) (BasicRenderer, error) {
	res := BasicRenderer{interpolationFunc: resize.Lanczos3}
	for i := range opts {
		err := opts[i](&res)
		if err != nil {
			return BasicRenderer{}, fmt.Errorf("failed to init basic renderer: %w", err)
		}
	}
	return res, nil
}

type BasicRenderer struct {
	interpolationFunc resize.InterpolationFunction
	ignoreAspectRatio bool
}

func (b *BasicRenderer) resize(i image.Image, bound image.Rectangle) image.Image {
	if b.ignoreAspectRatio {
		return resize.Resize(uint(bound.Dx()), uint(bound.Dy()), i, b.interpolationFunc)
	}

	scaledRect := ScaleRect(i.Bounds(), bound)
	return resize.Resize(uint(scaledRect.Dx()), uint(scaledRect.Dy()), i, b.interpolationFunc)
}

func (b *BasicRenderer) Render(l *layout.Layout, ai []provider.AlbumInfo) (image.Image, error) {
	if l == nil {
		return nil, errors.New("cannot render from nil layout")
	}

	var (
		resImg = image.NewRGBA64(image.Rectangle{
			Max: image.Point{
				X: int(l.Bounds.Width),
				Y: int(l.Bounds.Height),
			},
		})
		n = 0
	)

	for _, section := range l.Regions {
		if len(ai) <= n {
			break
		}

		var (
			sectBounds = GetBounds(section.Members)
			sectImg    = b.resize(ai[n].Art, sectBounds)

			// Center image
			xCenteredOffset = (sectImg.Bounds().Dx() - sectBounds.Dx()) / 2
			yCenteredOffset = (sectImg.Bounds().Dy() - sectBounds.Dy()) / 2
		)

		for _, p := range section.Members {
			col := sectImg.At(
				int(p.X)-sectBounds.Min.X+xCenteredOffset,
				int(p.Y)-sectBounds.Min.Y+yCenteredOffset,
			)
			resImg.Set(int(p.X), int(p.Y), col)
		}

		n += 1
	}

	return resImg, nil
}

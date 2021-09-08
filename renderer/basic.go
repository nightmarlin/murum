package renderer

import (
	"errors"
	"fmt"
	"image"

	"github.com/nfnt/resize"

	"github.com/nightmarlin/murum/layout"
	"github.com/nightmarlin/murum/provider"
)

// BasicRendererOption allows changes to the basic renderer.
type BasicRendererOption func(b *basicRenderer) error

// BasicWithInterpolationFunc sets the interpolation function used by the basic renderer to resize
// images to fit the bounding rectangle of a layout.L region
func BasicWithInterpolationFunc(intFunc resize.InterpolationFunction) BasicRendererOption {
	return func(b *basicRenderer) error {
		b.interpolationFunc = intFunc
		return nil
	}
}

// BasicWithAspectRatioIgnored allows images to be scaled to fit the bounding rectangle of a
// layout.L region without preserving their aspect ratio. This reduces the number of calculations
// done at the cost of skewing images
func BasicWithAspectRatioIgnored() BasicRendererOption {
	return func(b *basicRenderer) error {
		b.ignoreAspectRatio = true
		return nil
	}
}

// NewBasic creates and returns a new basicRenderer that implements the Renderer interface. The
// renderer simply fills in regions of the layout.L with images passed to it and returns the image.
// If there aren't enough images, some sections of the returned image will have the default color.
// Note: If BasicWithInterpolationFunc is not passed as an option, the basicRenderer will use the
// resize.Lanczos3 algorithm to resize images to fit the bounding rectangles.
func NewBasic(opts ...BasicRendererOption) (*basicRenderer, error) {
	res := &basicRenderer{interpolationFunc: resize.Lanczos3}
	for i := range opts {
		err := opts[i](res)
		if err != nil {
			return nil, fmt.Errorf("failed to init basic renderer: %w", err)
		}
	}
	return res, nil
}

type basicRenderer struct {
	interpolationFunc resize.InterpolationFunction
	ignoreAspectRatio bool
}

func (b *basicRenderer) resize(i image.Image, bound image.Rectangle) image.Image {
	if b.ignoreAspectRatio {
		return resize.Resize(uint(bound.Dx()), uint(bound.Dy()), i, b.interpolationFunc)
	}

	scaledRect := ScaleRect(i.Bounds(), bound)
	return resize.Resize(uint(scaledRect.Dx()), uint(scaledRect.Dy()), i, b.interpolationFunc)
}

func (b *basicRenderer) Render(l *layout.L, ai []provider.AlbumInfo) (image.Image, error) {
	if l == nil {
		return nil, errors.New("cannot render from nil layout")
	}

	var (
		resImg = image.NewRGBA64(l.Bounds)
		n      = 0
	)

	for _, section := range l.Points {
		if len(ai) <= n {
			break
		}

		var (
			sectBounds = GetBounds(section)
			sectImg    = b.resize(ai[n].Art, sectBounds)

			// TODO: Center image correctly
			// xCenteredOffset = (sectBounds.Dx() - sectImg.Bounds().Dx()) / 2 // (BoundsDX - ImgDX)/2
			// yCenteredOffset = (sectBounds.Dy() - sectImg.Bounds().Dy()) / 2
		)
		n++

		for _, p := range section {
			// TODO: Ensure image correctly fills region.
			col := sectImg.At(
				p.X-sectBounds.Min.X, // +xCenteredOffset,
				p.Y-sectBounds.Min.Y, // +yCenteredOffset,
			)
			resImg.Set(p.X, p.Y, col)
		}
	}

	return resImg, nil
}

package layout

import (
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestL_addClosestPixel(t *testing.T) {
	testTable := []struct {
		Name string

		ToAddTo L

		Reference image.Point
		ToAdd     image.Point

		Expect L
	}{
		{}, // TODO: This
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

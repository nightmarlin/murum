package provider

import (
	"image"
	"time"
)

type AlbumInfo struct {
	Name       string      // The name of the album
	Artist     string      // Who made the album
	Art        image.Image // The album cover art
	PlayCount  *int        // How many times songs from this album have been played
	LastPlayed *time.Time  // When this album was played most recently
}

// Provider retrieves listening history data.
type Provider interface {
	// Fetch retrieves the first n AlbumInfo entities that the Provider can provide.
	Fetch(n int) ([]AlbumInfo, error)
}

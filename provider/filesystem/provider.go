package filesystem

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/nightmarlin/murum/provider"
)

type FSOption func(fp *fsProvider) error

// WithFilter adds a regex filter to the file search. Only files that match *any* provided filter
// will be used.
func WithFilter(r regexp.Regexp) FSOption {
	return func(fp *fsProvider) error {
		fp.filters = append(fp.filters, r)
		return nil
	}
}

// WithWorkingDir sets the root directory to search in
func WithWorkingDir(path string) FSOption {
	return func(fp *fsProvider) error {
		p, err := filepath.Abs(path)
		if err != nil {
			return fmt.Errorf("failed to create absolute path: %w", err)
		}

		info, err := os.Stat(p)
		if err != nil {
			return fmt.Errorf("unable to stat file path: %w", err)
		}

		if !info.IsDir() {
			return errors.New("file path must be a directory")
		}

		fp.rootDir = p
		return nil
	}
}

// WithExtensions adds supported extensions. Note that for an image extension to be decode-able it
// must be imported (for side effects) by a built package. The following extensions are supported by
// default, but must be enabled individually:
//   - PNG
//   - JPG
//   - GIF
func WithExtensions(ext ...string) FSOption {
	return func(fp *fsProvider) error {
		fp.extensions = append(fp.extensions, ext...)
		return nil
	}
}

// WithDefaults adds the JPG, JPEG and PNG extensions and sets the working directory to the CWD.
func WithDefaults() FSOption {
	return func(fp *fsProvider) error {
		_ = WithExtensions("jpg", "jpeg", "png")

		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("failed to get cwd: %w", err)
		}
		fp.rootDir = cwd

		return nil
	}
}

type fsProvider struct {
	// Only files that match any of the provided extensions will be used
	extensions []string
	// If present, only file names (excluding extensions) that match any of these filters will be used
	// (on top of the standard filters)
	filters []regexp.Regexp
	// The directory to search in
	rootDir string
}

// New returns a provider.Provider that gets provider.AlbumInfo from a local directory. Images in
// that directory with names of the form "Artist Name - Album Name" will be converted into the
// corresponding information structure. If WithWorkingDir isn't provided, the current working
// directory will be used.
func New(opts ...FSOption) (*fsProvider, error) {
	fp := &fsProvider{}
	for i := range opts {
		err := opts[i](fp)
		if err != nil {
			return nil, fmt.Errorf("failed to initialise filesystem provider: %w", err)
		}
	}

	// Fill in CWD if root directory not set
	if fp.rootDir == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return nil, fmt.Errorf("failed to get cwd: %w", err)
		}
		fp.rootDir = cwd
	}

	return fp, nil
}

func (fp *fsProvider) listFiles() ([]string, error) {
	// Read all files in directory
	dInfo, err := os.ReadDir(fp.rootDir)
	if err != nil {
		return nil, fmt.Errorf("unable to read directory: %w", err)
	}

	names := make([]string, 0)
	for _, i := range dInfo {
		if i.IsDir() { // skip directories
			continue
		}

		fName := i.Name()
		ext, allowed := filepath.Ext(fName), false
		for _, knownExt := range fp.extensions {
			if knownExt == ext {
				allowed = true
				break
			}
		}
		if !allowed { // only support specified extensions
			continue
		}

		if len(fp.filters) != 0 {
			nameNoExt, allowed := strings.TrimSuffix(fName, ext), false
			for _, exp := range fp.filters {
				if exp.MatchString(nameNoExt) {
					allowed = true
					break
				}
			}
			if !allowed { // only allow files that match a filter
				continue
			}
		}

		names = append(names, filepath.Join(fp.rootDir, fName))
	}

	return names, nil
}

func extractFromFile(path string) (provider.AlbumInfo, error) {
	imgBytes, err := os.ReadFile(path)
	if err != nil {
		return provider.AlbumInfo{}, fmt.Errorf("unable to open file: %w", err)
	}

	img, _, err := image.Decode(bytes.NewBuffer(imgBytes))
	if err != nil {
		return provider.AlbumInfo{}, fmt.Errorf("unable to decode image: %w", err)
	}

	pathInfo := strings.SplitN(path, " - ", 2)
	artist := pathInfo[0]
	albumName := ""
	if len(pathInfo) > 1 {
		albumName = pathInfo[1]
	}

	return provider.AlbumInfo{Name: albumName, Artist: artist, Art: img}, nil
}

func (fp *fsProvider) Fetch(n int) ([]provider.AlbumInfo, error) {
	if n <= 0 {
		return nil, nil
	}

	allFiles, err := fp.listFiles()
	if err != nil {
		return nil, err
	}

	toRetrieve := int(math.Min(float64(n), float64(len(allFiles))))
	res := make([]provider.AlbumInfo, toRetrieve)
	for i := 0; i < toRetrieve; i++ {
		ai, err := extractFromFile(allFiles[i])
		if err != nil {
			return nil, err
		}
		res[i] = ai
	}

	return res, nil
}

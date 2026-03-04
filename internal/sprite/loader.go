package sprite

import (
	"fmt"
	"image/png"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	spriteURLTemplate = "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/%d.png"
	cacheDir          = "assets/cache"
)

// LoadSprite downloads a Pokemon sprite by ID, caching it to disk.
// Returns an *ebiten.Image ready for rendering.
func LoadSprite(id int) (*ebiten.Image, error) {
	if err := os.MkdirAll(cacheDir, 0o755); err != nil {
		return nil, fmt.Errorf("create cache dir: %w", err)
	}

	cachePath := filepath.Join(cacheDir, fmt.Sprintf("%d.png", id))

	// Try loading from cache first.
	if img, err := loadFromDisk(cachePath); err == nil {
		return img, nil
	}

	// Download from PokeAPI sprites repo.
	url := fmt.Sprintf(spriteURLTemplate, id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("download sprite %d: %w", id, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("download sprite %d: status %d", id, resp.StatusCode)
	}

	// Save to cache.
	f, err := os.Create(cachePath)
	if err != nil {
		return nil, fmt.Errorf("create cache file: %w", err)
	}

	if _, err := io.Copy(f, resp.Body); err != nil {
		f.Close()
		return nil, fmt.Errorf("write cache file: %w", err)
	}
	f.Close()

	return loadFromDisk(cachePath)
}

func loadFromDisk(path string) (*ebiten.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		return nil, fmt.Errorf("decode png: %w", err)
	}

	return ebiten.NewImageFromImage(img), nil
}

package sprite

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const Scale = 3.0

// Sprite holds position, velocity, and the image for a Pokemon sprite.
type Sprite struct {
	X, Y   float64
	VX, VY float64
	Image  *ebiten.Image
}

// New creates a Sprite at the given position with a default velocity.
func New(img *ebiten.Image, x, y float64) *Sprite {
	return &Sprite{
		X:     x,
		Y:     y,
		VX:    2,
		VY:    1.5,
		Image: img,
	}
}

// Size returns the scaled width and height of the sprite.
func (s *Sprite) Size() (float64, float64) {
	b := s.Image.Bounds()
	return float64(b.Dx()) * Scale, float64(b.Dy()) * Scale
}

// Update moves the sprite and bounces off screen edges.
func (s *Sprite) Update(screenW, screenH int) {
	s.X += s.VX
	s.Y += s.VY

	w, h := s.Size()

	if s.X < 0 {
		s.X = 0
		s.VX = -s.VX
	} else if s.X+w > float64(screenW) {
		s.X = float64(screenW) - w
		s.VX = -s.VX
	}

	if s.Y < 0 {
		s.Y = 0
		s.VY = -s.VY
	} else if s.Y+h > float64(screenH) {
		s.Y = float64(screenH) - h
		s.VY = -s.VY
	}
}

// Draw renders the sprite on the screen at its current position.
func (s *Sprite) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(Scale, Scale)
	op.GeoM.Translate(s.X, s.Y)
	screen.DrawImage(s.Image, op)
}

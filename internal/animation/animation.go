package animation

import (
	"math"

	"github.com/dario/pokesprite/internal/sprite"
)

// Animator applies dance/idle behaviors to a sprite.
type Animator struct {
	frame   int
	baseY   float64
	started bool
}

// New creates a new Animator.
func New() *Animator {
	return &Animator{}
}

// Update advances the animation by one frame, applying idle bobbing
// on top of the sprite's bounce movement.
func (a *Animator) Update(s *sprite.Sprite, screenW, screenH int) {
	if !a.started {
		a.baseY = s.Y
		a.started = true
	}

	// Advance the bounce movement.
	s.Update(screenW, screenH)

	// Apply idle bobbing: a sine-wave offset on the Y axis.
	a.frame++
	bob := math.Sin(float64(a.frame)/15.0) * 4.0
	s.Y += bob

	// Track base Y for next frame (strip the bob we just added).
	a.baseY = s.Y - bob
}

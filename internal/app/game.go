package app

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/dario/pokesprite/internal/animation"
	"github.com/dario/pokesprite/internal/sprite"
)

const (
	ScreenWidth  = 480
	ScreenHeight = 480
)

// Game implements ebiten.Game for the PokeSprite overlay.
type Game struct {
	sprite   *sprite.Sprite
	animator *animation.Animator
}

// NewGame creates a Game with the given sprite image.
func NewGame(img *ebiten.Image) *Game {
	s := sprite.New(img, 100, 100)
	return &Game{
		sprite:   s,
		animator: animation.New(),
	}
}

func (g *Game) Update() error {
	g.animator.Update(g.sprite, ScreenWidth, ScreenHeight)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 0}) // transparent background
	g.sprite.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

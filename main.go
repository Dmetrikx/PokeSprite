package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/dario/pokesprite/internal/app"
	"github.com/dario/pokesprite/internal/pokemon"
	"github.com/dario/pokesprite/internal/sprite"
)

func main() {
	chosen := pickStarter()

	fmt.Printf("Loading %s...\n", chosen.Name)
	img, err := sprite.LoadSprite(chosen.ID)
	if err != nil {
		log.Fatalf("Failed to load sprite: %v", err)
	}

	game := app.NewGame(img)

	ebiten.SetWindowSize(app.ScreenWidth, app.ScreenHeight)
	ebiten.SetWindowTitle("PokeSprite - " + chosen.Name)
	ebiten.SetWindowDecorated(false)
	ebiten.SetWindowFloating(true)

	if err := ebiten.RunGameWithOptions(game, &ebiten.RunGameOptions{
		ScreenTransparent: true,
	}); err != nil {
		log.Fatal(err)
	}
}

func pickStarter() pokemon.Pokemon {
	starters := pokemon.Starters()

	fmt.Println("Pick your starter Pokemon:")
	for i, p := range starters {
		fmt.Printf("  %d. %s\n", i+1, p.Name)
	}
	fmt.Print("Enter choice (1-3): ")

	var choice int
	if _, err := fmt.Scan(&choice); err != nil || choice < 1 || choice > 3 {
		fmt.Println("Invalid choice, defaulting to Bulbasaur.")
		return starters[0]
	}

	selected := starters[choice-1]
	fmt.Printf("You chose %s!\n", selected.Name)
	os.Stdout.Sync()
	return selected
}

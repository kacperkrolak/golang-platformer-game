package main

import (
	_ "image/png"
	"kacperkrolak/golang-platformer-game/pkg/game"
	"log"
	"os"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	projectRoot, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	g := game.MakeGame(
		filepath.Join(projectRoot, "assets", "maps", "1.txt"),
		filepath.Join(projectRoot, "assets", "tiles", "sheet.png"),
		filepath.Join(projectRoot, "assets", "tiles", "characters.png"),
	)

	ebiten.SetWindowSize(game.ScreenWidth*2, game.ScreenHeight*2)
	ebiten.SetWindowTitle("Tiles (Ebitengine Demo)")
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}

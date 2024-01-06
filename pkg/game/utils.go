package game

import (
	"image"
	"kacperkrolak/golang-platformer-game/pkg/gamemap"
	"kacperkrolak/golang-platformer-game/pkg/gamemap/parser"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func loadGameMap(mapFile string, tileSize int) (gamemap.Map, error) {
	file, err := os.Open(mapFile)
	if err != nil {
		return gamemap.Map{}, err
	}
	defer file.Close()

	parser := parser.MapDataParser{}
	tiles, err := parser.LoadTiles(file)
	if err != nil {
		return gamemap.Map{}, err
	}

	gameMap := gamemap.MakeMap(tiles, tileSize)

	return gameMap, nil
}

func loadTilesImage(textureFile string) (*ebiten.Image, error) {
	// Decode an image from the image file's byte slice.
	file, err := os.Open(textureFile)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return ebiten.NewImageFromImage(img), nil
}

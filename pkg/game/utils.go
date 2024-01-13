package game

import (
	"image"
	"kacperkrolak/golang-platformer-game/pkg/world/gamemap"
	"kacperkrolak/golang-platformer-game/pkg/world/parser"
	"kacperkrolak/golang-platformer-game/pkg/world/tilemap"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func loadWorldMap(mapFile string, tileSize int) (gamemap.Map, tilemap.Map, error) {
	file, err := os.Open(mapFile)
	if err != nil {
		return gamemap.Map{}, tilemap.Map{}, err
	}
	defer file.Close()

	parser := parser.Parser{}
	blocks, tiles, err := parser.Load(file)
	if err != nil {
		return gamemap.Map{}, tilemap.Map{}, err
	}

	gameMap := gamemap.MakeMap(blocks, tileSize)
	tileMap := tilemap.MakeMap(tiles, tileSize)

	return gameMap, tileMap, nil
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

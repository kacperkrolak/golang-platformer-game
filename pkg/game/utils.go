package game

import (
	"image"
	"kacperkrolak/golang-platformer-game/pkg/parser"
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"
	"kacperkrolak/golang-platformer-game/pkg/world/gamemap"
	"kacperkrolak/golang-platformer-game/pkg/world/tilemap"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func loadWorldMap(mapFile string, tileSize int, tileSheet *ebiten.Image) (gamemap.Map, tilemap.Map, []vector.Vector2, vector.Vector2, error) {
	file, err := os.Open(mapFile)
	if err != nil {
		return gamemap.Map{}, tilemap.Map{}, []vector.Vector2{}, vector.Vector2{}, err
	}
	defer file.Close()

	parser := parser.Parser{
		TileSheet: tileSheet,
		TileSize:  tileSize,
	}

	parsedData, err := parser.Load(file)
	if err != nil {
		return gamemap.Map{}, tilemap.Map{}, []vector.Vector2{}, vector.Vector2{}, err
	}

	gameMap := gamemap.MakeMap(parsedData.Blocks, tileSize)
	tileMap := tilemap.MakeMap(parsedData.Tiles, tileSize)

	return gameMap, tileMap, parsedData.CoinPositions, parsedData.SpawnPoint, nil
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

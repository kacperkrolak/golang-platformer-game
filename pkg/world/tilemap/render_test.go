package tilemap

import (
	"testing"

	"kacperkrolak/golang-platformer-game/pkg/physics/vector"
	"kacperkrolak/golang-platformer-game/pkg/world/tilemap/tile"
	tilemock "kacperkrolak/golang-platformer-game/pkg/world/tilemap/tile/mock"

	"github.com/hajimehoshi/ebiten/v2"
)

func TestMap_Draw(t *testing.T) {
	screen := &ebiten.Image{}
	cameraOffset := vector.Vector2{}
	img := &ebiten.Image{}
	tileSize := 32

	tiles := [][]tile.Tile{
		{&tilemock.MockTile{}, &tilemock.MockTile{}},
		{&tilemock.MockTile{}, &tilemock.MockTile{}},
		{&tilemock.MockTile{}, &tilemock.MockTile{}},
	}

	for i, t := range tiles {
		for j, tile := range t {
			tile.(*tilemock.MockTile).On("Draw", screen, cameraOffset.Added(vector.Vector2{float64(j * tileSize), float64(i * tileSize)}), img, tileSize).Return()
		}
	}

	testMap := Map{
		Tiles: tiles,
	}

	testMap.Draw(screen, cameraOffset, img, tileSize)
}

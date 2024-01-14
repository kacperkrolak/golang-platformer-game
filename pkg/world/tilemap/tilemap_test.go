package tilemap

import (
	"kacperkrolak/golang-platformer-game/pkg/world/tilemap/tile"
	tilemock "kacperkrolak/golang-platformer-game/pkg/world/tilemap/tile/mock"
	"testing"

	"github.com/stretchr/testify/mock"
)

// Assert every tile is updated.
func TestMap_FindVariants(t *testing.T) {
	testMap := Map{
		Tiles: [][]tile.Tile{
			{&tilemock.MockTile{}, &tilemock.MockTile{}},
			{&tilemock.MockTile{}, &tilemock.MockTile{}},
		},
	}

	for _, row := range testMap.Tiles {
		for _, tile := range row {
			tile.(*tilemock.MockTile).On("UpdateVariant", mock.Anything).Return()
		}
	}

	testMap.FindVariants()
}

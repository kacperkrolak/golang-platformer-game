package gamemap

import (
	"kacperkrolak/golang-platformer-game/pkg/gamemap/tile"
	tilemock "kacperkrolak/golang-platformer-game/pkg/gamemap/tile/mock"
	"kacperkrolak/golang-platformer-game/pkg/physics/box"
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestMap_CreateHitboxes(t *testing.T) {
	m := Map{
		Tiles: [][]tile.Tile{
			{tilemock.NewEmptyTile(), tilemock.NewEmptyTile(), tilemock.NewEmptyTile()},
			{tilemock.NewCollidableTile(), tilemock.NewCollidableTile(), tilemock.NewCollidableTile()},
			{tilemock.NewEmptyTile(), tilemock.NewCollidableTile(), tilemock.NewEmptyTile()},
		},
	}

	tileSize := 10
	m.CreateHitboxes(tileSize)

	expectedHitboxes := [][]box.Box{
		{box.Box{}, box.Box{}, box.Box{}},
		{{Position: vector.Vector2{X: 0, Y: 10}, Size: vector.Vector2{X: 10, Y: 10}}, {Position: vector.Vector2{X: 10, Y: 10}, Size: vector.Vector2{X: 10, Y: 10}}, {Position: vector.Vector2{X: 20, Y: 10}, Size: vector.Vector2{X: 10, Y: 10}}},
		{box.Box{}, {Position: vector.Vector2{X: 10, Y: 20}, Size: vector.Vector2{X: 10, Y: 10}}, box.Box{}},
	}

	for i, hitboxes := range expectedHitboxes {
		for j, hitbox := range hitboxes {
			if m.Tiles[i][j].Hitbox() != hitbox {
				t.Errorf("Expected hitbox %v, got %v", hitbox, m.Tiles[i][j].Hitbox())
			}
		}
	}
}

func TestMap_FindVariants(t *testing.T) {
	tiles := [][]*tilemock.MockTile{
		{tilemock.NewEmptyTile(), tilemock.NewEmptyTile(), tilemock.NewEmptyTile()},
		{tilemock.NewCollidableTile(), tilemock.NewCollidableTile(), tilemock.NewCollidableTile()},
		{tilemock.NewEmptyTile(), tilemock.NewCollidableTile(), tilemock.NewEmptyTile()},
	}

	var mapTiles [][]tile.Tile
	for _, row := range tiles {
		var mapRow []tile.Tile
		for i := 0; i < len(row); i++ {
			row[i].On("UpdateVariant", mock.Anything).Return()
			mapRow = append(mapRow, row[i])
		}
		mapTiles = append(mapTiles, mapRow)
	}

	m := Map{
		Tiles: mapTiles,
	}

	m.FindVariants()

	for _, row := range tiles {
		for i := 0; i < len(row); i++ {
			row[i].AssertExpectations(t)
		}
	}
}
func TestMakeMap(t *testing.T) {
	tiles := [][]tile.Tile{
		{tilemock.NewEmptyTile(), tilemock.NewEmptyTile(), tilemock.NewEmptyTile()},
		{tilemock.NewCollidableTile(), tilemock.NewCollidableTile(), tilemock.NewCollidableTile()},
		{tilemock.NewEmptyTile(), tilemock.NewCollidableTile(), tilemock.NewEmptyTile()},
	}
	tileSize := 10

	expectedMap := Map{
		Tiles: [][]tile.Tile{
			{tilemock.NewEmptyTile(), tilemock.NewEmptyTile(), tilemock.NewEmptyTile()},
			{tilemock.NewCollidableTile(), tilemock.NewCollidableTile(), tilemock.NewCollidableTile()},
			{tilemock.NewEmptyTile(), tilemock.NewCollidableTile(), tilemock.NewEmptyTile()},
		},
	}

	expectedMap.FindVariants()
	expectedMap.CreateHitboxes(tileSize)

	result := MakeMap(tiles, tileSize)

	// Check if the generated map matches the expected map
	if !reflect.DeepEqual(result, expectedMap) {
		t.Errorf("Expected map %+v, got %+v", expectedMap, result)
	}
}

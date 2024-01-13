package gamemap

import (
	"kacperkrolak/golang-platformer-game/pkg/physics/box"
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"
	"kacperkrolak/golang-platformer-game/pkg/world/gamemap/block"
	blockmock "kacperkrolak/golang-platformer-game/pkg/world/gamemap/block/mock"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestMap_CreateHitboxes(t *testing.T) {
	m := Map{
		Blocks: [][]block.Block{
			{blockmock.NewEmptyTile(), blockmock.NewEmptyTile(), blockmock.NewEmptyTile()},
			{blockmock.NewCollidableTile(), blockmock.NewCollidableTile(), blockmock.NewCollidableTile()},
			{blockmock.NewEmptyTile(), blockmock.NewCollidableTile(), blockmock.NewEmptyTile()},
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
			if m.Blocks[i][j].Hitbox() != hitbox {
				t.Errorf("Expected hitbox %v, got %v", hitbox, m.Blocks[i][j].Hitbox())
			}
		}
	}
}

func TestMap_FindVariants(t *testing.T) {
	tiles := [][]*blockmock.MockBlock{
		{blockmock.NewEmptyTile(), blockmock.NewEmptyTile(), blockmock.NewEmptyTile()},
		{blockmock.NewCollidableTile(), blockmock.NewCollidableTile(), blockmock.NewCollidableTile()},
		{blockmock.NewEmptyTile(), blockmock.NewCollidableTile(), blockmock.NewEmptyTile()},
	}

	var mapBlocks [][]block.Block
	for _, row := range tiles {
		var mapRow []block.Block
		for i := 0; i < len(row); i++ {
			row[i].On("UpdateVariant", mock.Anything).Return()
			mapRow = append(mapRow, row[i])
		}
		mapBlocks = append(mapBlocks, mapRow)
	}

	m := Map{
		Blocks: mapBlocks,
	}

	m.FindVariants()

	for _, row := range tiles {
		for i := 0; i < len(row); i++ {
			row[i].AssertExpectations(t)
		}
	}
}
func TestMakeMap(t *testing.T) {
	tiles := [][]block.Block{
		{blockmock.NewEmptyTile(), blockmock.NewEmptyTile(), blockmock.NewEmptyTile()},
		{blockmock.NewCollidableTile(), blockmock.NewCollidableTile(), blockmock.NewCollidableTile()},
		{blockmock.NewEmptyTile(), blockmock.NewCollidableTile(), blockmock.NewEmptyTile()},
	}
	tileSize := 10

	expectedMap := Map{
		Blocks: [][]block.Block{
			{blockmock.NewEmptyTile(), blockmock.NewEmptyTile(), blockmock.NewEmptyTile()},
			{blockmock.NewCollidableTile(), blockmock.NewCollidableTile(), blockmock.NewCollidableTile()},
			{blockmock.NewEmptyTile(), blockmock.NewCollidableTile(), blockmock.NewEmptyTile()},
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

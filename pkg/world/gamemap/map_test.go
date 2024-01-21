package gamemap

import (
	"kacperkrolak/golang-platformer-game/pkg/physics/box"
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"
	"kacperkrolak/golang-platformer-game/pkg/world/gamemap/block"
	blockmock "kacperkrolak/golang-platformer-game/pkg/world/gamemap/block/mock"
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
			row[i].On("AdaptToNeighbours", mock.Anything).Return()
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

func TestMap_Height(t *testing.T) {
	m := Map{
		Blocks:   [][]block.Block{},
		TileSize: 10,
	}

	expectedHeight := 0
	if result := m.Height(); result != expectedHeight {
		t.Errorf("Expected height %d, got %d", expectedHeight, result)
	}

	m.Blocks = [][]block.Block{
		{blockmock.NewEmptyTile(), blockmock.NewEmptyTile(), blockmock.NewEmptyTile()},
		{blockmock.NewCollidableTile(), blockmock.NewCollidableTile(), blockmock.NewCollidableTile()},
		{blockmock.NewEmptyTile(), blockmock.NewCollidableTile(), blockmock.NewEmptyTile()},
	}

	expectedHeight = 30
	if result := m.Height(); result != expectedHeight {
		t.Errorf("Expected height %d, got %d", expectedHeight, result)
	}
}

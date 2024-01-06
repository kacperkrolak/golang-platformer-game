package gamemap

import (
	"kacperkrolak/golang-platformer-game/pkg/physics/box"
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"
	"reflect"
	"testing"
)

func TestMap_CreateHitboxes(t *testing.T) {
	m := Map{
		Tiles: [][]Tile{
			{{Type: EMPTY}, {Type: EMPTY}, {Type: EMPTY}},
			{{Type: DIRT}, {Type: DIRT}, {Type: DIRT}},
			{{Type: EMPTY}, {Type: DIRT}, {Type: EMPTY}},
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
			if m.Tiles[i][j].Hitbox != hitbox {
				t.Errorf("Expected hitbox %v, got %v", hitbox, m.Tiles[i][j].Hitbox)
			}
		}
	}
}

func TestMap_FindVariants(t *testing.T) {
	m := Map{
		Tiles: [][]Tile{
			{{Type: EMPTY}, {Type: EMPTY}, {Type: EMPTY}},
			{{Type: DIRT}, {Type: DIRT}, {Type: DIRT}},
			{{Type: EMPTY}, {Type: DIRT}, {Type: EMPTY}},
		},
	}

	m.FindVariants()

	// Empty tiles should have no variants.
	expectedTiles := [][]uint8{
		{0, 0, 0},
		{TOP_BOTTOM, TOP, TOP_BOTTOM},
		{0, LEFT_RIGHT, 0},
	}

	for i, row := range expectedTiles {
		for j, variant := range row {
			if m.Tiles[i][j].Variant != variant {
				t.Errorf("Expected tile %d:%d to have variant %+v, got %+v", j, i, variant, m.Tiles[i][j].Variant)
			}
		}
	}
}
func TestMakeMap(t *testing.T) {
	tiles := [][]Tile{
		{{Type: EMPTY}, {Type: EMPTY}, {Type: EMPTY}},
		{{Type: DIRT}, {Type: DIRT}, {Type: DIRT}},
		{{Type: EMPTY}, {Type: DIRT}, {Type: EMPTY}},
	}
	tileSize := 10

	expectedMap := Map{
		Tiles: [][]Tile{
			{{Type: EMPTY}, {Type: EMPTY}, {Type: EMPTY}},
			{{Type: DIRT}, {Type: DIRT}, {Type: DIRT}},
			{{Type: EMPTY}, {Type: DIRT}, {Type: EMPTY}},
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

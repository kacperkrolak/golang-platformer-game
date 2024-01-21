package parser

import (
	"reflect"
	"strings"
	"testing"

	"kacperkrolak/golang-platformer-game/pkg/physics/vector"
	"kacperkrolak/golang-platformer-game/pkg/world/block/empty"
	"kacperkrolak/golang-platformer-game/pkg/world/block/ground"
	"kacperkrolak/golang-platformer-game/pkg/world/block/stone"
	"kacperkrolak/golang-platformer-game/pkg/world/gamemap/block"
	"kacperkrolak/golang-platformer-game/pkg/world/tilemap/tile"

	"github.com/hajimehoshi/ebiten/v2"
)

func TestParser_Load(t *testing.T) {
	tileSize := 10
	tileSheet := &ebiten.Image{}

	mapInput := `
author=John Doe
---
cS_
xxx
_xc
---
_xx
x__
___
---
`
	expected := ParsedData{
		Blocks: [][]block.Block{
			{&empty.Block{}, &empty.Block{}, &empty.Block{}},
			{&ground.Block{}, &ground.Block{}, &ground.Block{}},
			{&empty.Block{}, &ground.Block{}, &empty.Block{}},
		},
		Tiles: [][]tile.Tile{
			{&empty.Tile{}, &stone.Tile{}, &stone.Tile{}},
			{&stone.Tile{}, &empty.Tile{}, &empty.Tile{}},
			{&empty.Tile{}, &empty.Tile{}, &empty.Tile{}},
		},
		CoinPositions: []vector.Vector2{
			{X: 0, Y: 0},
			{X: 2 * float64(tileSize), Y: 2 * float64(tileSize)},
		},
		SpawnPoint: vector.Vector2{X: 1 * float64(tileSize), Y: 0},
	}

	parser := Parser{
		TileSheet: tileSheet,
		TileSize:  tileSize,
	}

	mapReader := strings.NewReader(mapInput)
	parsedData, err := parser.Load(mapReader)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	for i, row := range expected.Blocks {
		for j, block := range row {
			if reflect.TypeOf(parsedData.Blocks[i][j]) != reflect.TypeOf(block) {
				t.Errorf("expected block %v, got %v", block, parsedData.Blocks[i][j])
			}
		}
	}

	for i, row := range expected.Tiles {
		for j, tile := range row {
			if reflect.TypeOf(parsedData.Tiles[i][j]) != reflect.TypeOf(tile) {
				t.Errorf("expected tile %v, got %v", tile, parsedData.Tiles[i][j])
			}
		}
	}

	if !reflect.DeepEqual(parsedData.CoinPositions, expected.CoinPositions) {
		t.Errorf("expected coin positions %+v, got %+v", expected.CoinPositions, parsedData.CoinPositions)
	}

	if !reflect.DeepEqual(parsedData.SpawnPoint, expected.SpawnPoint) {
		t.Errorf("expected spawn point %v, got %v", expected.SpawnPoint, parsedData.SpawnPoint)
	}
}

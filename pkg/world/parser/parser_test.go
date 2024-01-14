package parser

import (
	"bufio"
	"kacperkrolak/golang-platformer-game/pkg/world/block/empty"
	"kacperkrolak/golang-platformer-game/pkg/world/block/ground"
	"kacperkrolak/golang-platformer-game/pkg/world/block/spikes"
	"kacperkrolak/golang-platformer-game/pkg/world/block/stone"
	"kacperkrolak/golang-platformer-game/pkg/world/gamemap/block"
	"kacperkrolak/golang-platformer-game/pkg/world/tilemap/tile"
	"reflect"
	"strings"
	"testing"
)

func TestParser_readMapData(t *testing.T) {
	testCases := []struct {
		name    string
		input   string
		wanted  [][]block.Block
		wantErr bool
	}{
		{
			name:    "should return an error if the reader is nil or empty",
			input:   "",
			wanted:  nil,
			wantErr: true,
		},
		{
			name: "should return a correct map given the following input",
			input: `_____
x^___
xxxxx
---
`,
			wanted: [][]block.Block{
				{
					&empty.Block{},
					&empty.Block{},
					&empty.Block{},
					&empty.Block{},
					&empty.Block{},
				},
				{
					&ground.Block{},
					&spikes.Block{},
					&empty.Block{},
					&empty.Block{},
					&empty.Block{},
				},
				{
					&ground.Block{},
					&ground.Block{},
					&ground.Block{},
					&ground.Block{},
					&ground.Block{},
				},
			},
			wantErr: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			parser := Parser{}
			reader := strings.NewReader(testCase.input)
			scanner := bufio.NewScanner(reader)
			got, err := parser.readMapData(scanner)
			if (err != nil) != testCase.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}

			for y, row := range got {
				for x, tile := range row {
					if reflect.TypeOf(tile) != reflect.TypeOf(testCase.wanted[y][x]) {
						t.Errorf("Load() expected tile type %v at position %d:%d, got %v", reflect.TypeOf(testCase.wanted[y][x]), x, y, reflect.TypeOf(tile))
					}
				}
			}
		})
	}

}

func TestParser_readTilemMapData(t *testing.T) {
	testCases := []struct {
		name    string
		input   string
		wanted  [][]tile.Tile
		wantErr bool
	}{
		{
			name:    "should return an error if the reader is nil or empty",
			input:   "",
			wanted:  nil,
			wantErr: true,
		},
		{
			name: "should return a correct map given the following input",
			input: `_____
xx___
xxxxx
`,
			wanted: [][]tile.Tile{
				{
					&empty.Tile{},
					&empty.Tile{},
					&empty.Tile{},
					&empty.Tile{},
					&empty.Tile{},
				},
				{
					&stone.Tile{},
					&stone.Tile{},
					&empty.Tile{},
					&empty.Tile{},
					&empty.Tile{},
				},
				{
					&stone.Tile{},
					&stone.Tile{},
					&stone.Tile{},
					&stone.Tile{},
					&stone.Tile{},
				},
			},
			wantErr: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			parser := Parser{}
			reader := strings.NewReader(testCase.input)
			scanner := bufio.NewScanner(reader)
			got, err := parser.readTileMapData(scanner)
			if (err != nil) != testCase.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}

			for y, row := range got {
				for x, tile := range row {
					if reflect.TypeOf(tile) != reflect.TypeOf(testCase.wanted[y][x]) {
						t.Errorf("Load() expected tile type %v at position %d:%d, got %v", reflect.TypeOf(testCase.wanted[y][x]), x, y, reflect.TypeOf(tile))
					}
				}
			}
		})
	}

}

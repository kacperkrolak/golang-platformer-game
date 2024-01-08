package parser

import (
	"kacperkrolak/golang-platformer-game/pkg/gamemap/tile"
	"kacperkrolak/golang-platformer-game/pkg/gamemap/tile/empty"
	"kacperkrolak/golang-platformer-game/pkg/gamemap/tile/ground"
	"kacperkrolak/golang-platformer-game/pkg/gamemap/tile/spikes"
	"reflect"
	"strings"
	"testing"
)

func TestParser_LoadTiles(t *testing.T) {
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
			input: `
---
_____
x^___
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
					&ground.Tile{},
					&spikes.Tile{},
					&empty.Tile{},
					&empty.Tile{},
					&empty.Tile{},
				},
				{
					&ground.Tile{},
					&ground.Tile{},
					&ground.Tile{},
					&ground.Tile{},
					&ground.Tile{},
				},
			},
			wantErr: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			parser := MapDataParser{}
			got, err := parser.LoadTiles(strings.NewReader(testCase.input))
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

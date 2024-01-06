package parser

import (
	"kacperkrolak/golang-platformer-game/pkg/gamemap"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParser_LoadTiles(t *testing.T) {
	testCases := []struct {
		name    string
		input   string
		wanted  [][]gamemap.Tile
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
xx___
xxxxx
`,
			wanted: [][]gamemap.Tile{
				{
					gamemap.Tile{Type: gamemap.EMPTY},
					gamemap.Tile{Type: gamemap.EMPTY},
					gamemap.Tile{Type: gamemap.EMPTY},
					gamemap.Tile{Type: gamemap.EMPTY},
					gamemap.Tile{Type: gamemap.EMPTY},
				},
				{
					gamemap.Tile{Type: gamemap.DIRT},
					gamemap.Tile{Type: gamemap.DIRT},
					gamemap.Tile{Type: gamemap.EMPTY},
					gamemap.Tile{Type: gamemap.EMPTY},
					gamemap.Tile{Type: gamemap.EMPTY},
				},
				{
					gamemap.Tile{Type: gamemap.DIRT},
					gamemap.Tile{Type: gamemap.DIRT},
					gamemap.Tile{Type: gamemap.DIRT},
					gamemap.Tile{Type: gamemap.DIRT},
					gamemap.Tile{Type: gamemap.DIRT},
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

			if !testCase.wantErr && !cmp.Equal(got, testCase.wanted) {
				t.Errorf("Load() got = %v, want %v", got, testCase.wanted)
			}
		})
	}

}

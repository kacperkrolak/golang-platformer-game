/*
Package parser is responsible for parsing a map from a reader.

The map is a text file with three parts separated by three dashes.
The first part is the meta data, the second part is the block data and the third part is the tile data.
An example map may look like this:
AUTHOR=someone
---
_____
xx___
xxxxx
---
xxxxx
x____
x____
*/
package parser

import (
	"bufio"
	"fmt"
	"io"
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"
	"kacperkrolak/golang-platformer-game/pkg/world/block/empty"
	"kacperkrolak/golang-platformer-game/pkg/world/block/ground"
	"kacperkrolak/golang-platformer-game/pkg/world/block/spikes"
	"kacperkrolak/golang-platformer-game/pkg/world/block/spring"
	"kacperkrolak/golang-platformer-game/pkg/world/block/stone"
	"kacperkrolak/golang-platformer-game/pkg/world/gamemap/block"
	"kacperkrolak/golang-platformer-game/pkg/world/tilemap/tile"

	"github.com/hajimehoshi/ebiten/v2"
)

// Parser is responsible for parsing a map from a reader.
type Parser struct {
	TileSheet *ebiten.Image
	TileSize  int
}

// Reads a map from a reader and returns a slice of blocks and tiles.
//
// The maps are in the following format:
// <variables like in .env>
// ---
// <multiple rows of the same length representing blocks>
// ---
// <multiple rows of the same length representing tiles>
func (parser Parser) Load(reader io.Reader) (ParsedData, error) {
	parsedData := ParsedData{}
	if reader == nil {
		return ParsedData{}, fmt.Errorf("reader cannot be nil")
	}

	scanner := bufio.NewScanner(reader)
	if err := parser.readMetaData(scanner, &parsedData); err != nil {
		return ParsedData{}, err
	}

	if err := parser.readMapData(scanner, &parsedData); err != nil {
		return ParsedData{}, err
	}

	if err := parser.readTileMapData(scanner, &parsedData); err != nil {
		return ParsedData{}, err
	}

	return parsedData, nil
}

// Reads the first part of the map which is the meta data.
func (parser Parser) readMetaData(scanner *bufio.Scanner, parsedData *ParsedData) error {
	for scanner.Scan() {
		line := scanner.Text()
		if line == "---" {
			return nil
		}
	}

	return fmt.Errorf("meta data must end with three dashes")
}

// Reads the second part of the map which is the block data.
func (parser Parser) readMapData(scanner *bufio.Scanner, parsedData *ParsedData) error {
	blocks := make([][]block.Block, 0)
	coinPositions := make([]vector.Vector2, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "---" {
			break
		}

		row := make([]block.Block, 0)
		for _, char := range line {
			switch char {
			case 'x':
				row = append(row, &ground.Block{})
			case '^':
				row = append(row, &spikes.Block{})
			case 's':
				row = append(row, &spring.Block{})
			default:
				row = append(row, &empty.Block{})
				if char == 'c' {
					coinPositions = append(coinPositions, vector.Vector2{X: float64(len(row)-1) * float64(parser.TileSize), Y: float64(len(blocks)) * float64(parser.TileSize)})
				}
				if char == 'S' {
					parsedData.SpawnPoint = vector.Vector2{X: float64(len(row)-1) * float64(parser.TileSize), Y: float64(len(blocks)) * float64(parser.TileSize)}
				}
			}
		}

		if len(blocks) > 0 && len(row) != len(blocks[0]) {
			return fmt.Errorf("all rows must have the same length")
		}

		blocks = append(blocks, row)
	}

	if len(blocks) == 0 {
		return fmt.Errorf("map must have at least one row")
	}

	parsedData.Blocks = blocks
	parsedData.CoinPositions = coinPositions

	return nil
}

// Reads the third part of the map which is the tile data.
func (parser Parser) readTileMapData(scanner *bufio.Scanner, parsedData *ParsedData) error {
	tiles := make([][]tile.Tile, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		row := make([]tile.Tile, 0)
		for _, char := range line {
			if char == 'x' {
				row = append(row, &stone.Tile{})
			} else {
				row = append(row, &empty.Tile{})
			}
		}

		if len(tiles) > 0 && len(row) != len(tiles[0]) {
			return fmt.Errorf("all rows must have the same length")
		}

		tiles = append(tiles, row)
	}

	if len(tiles) == 0 {
		return fmt.Errorf("map must have at least one row")
	}

	parsedData.Tiles = tiles
	return nil
}

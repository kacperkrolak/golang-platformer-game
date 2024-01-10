package parser

import (
	"bufio"
	"fmt"
	"io"
	"kacperkrolak/golang-platformer-game/pkg/gamemap/tile"
	"kacperkrolak/golang-platformer-game/pkg/gamemap/tile/empty"
	"kacperkrolak/golang-platformer-game/pkg/gamemap/tile/ground"
	"kacperkrolak/golang-platformer-game/pkg/gamemap/tile/spikes"
	"kacperkrolak/golang-platformer-game/pkg/gamemap/tile/spring"
)

type MapDataParser struct {
}

func (parser MapDataParser) LoadTiles(reader io.Reader) ([][]tile.Tile, error) {
	if reader == nil {
		return nil, fmt.Errorf("reader cannot be nil")
	}

	scanner := bufio.NewScanner(reader)
	parser.readMetaData(scanner)
	tiles, err := parser.readMapData(scanner)
	if err != nil {
		return nil, err
	}

	return tiles, nil
}

func (parser MapDataParser) readMetaData(scanner *bufio.Scanner) error {
	for scanner.Scan() {
		line := scanner.Text()
		if line == "---" {
			return nil
		}
	}

	return fmt.Errorf("meta data must end with three dashes")
}

func (parser MapDataParser) readMapData(scanner *bufio.Scanner) ([][]tile.Tile, error) {
	tiles := make([][]tile.Tile, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		row := make([]tile.Tile, 0)
		for _, char := range line {
			if char == '_' {
				row = append(row, &empty.Tile{})
			}
			if char == 'x' {
				row = append(row, &ground.Tile{})
			}
			if char == '^' {
				row = append(row, &spikes.Tile{})
			}
			if char == 's' {
				row = append(row, &spring.Tile{})
			}
		}

		if len(tiles) > 0 && len(row) != len(tiles[0]) {
			return nil, fmt.Errorf("all rows must have the same length")
		}

		tiles = append(tiles, row)
	}

	if len(tiles) == 0 {
		return nil, fmt.Errorf("map must have at least one row")
	}

	return tiles, nil
}

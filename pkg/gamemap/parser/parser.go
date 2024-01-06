package parser

import (
	"bufio"
	"fmt"
	"io"
	"kacperkrolak/golang-platformer-game/pkg/gamemap"
)

type MapDataParser struct {
}

func (parser MapDataParser) LoadTiles(reader io.Reader) ([][]gamemap.Tile, error) {
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

func (parser MapDataParser) readMapData(scanner *bufio.Scanner) ([][]gamemap.Tile, error) {
	tiles := make([][]gamemap.Tile, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		row := make([]gamemap.Tile, 0)
		for _, char := range line {
			row = append(row, gamemap.Tile{Type: gamemap.TileType(char)})
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

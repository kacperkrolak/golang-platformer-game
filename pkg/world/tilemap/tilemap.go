package tilemap

import (
	"kacperkrolak/golang-platformer-game/pkg/world/tilemap/tile"
)

type Map struct {
	Tiles    [][]tile.Tile
	TileSize int
}

func MakeMap(tiles [][]tile.Tile, tileSize int) Map {
	m := Map{
		Tiles:    tiles,
		TileSize: tileSize,
	}
	m.FindVariants()

	return m
}

// For each tile, chech if there is different tile on the left, right, top or bottom.
func (m Map) FindVariants() {
	for y, row := range m.Tiles {
		for x, tileInstance := range row {
			var neighbours [4]tile.Tile
			if x > 0 {
				neighbours[0] = m.Tiles[y][x-1]
			}
			if y > 0 {
				neighbours[1] = m.Tiles[y-1][x]
			}
			if x < len(row)-1 {
				neighbours[2] = m.Tiles[y][x+1]
			}
			if y < len(m.Tiles)-1 {
				neighbours[3] = m.Tiles[y+1][x]
			}

			tileInstance.UpdateVariant(neighbours)
		}
	}
}

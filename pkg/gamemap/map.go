package gamemap

import (
	"io"
	"kacperkrolak/golang-platformer-game/pkg/gamemap/tile"
	"kacperkrolak/golang-platformer-game/pkg/physics/box"
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"
)

type Parser interface {
	LoadTiles(io.Reader) [][]tile.Tile
}

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
	m.CreateHitboxes(tileSize)

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

func (m Map) CreateHitboxes(tileSize int) {
	for y, row := range m.Tiles {
		for x, tile := range row {
			if tile.IsCollidable() {
				m.Tiles[y][x].SetHitbox(box.Box{
					Position: vector.Vector2{X: float64(x * tileSize), Y: float64(y * tileSize)},
					Size:     vector.Vector2{X: float64(tileSize), Y: float64(tileSize)},
				})
			}
		}
	}
}

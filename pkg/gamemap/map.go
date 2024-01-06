package gamemap

import (
	"io"
	"kacperkrolak/golang-platformer-game/pkg/physics/box"
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"
)

type Parser interface {
	LoadTiles(io.Reader) [][]Tile
}

type Map struct {
	Tiles    [][]Tile
	TileSize int
}

func MakeMap(tiles [][]Tile, tileSize int) Map {
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
		for x, tile := range row {
			if tile.Type == EMPTY {
				continue
			}

			variant := uint8(0)
			if x > 0 && m.Tiles[y][x-1].Type != tile.Type {
				variant |= LEFT
			}
			if x < len(row)-1 && m.Tiles[y][x+1].Type != tile.Type {
				variant |= RIGHT
			}
			if y > 0 && m.Tiles[y-1][x].Type != tile.Type {
				variant |= TOP
			}
			if y < len(m.Tiles)-1 && m.Tiles[y+1][x].Type != tile.Type {
				variant |= BOTTOM
			}

			m.Tiles[y][x].Variant = variant
		}
	}
}

func (m Map) CreateHitboxes(tileSize int) {
	for y, row := range m.Tiles {
		for x, tile := range row {
			if tile.IsCollidable() {
				m.Tiles[y][x].Hitbox = box.Box{
					Position: vector.Vector2{X: float64(x * tileSize), Y: float64(y * tileSize)},
					Size:     vector.Vector2{X: float64(tileSize), Y: float64(tileSize)},
				}
			}
		}
	}
}

package gamemap

import (
	"kacperkrolak/golang-platformer-game/pkg/physics/box"
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"
	"kacperkrolak/golang-platformer-game/pkg/world/gamemap/block"
)

// Map represents a 2D array of blocks have physics and can be drawn.
type Map struct {
	Blocks   [][]block.Block
	TileSize int
}

// MakeMap takes and updates blocks based on their position to create a map.
func MakeMap(blocks [][]block.Block, tileSize int) Map {
	m := Map{
		Blocks:   blocks,
		TileSize: tileSize,
	}
	m.FindVariants()
	m.CreateHitboxes(tileSize)

	return m
}

// For each tile, chech if there is different tile on the left, right, top or bottom.
func (m Map) FindVariants() {
	for y, row := range m.Blocks {
		for x, tileInstance := range row {
			var neighbours [4]block.Block
			if x > 0 {
				neighbours[0] = m.Blocks[y][x-1]
			}
			if y > 0 {
				neighbours[1] = m.Blocks[y-1][x]
			}
			if x < len(row)-1 {
				neighbours[2] = m.Blocks[y][x+1]
			}
			if y < len(m.Blocks)-1 {
				neighbours[3] = m.Blocks[y+1][x]
			}

			tileInstance.AdaptToNeighbours(neighbours)
		}
	}
}

// Add hitbox to each block based on its position.
func (m Map) CreateHitboxes(tileSize int) {
	for y, row := range m.Blocks {
		for x, tile := range row {
			if tile.IsCollidable() {
				m.Blocks[y][x].SetHitbox(box.Box{
					Position: vector.Vector2{X: float64(x * tileSize), Y: float64(y * tileSize)},
					Size:     vector.Vector2{X: float64(tileSize), Y: float64(tileSize)},
				})
			}
		}
	}
}

// Return height of the map in pixels.
func (m Map) Height() int {
	return len(m.Blocks) * m.TileSize
}

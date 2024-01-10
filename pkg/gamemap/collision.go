package gamemap

import (
	"kacperkrolak/golang-platformer-game/pkg/gamemap/tile"
	"kacperkrolak/golang-platformer-game/pkg/physics/box"
	"math"
)

func isInteger(f float64) bool {
	return f == math.Trunc(f)
}

func (m Map) CollidesWith(box box.Box) []tile.Tile {
	// The is made of tiles of the same size whuch don't move nor overlap,
	// so instead if checking if the box collides with each tile, we can
	// use a simple formula to get the tiles that the box is colliding with
	// and check if they are collidable.
	left := int(box.Left()) / m.TileSize
	if left < 0 {
		left = 0
	}

	// int(box.Right()) / m.TileSize is wrong if box.Right() is exactly a multiple of m.TileSize.
	right := int(box.Right()) / m.TileSize
	if isInteger(box.Right() / float64(m.TileSize)) {
		right--
	}
	if right >= len(m.Tiles[0]) {
		right = len(m.Tiles[0]) - 1
	}

	top := int(box.Top()) / m.TileSize
	if top < 0 {
		top = 0
	}

	// Same as for right, we need to subtract 1 if box.Bottom() is exactly a multiple of m.TileSize.
	bottom := int(box.Bottom()) / m.TileSize
	if isInteger(box.Bottom() / float64(m.TileSize)) {
		bottom--
	}
	if bottom >= len(m.Tiles) {
		bottom = len(m.Tiles) - 1
	}

	collidedWith := make([]tile.Tile, 0)
	for y := top; y <= bottom; y++ {
		for x := left; x <= right; x++ {
			if m.Tiles[y][x].IsCollidable() {
				collidedWith = append(collidedWith, m.Tiles[y][x])
			}
		}
	}

	return collidedWith
}

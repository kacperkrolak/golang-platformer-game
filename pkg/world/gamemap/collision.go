package gamemap

import (
	"kacperkrolak/golang-platformer-game/pkg/physics/box"
	"kacperkrolak/golang-platformer-game/pkg/world/gamemap/block"
	"math"
)

func isInteger(f float64) bool {
	return f == math.Trunc(f)
}

// Intenger divide a number but round it down.
//
// While diving positive integers, the result is rounded down, but when we
// divide negative integers, the result is rounded up. This function always
// rounds down.
func integerDivide(a, b int) int {
	if a < 0 {
		return (a - b + 1) / b
	}
	return a / b
}

func (m Map) CollidesWith(box box.Box) []block.Block {
	// The is made of tiles of the same size whuch don't move nor overlap,
	// so instead if checking if the box collides with each tile, we can
	// use a simple formula to get the tiles that the box is colliding with
	// and check if they are collidable.
	left := integerDivide(int(box.Left()), m.TileSize)
	if left > len(m.Blocks[0]) {
		return make([]block.Block, 0)
	}
	if left < 0 {
		left = 0
	}

	// int(box.Right()) / m.TileSize is wrong if box.Right() is exactly a multiple of m.TileSize.
	right := integerDivide(int(box.Right()), m.TileSize)
	if isInteger(box.Right() / float64(m.TileSize)) {
		right--
	}
	if right < 0 {
		return make([]block.Block, 0)
	}
	if right >= len(m.Blocks[0]) {
		right = len(m.Blocks[0]) - 1
	}

	top := integerDivide(int(box.Top()), m.TileSize)
	if top > len(m.Blocks) {
		return make([]block.Block, 0)
	}
	if top < 0 {
		top = 0
	}

	// Same as for right, we need to subtract 1 if box.Bottom() is exactly a multiple of m.TileSize.
	bottom := integerDivide(int(box.Bottom()), m.TileSize)
	if bottom < 0 {
		return make([]block.Block, 0)
	}
	if isInteger(box.Bottom() / float64(m.TileSize)) {
		bottom--
	}
	if bottom >= len(m.Blocks) {
		bottom = len(m.Blocks) - 1
	}

	collidedWith := make([]block.Block, 0)
	for y := top; y <= bottom; y++ {
		for x := left; x <= right; x++ {
			if m.Blocks[y][x].IsCollidable() {
				collidedWith = append(collidedWith, m.Blocks[y][x])
			}
		}
	}

	return collidedWith
}

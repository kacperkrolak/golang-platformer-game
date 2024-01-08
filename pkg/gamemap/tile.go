package gamemap

import "kacperkrolak/golang-platformer-game/pkg/physics/box"

type TileType rune

const (
	EMPTY  TileType = '_'
	DIRT   TileType = 'x'
	SPIKES TileType = '^'
)

// Flags to indecate whether there different type of tile on the left, right, top or bottom.
const (
	LEFT              = 1 << iota
	RIGHT             = 1 << iota
	TOP               = 1 << iota
	BOTTOM            = 1 << iota
	SURROUNDED        = LEFT | RIGHT | TOP | BOTTOM
	ALONE             = 0
	LEFT_RIGHT        = LEFT | RIGHT
	TOP_BOTTOM        = TOP | BOTTOM
	LEFT_TOP          = LEFT | TOP
	LEFT_BOTTOM       = LEFT | BOTTOM
	RIGHT_TOP         = RIGHT | TOP
	RIGHT_BOTTOM      = RIGHT | BOTTOM
	LEFT_TOP_RIGHT    = LEFT | TOP | RIGHT
	BOTTOM_LEFT_TOP   = BOTTOM | LEFT | TOP
	BOTTOM_RIGHT_TOP  = BOTTOM | RIGHT | TOP
	BOTTOM_LEFT_RIGHT = BOTTOM | LEFT | RIGHT
)

type Tile struct {
	Type    TileType
	Variant uint8
	Hitbox  box.Box
}

func (t Tile) IsCollidable() bool {
	return t.Type != EMPTY
}

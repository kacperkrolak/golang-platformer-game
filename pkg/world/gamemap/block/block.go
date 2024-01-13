package block

import (
	"kacperkrolak/golang-platformer-game/pkg/physics/box"
	"kacperkrolak/golang-platformer-game/pkg/physics/rigidbody"
	"kacperkrolak/golang-platformer-game/pkg/world/tilemap/tile"
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

type Block interface {
	IsCollidable() bool
	IsSolid() bool
	Hitbox() box.Box
	SetHitbox(box.Box)
	IsDeadly() bool
	OnCollision(*rigidbody.Rigidbody)
	AdaptToNeighbours([4]Block) // Neighbours: left, top, right, bottom
	tile.Tile
}

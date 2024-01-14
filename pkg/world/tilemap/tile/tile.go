// Package tile defines a structure representing a square texture
// which can change its appearance based on its neighbours.
package tile

import (
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"

	"github.com/hajimehoshi/ebiten/v2"
)

// The following constants are used to determine which sides of a tile are
// surrounded by other tiles.
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

// Tile is a square texture which can change its appearance based on its neighbours.
type Tile interface {
	GetGroup() string                                       // Tiles in the same group can be visually connected
	UpdateVariant([4]Tile)                                  // Neighbours: left, top, right, bottom
	Draw(*ebiten.Image, vector.Vector2, *ebiten.Image, int) // Screen, camera offset, tile sheet, tile size
}

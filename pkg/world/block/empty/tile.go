package empty

import (
	"kacperkrolak/golang-platformer-game/pkg/physics/box"
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"
	"kacperkrolak/golang-platformer-game/pkg/world/tilemap/tile"

	"github.com/hajimehoshi/ebiten/v2"
)

type Tile struct {
	variant uint8
	hitbox  box.Box
}

func (t *Tile) UpdateVariant(neighbours [4]tile.Tile) {
}

func (t Tile) Draw(screen *ebiten.Image, screenPosition vector.Vector2, tileSheet *ebiten.Image, tileSize int) {
}

func (t Tile) GetGroup() string {
	return "empty"
}

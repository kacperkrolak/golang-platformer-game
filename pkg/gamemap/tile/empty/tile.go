package empty

import (
	"kacperkrolak/golang-platformer-game/pkg/gamemap/tile"
	"kacperkrolak/golang-platformer-game/pkg/physics/box"
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"

	"github.com/hajimehoshi/ebiten/v2"
)

type Tile struct {
	variant uint8
	hitbox  box.Box
}

func (t Tile) IsCollidable() bool {
	return false
}

func (t Tile) IsSolid() bool {
	return false
}

func (t *Tile) UpdateVariant(neighbours [4]tile.Tile) {
}

func (t Tile) Draw(screen *ebiten.Image, screenPosition vector.Vector2, tileSheet *ebiten.Image, tileSize int) {
}

func (t Tile) Hitbox() box.Box {
	return t.hitbox
}

func (t *Tile) SetHitbox(hitbox box.Box) {
	t.hitbox = hitbox
}

func (t Tile) IsDeadly() bool {
	return false
}

package mock

import (
	"kacperkrolak/golang-platformer-game/pkg/gamemap/tile"
	"kacperkrolak/golang-platformer-game/pkg/physics/box"
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/stretchr/testify/mock"
)

type MockTile struct {
	mock.Mock
	collidable bool
	solid      bool
	hitbox     box.Box
}

func NewEmptyTile() *MockTile {
	return &MockTile{
		collidable: false,
		solid:      false,
	}
}

func NewCollidableTile() *MockTile {
	return &MockTile{
		collidable: true,
		solid:      true,
	}
}

func (t MockTile) IsCollidable() bool {
	return t.collidable
}

func (t MockTile) IsSolid() bool {
	return t.solid
}

func (t MockTile) UpdateVariant(neighbours [4]tile.Tile) {
	t.Called(neighbours)
}

func (t MockTile) Draw(screen *ebiten.Image, screenPosition vector.Vector2, tileSheet *ebiten.Image, tileSize int) {
}

func (t MockTile) Hitbox() box.Box {
	return t.hitbox
}

func (t *MockTile) SetHitbox(hitbox box.Box) {
	t.hitbox = hitbox
}

func (t MockTile) IsDeadly() bool {
	return false
}

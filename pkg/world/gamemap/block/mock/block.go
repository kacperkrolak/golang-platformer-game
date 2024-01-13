package mock

import (
	"kacperkrolak/golang-platformer-game/pkg/physics/box"
	"kacperkrolak/golang-platformer-game/pkg/physics/rigidbody"
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"
	"kacperkrolak/golang-platformer-game/pkg/world/gamemap/block"
	"kacperkrolak/golang-platformer-game/pkg/world/tilemap/tile"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/stretchr/testify/mock"
)

type MockBlock struct {
	mock.Mock
	collidable bool
	solid      bool
	hitbox     box.Box
}

func NewEmptyTile() *MockBlock {
	return &MockBlock{
		collidable: false,
		solid:      false,
	}
}

func NewCollidableTile() *MockBlock {
	return &MockBlock{
		collidable: true,
		solid:      true,
	}
}

func (t MockBlock) IsCollidable() bool {
	return t.collidable
}

func (t MockBlock) IsSolid() bool {
	return t.solid
}

func (t MockBlock) AdaptToNeighbours(neighbours [4]block.Block) {
	t.Called(neighbours)
}

func (t MockBlock) Draw(screen *ebiten.Image, screenPosition vector.Vector2, tileSheet *ebiten.Image, tileSize int) {
}

func (t MockBlock) Hitbox() box.Box {
	return t.hitbox
}

func (t *MockBlock) SetHitbox(hitbox box.Box) {
	t.hitbox = hitbox
}

func (t MockBlock) IsDeadly() bool {
	return false
}

func (t MockBlock) OnCollision(rigidbody *rigidbody.Rigidbody) {
}

func (t MockBlock) UpdateVariant(neighbours [4]tile.Tile) {
}

func (t MockBlock) GetGroup() string {
	return ""
}

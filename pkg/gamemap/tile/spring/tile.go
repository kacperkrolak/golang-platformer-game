package spring

import (
	"image"
	"kacperkrolak/golang-platformer-game/pkg/gamemap/tile"
	"kacperkrolak/golang-platformer-game/pkg/physics/box"
	"kacperkrolak/golang-platformer-game/pkg/physics/rigidbody"
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"

	"github.com/hajimehoshi/ebiten/v2"
)

type Tile struct {
	hitbox box.Box
}

func (t Tile) IsCollidable() bool {
	return true
}

func (t Tile) IsSolid() bool {
	return true
}

var TilePosition vector.Vector2 = vector.Vector2{X: 0, Y: 90}

func (t Tile) Draw(screen *ebiten.Image, screenPosition vector.Vector2, tileSheet *ebiten.Image, tileSize int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(screenPosition.X, screenPosition.Y)

	tilePos := TilePosition
	tileImg := tileSheet.SubImage(image.Rect(int(tilePos.X), int(tilePos.Y), int(tilePos.X)+tileSize, int(tilePos.Y)+tileSize)).(*ebiten.Image)

	screen.DrawImage(tileImg, op)
}

func (t *Tile) UpdateVariant(neighbours [4]tile.Tile) {
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

func (t Tile) OnCollision(rigidbody *rigidbody.Rigidbody) {
	if rigidbody.Velocity.Y < -0 {
		return
	}

	pushStrength := rigidbody.Velocity.Y
	if pushStrength > 500 {
		pushStrength = 500
	}

	rigidbody.AddForce(vector.Up().Scaled(pushStrength))
}

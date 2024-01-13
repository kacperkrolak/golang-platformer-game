package spring

import (
	"kacperkrolak/golang-platformer-game/pkg/physics/box"
	"kacperkrolak/golang-platformer-game/pkg/physics/rigidbody"
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"
	"kacperkrolak/golang-platformer-game/pkg/world/gamemap/block"
	"kacperkrolak/golang-platformer-game/pkg/world/tilemap/tile"
)

type Block struct {
	Tile
	hitbox box.Box
}

func (b Block) IsCollidable() bool {
	return true
}

func (b Block) IsSolid() bool {
	return true
}

func (b Block) Hitbox() box.Box {
	return b.hitbox
}

func (b *Block) SetHitbox(hitbox box.Box) {
	b.hitbox = hitbox
}

func (b Block) IsDeadly() bool {
	return false
}

func (b Block) OnCollision(rigidbody *rigidbody.Rigidbody) {
	if rigidbody.Velocity.Y < -0 {
		return
	}

	pushStrength := rigidbody.Velocity.Y
	if pushStrength > 500 {
		pushStrength = 500
	}

	rigidbody.AddForce(vector.Up().Scaled(pushStrength))
}

func (b *Block) AdaptToNeighbours(neighbours [4]block.Block) {
	neighbourTiles := [4]tile.Tile{}
	for i, neighbour := range neighbours {
		neighbourTiles[i] = neighbour
	}

	b.UpdateVariant(neighbourTiles)
}

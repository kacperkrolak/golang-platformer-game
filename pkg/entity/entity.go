package entity

import (
	"kacperkrolak/golang-platformer-game/pkg/physics/rigidbody"
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Entity interface {
	Update(deltaTime time.Duration)
	Draw(*ebiten.Image, vector.Vector2)
	OnCollisionWithPlayer()
	IsDead() bool
	GetRigidbody() *rigidbody.Rigidbody
}

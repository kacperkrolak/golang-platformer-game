package coin

import (
	"kacperkrolak/golang-platformer-game/pkg/physics/rigidbody"
	"kacperkrolak/golang-platformer-game/pkg/visuals/particle"
	"kacperkrolak/golang-platformer-game/pkg/visuals/particle/confetti"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Coin struct {
	Rigidbody      rigidbody.Rigidbody
	TimePassed     time.Duration
	TileSheet      *ebiten.Image
	ParticleSystem *particle.ParticleSystem
	dead           bool
}

func (c *Coin) Update(deltaTime time.Duration) {
	c.TimePassed += deltaTime
}

func (c *Coin) OnCollisionWithPlayer() {
	c.ParticleSystem.AddParticles(confetti.CreateEffect(c.Rigidbody.Hitbox.Center(), 10, 3, 2, 40))
	c.dead = true
}

func (c *Coin) IsDead() bool {
	return c.dead
}

func (c *Coin) GetRigidbody() *rigidbody.Rigidbody {
	return &c.Rigidbody
}

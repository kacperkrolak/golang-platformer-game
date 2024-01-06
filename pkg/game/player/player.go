package player

import (
	"kacperkrolak/golang-platformer-game/pkg/physics/box"
	"kacperkrolak/golang-platformer-game/pkg/physics/rigidbody"
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"
	"kacperkrolak/golang-platformer-game/pkg/visuals/particle"
	"kacperkrolak/golang-platformer-game/pkg/visuals/particle/smoke"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	Rigidbody     rigidbody.Rigidbody
	Speed         float64 // How many tiles can a player move in 1 second
	CameraOffsetX float64
	Grounded      bool
	Frame         int
	State         int
	IsMoving      bool
	FacingRight   bool
}

func (p *Player) Update(tps float64, tileSize int, particleSystem *particle.ParticleSystem) error {
	p.Frame += 1
	if p.Frame%15 == 0 {
		p.State = (p.State + 1) % 2
		p.IsMoving = false
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.Rigidbody.Hitbox.Position.X += p.Speed / tps
		p.IsMoving = true
		p.FacingRight = true
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.Rigidbody.Hitbox.Position.X -= p.Speed / tps
		p.IsMoving = true
		p.FacingRight = false
	}

	if p.Grounded && ebiten.IsKeyPressed(ebiten.KeySpace) {
		particleSystem.AddParticles(smoke.CreateEffect(p.SurfaceDetector().Center(), 5, 7, 0.75, 60))
		JUMP_FORCE := 4.0
		p.Rigidbody.AddForce(vector.Vector2{X: 0, Y: -JUMP_FORCE})
	}

	// Gravity
	p.Rigidbody.AddForce(vector.Vector2{X: 0, Y: 9.81 / tps})
	// p.Rigidbody.AddForce(vector.Friction(p.Rigidbody.Velocity, 0.5/tps))
	p.Rigidbody.ApplyAcceleration()
	p.Rigidbody.ApplyVelocity()
	p.Rigidbody.LimitHorizontalVelocity(p.Speed)
	p.Rigidbody.LimitHorizontalVelocity(20)

	return nil
}

func (p Player) GetRigidbody() rigidbody.Rigidbody {
	return p.Rigidbody
}

func (p Player) SurfaceDetector() box.Box {
	playerBottom := p.Rigidbody.Hitbox.Bottom()
	playerLeft := p.Rigidbody.Hitbox.Left()

	return box.Box{
		Position: vector.Vector2{X: playerLeft, Y: playerBottom},
		Size:     vector.Vector2{X: p.Rigidbody.Hitbox.Size.X, Y: 2},
	}
}

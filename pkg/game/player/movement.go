package player

import (
	"image/color"
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"
	"kacperkrolak/golang-platformer-game/pkg/visuals/particle/smoke"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func (p *Player) HandleInput(deltaTime float64) {
	if p.MovingCooldown > 0 {
		return
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.Rigidbody.Hitbox.Position.X += p.motion.Speed * deltaTime
		p.IsMoving = true
		p.FacingRight = true
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.Rigidbody.Hitbox.Position.X -= p.motion.Speed * deltaTime
		p.IsMoving = true
		p.FacingRight = false
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) && p.JumpingCooldown <= 0 {
		if p.motion.Grounded {
			p.Jump()
			p.JumpingCooldown = time.Millisecond * time.Duration(200)
			p.wallJumpData.WallSlidingCooldown = time.Millisecond * time.Duration(200)
		} else if p.wallJumpData.IsWallSliding {
			p.WallJump()
			p.JumpingCooldown = time.Millisecond * time.Duration(200)
			p.wallJumpData.WallSlidingCooldown = time.Millisecond * time.Duration(300)
			p.MovingCooldown = time.Millisecond * time.Duration(200)
		}
	}
}

func (p *Player) Jump() {
	p.motion.IsJumping = true
	p.ParticleSystem.AddParticles(smoke.CreateEffect(p.SurfaceDetector().Center(), 5, 7, 0.75, 60, color.RGBA{R: 255, G: 255, B: 255, A: 255}))
	JUMP_FORCE := 4.0
	p.Rigidbody.AddForce(vector.Vector2{X: 0, Y: -1}.Scaled(JUMP_FORCE))
}

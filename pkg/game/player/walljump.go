package player

import (
	"image/color"
	"kacperkrolak/golang-platformer-game/pkg/physics/box"
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"
	"kacperkrolak/golang-platformer-game/pkg/visuals/particle/smoke"
	"time"
)

type wallJumpData struct {
	WallSlidingTime     time.Duration // Allows user to jump shortly after stopping wall sliding
	WallSlidingCooldown time.Duration // Disables wall sliding for a short time after jumping
	IsWallSliding       bool
	IsWallJumping       bool
	WallJumpRight       bool
}

func (p *Player) WallJump() {
	p.wallJumpData.IsWallJumping = true
	p.ParticleSystem.AddParticles(smoke.CreateEffect(p.WallDetector().Center(), 5, 7, 0.75, 60, color.RGBA{R: 255, G: 255, B: 255, A: 255}))
	JUMP_FORCE := 5.0
	jumpVector := vector.Vector2{X: 0.6, Y: -0.8}
	if p.wallJumpData.WallJumpRight {
		jumpVector.X = -jumpVector.X
	}

	// Jumping should reset the velocity, because otherwise the speed can stack up
	p.Rigidbody.Velocity.Y = 0
	p.Rigidbody.Acceleration = jumpVector.Scaled(JUMP_FORCE)
}

func (p Player) WallDetector() box.Box {
	size := vector.Vector2{X: 2, Y: p.Rigidbody.Hitbox.Size.Y - 2}
	left := p.Rigidbody.Hitbox.Right()
	if !p.FacingRight {
		left = p.Rigidbody.Hitbox.Left() - size.X
	}

	return box.Box{
		Position: vector.Vector2{X: left, Y: p.Rigidbody.Hitbox.Top() + 1},
		Size:     size,
	}
}

func (p *Player) UpdateWallSlidingState(wallSliding bool) {
	if !p.wallJumpData.IsWallSliding && wallSliding && p.wallJumpData.WallSlidingTime <= 0 {
		p.ParticleSystem.AddParticles(smoke.CreateEffect(p.SurfaceDetector().Center(), 5, 1, 0.5, 60, color.RGBA{R: 222, G: 184, B: 135, A: 255}))
	}

	wallSlidingTimeLimit := time.Millisecond * time.Duration(500)
	if wallSliding {
		p.wallJumpData.WallSlidingTime = wallSlidingTimeLimit
		p.wallJumpData.IsWallJumping = false
		p.wallJumpData.WallJumpRight = p.FacingRight
	}

	p.wallJumpData.IsWallSliding = wallSliding
}

func (p Player) IsWalled() bool {
	// Grounded player cannot be walled
	if p.motion.Grounded {
		return false
	}

	// Disables wall sliding for a short time after jumping
	if p.wallJumpData.WallSlidingCooldown > 0 {
		return false
	}

	return p.wallJumpData.IsWallSliding && !p.wallJumpData.IsWallJumping && p.wallJumpData.WallSlidingTime > 0
}

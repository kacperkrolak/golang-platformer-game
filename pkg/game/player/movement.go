package player

import (
	"image/color"
	"kacperkrolak/golang-platformer-game/pkg/input"
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"
	"kacperkrolak/golang-platformer-game/pkg/visuals/particle/smoke"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const ACCELERATION_RATE = 0.1
const DECCELERATION_RATE = 0.05

// Don't make it too small, as it's multiplied by ACCELERATION_RATE or DECCELERATION_RATE.
const ACCELERATION_AIR_RATE = 1
const DECCELERATION_RATE_AIR = 0.5

func (p *Player) HandleInput(deltaTime float64) {
	inputDirection := input.GetHorizontal()
	// targetSpeed := inputDirection * p.motion.Speed
	targetSpeed := inputDirection * p.Speed
	speedDifference := targetSpeed - p.Rigidbody.Velocity.X

	// We need different rates for moving and stopping.
	var accelerationRate float64
	// It could be 0 for keyboard, but 0.01 is better for joysticks.
	if math.Abs(targetSpeed) > 0.01 {
		accelerationRate = ACCELERATION_RATE
		if !p.IsGrounded() {
			accelerationRate *= ACCELERATION_AIR_RATE
		}
	} else {
		accelerationRate = DECCELERATION_RATE
		if !p.IsGrounded() {
			accelerationRate *= DECCELERATION_RATE_AIR
		}
	}

	// We raise the horizontal force to the power of 2 to make it increase more
	// when the speed difference is bigger.
	horizontalForce := math.Abs(speedDifference) * accelerationRate * math.Copysign(1, speedDifference)
	p.Rigidbody.AddForce(vector.Right().Scaled(horizontalForce).Scaled(p.motion.movementSlowdown))

	// We don't want to change facing direction if there is no input (direction is 0).
	if inputDirection > 0 {
		p.FacingRight = true
	} else if inputDirection < 0 {
		p.FacingRight = false
	}

	if input.IsJumpPressed() && p.JumpingCooldown <= 0 {
		if p.IsGroundedLate() {
			p.Jump()
			p.JumpingCooldown = time.Millisecond * time.Duration(200)
			p.wallJumpData.WallSlidingCooldown = time.Millisecond * time.Duration(200)
		} else if p.IsWalled() {
			p.WallJump()
			p.JumpingCooldown = time.Millisecond * time.Duration(200)
			p.wallJumpData.WallSlidingCooldown = time.Millisecond * time.Duration(300)
			p.MovingCooldown = time.Millisecond * time.Duration(300)
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyQ) && p.JumpingCooldown <= 0 {
		p.WallJump()
		p.JumpingCooldown = time.Millisecond * time.Duration(200)
		p.wallJumpData.WallSlidingCooldown = time.Millisecond * time.Duration(300)
		p.MovingCooldown = time.Millisecond * time.Duration(200)
	}
}

func (p *Player) Jump() {
	p.motion.IsJumping = true
	p.ParticleSystem.AddParticles(smoke.CreateEffect(p.SurfaceDetector().Center(), 5, 7, 0.75, 60, color.RGBA{R: 255, G: 255, B: 255, A: 255}))
	JUMP_FORCE := 250.0
	p.Rigidbody.AddForce(vector.Vector2{X: 0, Y: -1}.Scaled(JUMP_FORCE))
}

func (p Player) OnBumping(displacementVector vector.Vector2) {
	if displacementVector.X != 0 {
		p.Rigidbody.Velocity.X = 0
	}
	if displacementVector.Y > 0 {
		p.Rigidbody.Velocity.Y = 0
		p.Rigidbody.AddForce(vector.Down().Scaled(p.motion.Gravity).Scaled(2))
	}
	if displacementVector.Y < 0 {
		p.Rigidbody.Velocity.Y = 0
	}
}

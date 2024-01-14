/*
Package player implements the player character and its logic.
*/
package player

import (
	"image/color"
	"kacperkrolak/golang-platformer-game/pkg/input"
	"kacperkrolak/golang-platformer-game/pkg/physics/box"
	"kacperkrolak/golang-platformer-game/pkg/physics/rigidbody"
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"
	"kacperkrolak/golang-platformer-game/pkg/visuals/particle"
	"kacperkrolak/golang-platformer-game/pkg/visuals/particle/smoke"
	"math"
	"time"
)

type motionSettings struct {
	Speed                 float64
	Gravity               float64
	GroundedThisFrame     bool
	DurationSinceGrounded time.Duration
	IsJumping             bool
	movementSlowdown      float64
}

type Player struct {
	Rigidbody        rigidbody.Rigidbody
	PreviousVelocity vector.Vector2 // Collisions remove velocity, so this is useful for particle effects
	Speed            float64        // How many tiles can a player move in 1 second
	CameraOffsetX    float64
	Frame            int
	State            int
	FacingRight      bool
	motion           motionSettings
	JumpingCooldown  time.Duration
	MovingCooldown   time.Duration // Slows horizontal movement for a short time after wall jumping
	ParticleSystem   *particle.ParticleSystem
	wallJumpData     wallJumpData
}

func (p *Player) UpdateGroundedState(grounded bool) {
	p.motion.GroundedThisFrame = grounded
}

// To make the player feel more responsive, when the player is the highest point of the jump
// the gravity is reduced and the speed is increased. This makes the player feel like they are
// hanging in the air for a moment.
func (p *Player) AdjustMotionSettings() {
	speed := p.Speed
	p.motion.Gravity = 9.81 * 16 * 3
	p.motion.movementSlowdown = 1
	if p.MovingCooldown > 0 {
		p.motion.movementSlowdown = 0.1
	}

	if p.IsWallSliding() {
		p.motion.Gravity = p.motion.Gravity * 0.1
		return
	}

	// If the player is falling, increase gravity to make the fall faster.
	if p.Rigidbody.Velocity.Y > 0 {
		p.motion.Gravity = p.motion.Gravity * 1.02
	}

	// When player is at the highest point of the jump, reduce gravity and increase speed.
	slowGravityThreshold := 0.5
	if !p.motion.IsJumping || math.Abs(p.Rigidbody.Velocity.Y) < slowGravityThreshold {
		jumpHangGravityMultiplier := 0.8
		jumpHangSpeedMultiplier := 1.1
		p.motion.Gravity = p.motion.Gravity * jumpHangGravityMultiplier
		p.motion.Speed = speed * jumpHangSpeedMultiplier
	}
}

func (p *Player) Update(deltaTime time.Duration, tileSize int) error {
	deltaTimeFloat := deltaTime.Seconds()
	p.wallJumpData.WallSlidingTime -= deltaTime
	p.JumpingCooldown -= deltaTime
	p.wallJumpData.WallSlidingCooldown -= deltaTime
	p.MovingCooldown -= deltaTime

	if p.motion.GroundedThisFrame {
		if p.motion.DurationSinceGrounded > 0 && p.PreviousVelocity.Y > 0 {
			particleCount := uint(p.PreviousVelocity.Y) / 25
			p.ParticleSystem.AddParticles(smoke.CreateEffect(p.SurfaceDetector().Center(), particleCount, 4, 0.5, 60, color.RGBA{R: 222, G: 184, B: 135, A: 255}))
		}
		p.motion.IsJumping = false
		p.motion.DurationSinceGrounded = 0
	} else {
		p.motion.DurationSinceGrounded += deltaTime
	}

	if p.IsGrounded() && p.Rigidbody.Velocity.Y > 0 {
		p.Rigidbody.Velocity.Y = 0
	}

	p.AdjustMotionSettings()

	p.Frame += 1
	if p.Frame%15 == 0 {
		p.State = (p.State + 1) % 2
	}

	p.HandleInput(deltaTimeFloat)

	// Gravity
	p.Rigidbody.AddForce(vector.Vector2{X: 0, Y: p.motion.Gravity * deltaTimeFloat})

	// Friction
	if p.IsGrounded() && input.IsHorizontalIdle() {
		friction := 0.6
		frictionAmount := math.Min(math.Abs(p.Rigidbody.Velocity.X), friction)
		frictionForce := vector.Right().Scaled(-math.Copysign(frictionAmount, p.Rigidbody.Velocity.X))
		p.Rigidbody.AddForce(frictionForce)
	}

	p.Rigidbody.ApplyAcceleration()
	p.Rigidbody.ApplyVelocity(deltaTimeFloat)

	p.PreviousVelocity = p.Rigidbody.Velocity

	return nil
}

func (p Player) GetRigidbody() rigidbody.Rigidbody {
	return p.Rigidbody
}

func (p Player) SurfaceDetector() box.Box {
	playerBottom := p.Rigidbody.Hitbox.Bottom()
	playerLeft := p.Rigidbody.Hitbox.Left()

	return box.Box{
		Position: vector.Vector2{X: playerLeft + 2, Y: playerBottom},
		Size:     vector.Vector2{X: p.Rigidbody.Hitbox.Size.X - 4, Y: 2},
	}
}

// Check if the player touches the ground right now.
func (p Player) IsGrounded() bool {
	return p.motion.DurationSinceGrounded <= 0
}

// Player can jump a short time after leaving the ground.
func (p Player) IsGroundedLate() bool {
	return p.motion.DurationSinceGrounded <= time.Millisecond*100
}

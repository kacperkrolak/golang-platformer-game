package player

import (
	"image/color"
	"kacperkrolak/golang-platformer-game/pkg/physics/box"
	"kacperkrolak/golang-platformer-game/pkg/physics/rigidbody"
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"
	"kacperkrolak/golang-platformer-game/pkg/visuals/particle"
	"kacperkrolak/golang-platformer-game/pkg/visuals/particle/smoke"
	"math"
	"time"
)

type motionSettings struct {
	Speed     float64
	Gravity   float64
	Grounded  bool
	IsJumping bool
}

type Player struct {
	Rigidbody        rigidbody.Rigidbody
	PreviousVelocity vector.Vector2 // Collisions remove velocity, so this is useful for particle effects
	Speed            float64        // How many tiles can a player move in 1 second
	CameraOffsetX    float64
	Frame            int
	State            int
	IsMoving         bool
	FacingRight      bool
	motion           motionSettings
	JumpingCooldown  time.Duration
	MovingCooldown   time.Duration // Disables movement for a short time after wall jumping
	ParticleSystem   *particle.ParticleSystem
	wallJumpData     wallJumpData
}

func (p *Player) UpdateGroundedState(grounded bool) {
	if grounded && !p.motion.Grounded && p.PreviousVelocity.Y > 0 {
		particleCount := uint(p.PreviousVelocity.Y)
		p.ParticleSystem.AddParticles(smoke.CreateEffect(p.SurfaceDetector().Center(), particleCount, 4, 0.5, 60, color.RGBA{R: 222, G: 184, B: 135, A: 255}))
	}

	if grounded {
		p.motion.IsJumping = false
	}

	p.motion.Grounded = grounded
}

// To make the player feel more responsive, when the player is the highest point of the jump
// the gravity is reduced and the speed is increased. This makes the player feel like they are
// hanging in the air for a moment.
func (p *Player) AdjustMotionSettings() {
	baseGravity := 9.81
	if p.IsWalled() {
		p.motion.Gravity = 0
		p.Rigidbody.Velocity.Y = 1
		return
	}

	slowGravityThreshold := 0.5
	if !p.motion.IsJumping || math.Abs(p.Rigidbody.Velocity.Y) < slowGravityThreshold {
		p.motion.Gravity = baseGravity
		p.motion.Speed = p.Speed
	}

	jumpHangGravityMultiplier := 0.8
	jumpHangSpeedMultiplier := 1.1

	p.motion.Gravity = baseGravity * jumpHangGravityMultiplier
	p.motion.Speed = p.Speed * jumpHangSpeedMultiplier
}

func (p *Player) Update(deltaTime time.Duration, tileSize int) error {
	deltaTimeFloat := deltaTime.Seconds()
	p.wallJumpData.WallSlidingTime -= deltaTime
	p.JumpingCooldown -= deltaTime
	p.wallJumpData.WallSlidingCooldown -= deltaTime
	p.MovingCooldown -= deltaTime

	p.AdjustMotionSettings()

	p.Frame += 1
	if p.Frame%15 == 0 {
		p.State = (p.State + 1) % 2
		p.IsMoving = false
	}

	p.HandleInput(deltaTimeFloat)

	// Gravity
	if math.Abs(p.Rigidbody.Velocity.X) > 0 {
		p.Rigidbody.AddForce(vector.Vector2{X: -p.Rigidbody.Velocity.X / 10, Y: 0})
	}
	p.Rigidbody.AddForce(vector.Vector2{X: 0, Y: p.motion.Gravity * deltaTimeFloat})
	// p.Rigidbody.AddForce(vector.Friction(p.Rigidbody.Velocity, 0.5/tps))
	p.Rigidbody.ApplyAcceleration()
	p.Rigidbody.ApplyVelocity()
	p.Rigidbody.LimitHorizontalVelocity(p.motion.Speed)
	p.Rigidbody.LimitHorizontalVelocity(20)

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
		Position: vector.Vector2{X: playerLeft, Y: playerBottom},
		Size:     vector.Vector2{X: p.Rigidbody.Hitbox.Size.X, Y: 2},
	}
}

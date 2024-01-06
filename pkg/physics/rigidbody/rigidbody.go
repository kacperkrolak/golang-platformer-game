package rigidbody

import (
	"kacperkrolak/golang-platformer-game/pkg/physics/box"
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"
)

type Rigidbody struct {
	Hitbox       box.Box
	Velocity     vector.Vector2
	Acceleration vector.Vector2
}

// Checks if the rigidbody is colliding with the box and if so
// moves it out of the box and sets the velocity to 0 in the given axis.
func (rb *Rigidbody) MoveOutOfBox(b box.Box) {
	displacement := b.DisplacementVector(rb.Hitbox)
	if displacement.X != 0 {
		rb.Velocity.X = 0
		rb.Hitbox.Position.Add(displacement)
	}

	if displacement.Y != 0 {
		rb.Velocity.Y = 0
		rb.Hitbox.Position.Add(displacement)
	}
}

func (rb *Rigidbody) CollidesWith(b box.Box) bool {
	return b.DisplacementVector(rb.Hitbox) != vector.Vector2{X: 0, Y: 0}
}

func (rb *Rigidbody) AddForce(force vector.Vector2) {
	rb.Acceleration.Add(force)
}

func (rb *Rigidbody) ApplyAcceleration() {
	rb.Velocity.Add(rb.Acceleration)
	rb.Acceleration.X = 0
	rb.Acceleration.Y = 0
}

func (rb *Rigidbody) ApplyVelocity() {
	rb.Hitbox.Position.Add(rb.Velocity)
}

func (rb *Rigidbody) LimitHorizontalVelocity(maxVelocity float64) {
	if rb.Velocity.X > maxVelocity {
		rb.Velocity.X = maxVelocity
	} else if rb.Velocity.X < -maxVelocity {
		rb.Velocity.X = -maxVelocity
	}
}

func (rb *Rigidbody) LimitVerticalVelocity(maxVelocity float64) {
	if rb.Velocity.Y > maxVelocity {
		rb.Velocity.Y = maxVelocity
	} else if rb.Velocity.Y < -maxVelocity {
		rb.Velocity.Y = -maxVelocity
	}
}

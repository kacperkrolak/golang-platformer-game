/*
Package rigidbody represents a rectangular object with velocity and acceleration that can collide with other objects.
*/
package rigidbody

import (
	"kacperkrolak/golang-platformer-game/pkg/physics/box"
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"
)

// Rectangular object handling basic physics.
type Rigidbody struct {
	Hitbox       box.Box
	Velocity     vector.Vector2
	Acceleration vector.Vector2
}

// Checks if the rigidbody is colliding with the box and if so
// moves it out of the box and sets the velocity to 0 in the given axis.
func (rb *Rigidbody) MoveOutOfBox(b box.Box) vector.Vector2 {
	displacement := b.DisplacementVector(rb.Hitbox)
	rb.Hitbox.Position.Add(displacement)

	return displacement
}

// Checks if the rigidbody is colliding with the box.
func (rb *Rigidbody) CollidesWith(b box.Box) bool {
	return b.DisplacementVector(rb.Hitbox) != vector.Vector2{X: 0, Y: 0}
}

// Adds a vector to the rigidbody's acceleration.
func (rb *Rigidbody) AddForce(force vector.Vector2) {
	rb.Acceleration.Add(force)
}

// Adds the acceleration to the velocity and resets the acceleration.
func (rb *Rigidbody) ApplyAcceleration() {
	rb.Velocity.Add(rb.Acceleration)
	rb.Acceleration.X = 0
	rb.Acceleration.Y = 0
}

// Adds the velocity to the position.
func (rb *Rigidbody) ApplyVelocity(deltaTime float64) {
	rb.Hitbox.Position.Add(rb.Velocity.Scaled(deltaTime))
}

// Limits the X-axis velocity to the given value.
func (rb *Rigidbody) LimitHorizontalVelocity(maxVelocity float64) {
	if rb.Velocity.X > maxVelocity {
		rb.Velocity.X = maxVelocity
	} else if rb.Velocity.X < -maxVelocity {
		rb.Velocity.X = -maxVelocity
	}
}

// Limits the Y-axis velocity to the given value.
func (rb *Rigidbody) LimitVerticalVelocity(maxVelocity float64) {
	if rb.Velocity.Y > maxVelocity {
		rb.Velocity.Y = maxVelocity
	} else if rb.Velocity.Y < -maxVelocity {
		rb.Velocity.Y = -maxVelocity
	}
}

package camera

import (
	"kacperkrolak/golang-platformer-game/pkg/game/camera/mock"
	"kacperkrolak/golang-platformer-game/pkg/physics/box"
	"kacperkrolak/golang-platformer-game/pkg/physics/rigidbody"
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"
	"math"
	"testing"
)

func TestCamera_Update(t *testing.T) {
	targetPosition := vector.Vector2{X: 1, Y: 1}
	target := mock.Target{
		Rigidbody: rigidbody.Rigidbody{
			Hitbox: box.Box{
				Position: targetPosition,
				Size:     vector.Vector2{X: 1, Y: 1},
			},
		},
	}

	initialPosition := vector.Vector2{X: 0, Y: 0}
	initialVelocity := vector.Vector2{X: 0, Y: 0}
	c := &Camera{
		Position:   initialPosition,
		Velocity:   initialVelocity,
		SmoothTime: 0.2,
		Target:     target,
	}

	err := c.Update(0.1)
	if err != nil {
		t.Errorf("Update returned an error: %v", err)
	}

	// Position should be between old position and target position.
	if c.Position.X < math.Min(initialPosition.X, targetPosition.X) || c.Position.X > math.Max(initialPosition.X, targetPosition.X) {
		t.Errorf("expected X to be between %v and %v, got %v", initialPosition.X, targetPosition.X, c.Position.X)
	}

	if c.Position.Y < math.Min(initialPosition.Y, targetPosition.Y) || c.Position.Y > math.Max(initialPosition.Y, targetPosition.Y) {
		t.Errorf("expected Y to be between %v and %v, got %v", initialPosition.Y, targetPosition.Y, c.Position.Y)
	}

	// Velocity should not be zero.
	if c.Velocity == (vector.Vector2{X: 0, Y: 0}) {
		t.Errorf("expected velocity to not be zero, got %v", c.Velocity)
	}
}

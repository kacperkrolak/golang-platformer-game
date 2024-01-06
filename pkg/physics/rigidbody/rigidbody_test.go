package rigidbody

import (
	"kacperkrolak/golang-platformer-game/pkg/physics/box"
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"
	"testing"
)

func TestRigidbody_MoveOutOfBox(t *testing.T) {
	// We won't test every scenario here, because it's already tested in the box package.
	testCases := []struct {
		name     string
		box      box.Box
		box2     box.Box
		expected box.Box
	}{
		{
			name: "should move the box2 down when it's colliding with the box above",
			box: box.Box{
				Position: vector.Vector2{X: 0, Y: 0},
				Size:     vector.Vector2{X: 10, Y: 10},
			},
			box2: box.Box{
				Position: vector.Vector2{X: 5, Y: 5},
				Size:     vector.Vector2{X: 10, Y: 10},
			},
			expected: box.Box{
				Position: vector.Vector2{X: 5, Y: 10},
				Size:     vector.Vector2{X: 10, Y: 10},
			},
		},
		{
			name: "shouldn't move the box2 when it's not colliding with the box",
			box: box.Box{
				Position: vector.Vector2{X: 0, Y: 0},
				Size:     vector.Vector2{X: 10, Y: 10},
			},
			box2: box.Box{
				Position: vector.Vector2{X: 15, Y: 15},
				Size:     vector.Vector2{X: 10, Y: 10},
			},
			expected: box.Box{
				Position: vector.Vector2{X: 15, Y: 15},
				Size:     vector.Vector2{X: 10, Y: 10},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rigidbody := Rigidbody{
				Hitbox: tc.box2,
			}

			rigidbody.MoveOutOfBox(tc.box)
			if rigidbody.Hitbox != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, tc.box2)
			}
		})
	}

}

func TestCollidesWith(t *testing.T) {
	testCases := []struct {
		name      string
		rigidbody Rigidbody
		box       box.Box
		expected  bool
	}{
		{
			name: "should return true when the rigidbody collides with the box",
			rigidbody: Rigidbody{
				Hitbox: box.Box{
					Position: vector.Vector2{X: 0, Y: 0},
					Size:     vector.Vector2{X: 10, Y: 10},
				},
			},
			box: box.Box{
				Position: vector.Vector2{X: 5, Y: 5},
				Size:     vector.Vector2{X: 10, Y: 10},
			},
			expected: true,
		},
		{
			name: "should return false when the rigidbody does not collide with the box",
			rigidbody: Rigidbody{
				Hitbox: box.Box{
					Position: vector.Vector2{X: 0, Y: 0},
					Size:     vector.Vector2{X: 10, Y: 10},
				},
			},
			box: box.Box{
				Position: vector.Vector2{X: 15, Y: 15},
				Size:     vector.Vector2{X: 10, Y: 10},
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.rigidbody.CollidesWith(tc.box)
			if result != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestRigidbody_ApplyAcceleration(t *testing.T) {
	rb := Rigidbody{
		Velocity:     vector.Vector2{X: 1, Y: 2},
		Acceleration: vector.Vector2{X: 3, Y: 4},
	}

	rb.ApplyAcceleration()

	expectedVelocity := vector.Vector2{X: 4, Y: 6}
	expectedAcceleration := vector.Vector2{X: 0, Y: 0}

	if rb.Velocity != expectedVelocity {
		t.Errorf("expected velocity %v, got %v", expectedVelocity, rb.Velocity)
	}

	if rb.Acceleration != expectedAcceleration {
		t.Errorf("expected acceleration %v, got %v", expectedAcceleration, rb.Acceleration)
	}
}

func TestRigidbody_AddForce(t *testing.T) {
	rb := Rigidbody{
		Acceleration: vector.Vector2{X: 0, Y: 0},
	}

	force := vector.Vector2{X: 2, Y: 3}
	expectedAcceleration := vector.Vector2{X: 2, Y: 3}

	rb.AddForce(force)

	if rb.Acceleration != expectedAcceleration {
		t.Errorf("expected acceleration %v, got %v", expectedAcceleration, rb.Acceleration)
	}
}

func TestRigidbody_ApplyVelocity(t *testing.T) {
	testCases := []struct {
		name          string
		initialPos    vector.Vector2
		velocity      vector.Vector2
		expectedFinal vector.Vector2
	}{
		{
			name:          "should apply velocity to the position",
			initialPos:    vector.Vector2{X: 0, Y: 0},
			velocity:      vector.Vector2{X: 5, Y: 5},
			expectedFinal: vector.Vector2{X: 5, Y: 5},
		},
		{
			name:          "should apply negative velocity to the position",
			initialPos:    vector.Vector2{X: 10, Y: 10},
			velocity:      vector.Vector2{X: -3, Y: -3},
			expectedFinal: vector.Vector2{X: 7, Y: 7},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rb := Rigidbody{
				Hitbox:   box.Box{Position: tc.initialPos},
				Velocity: tc.velocity,
			}

			rb.ApplyVelocity()

			if rb.Hitbox.Position != tc.expectedFinal {
				t.Errorf("expected final position %v, got %v", tc.expectedFinal, rb.Hitbox.Position)
			}
		})
	}
}

func TestRigidbody_LimitHorizontalVelocity(t *testing.T) {
	testCases := []struct {
		name        string
		velocity    vector.Vector2
		maxVelocity float64
		expected    vector.Vector2
	}{
		{
			name:        "should limit positive horizontal velocity",
			velocity:    vector.Vector2{X: 5, Y: 0},
			maxVelocity: 3,
			expected:    vector.Vector2{X: 3, Y: 0},
		},
		{
			name:        "should limit negative horizontal velocity",
			velocity:    vector.Vector2{X: -5, Y: 0},
			maxVelocity: 3,
			expected:    vector.Vector2{X: -3, Y: 0},
		},
		{
			name:        "should not modify velocity when within limits",
			velocity:    vector.Vector2{X: 2, Y: 0},
			maxVelocity: 3,
			expected:    vector.Vector2{X: 2, Y: 0},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rb := Rigidbody{
				Velocity: tc.velocity,
			}

			rb.LimitHorizontalVelocity(tc.maxVelocity)

			if rb.Velocity != tc.expected {
				t.Errorf("expected velocity %v, got %v", tc.expected, rb.Velocity)
			}
		})
	}
}

func TestRigidbody_LimitVerticalVelocity(t *testing.T) {
	testCases := []struct {
		name        string
		velocity    vector.Vector2
		maxVelocity float64
		expected    vector.Vector2
	}{
		{
			name:        "should limit positive velocity",
			velocity:    vector.Vector2{X: 0, Y: 5},
			maxVelocity: 3,
			expected:    vector.Vector2{X: 0, Y: 3},
		},
		{
			name:        "should limit negative velocity",
			velocity:    vector.Vector2{X: 0, Y: -5},
			maxVelocity: 3,
			expected:    vector.Vector2{X: 0, Y: -3},
		},
		{
			name:        "should not modify velocity within the limit",
			velocity:    vector.Vector2{X: 0, Y: 2},
			maxVelocity: 3,
			expected:    vector.Vector2{X: 0, Y: 2},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rb := Rigidbody{
				Velocity: tc.velocity,
			}

			rb.LimitVerticalVelocity(tc.maxVelocity)

			if rb.Velocity != tc.expected {
				t.Errorf("expected velocity %v, got %v", tc.expected, rb.Velocity)
			}
		})
	}
}

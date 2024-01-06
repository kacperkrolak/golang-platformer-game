package box

import (
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDisplacementVector(t *testing.T) {
	testCases := []struct {
		name         string
		box1         Box
		box2         Box
		wantedVector vector.Vector2
	}{
		{
			name: "box2 should be moved down when it's colliding with the box above",
			box1: Box{
				Position: vector.Vector2{X: 0, Y: 0},
				Size:     vector.Vector2{X: 10, Y: 10},
			},
			box2: Box{
				Position: vector.Vector2{X: 5, Y: 5},
				Size:     vector.Vector2{X: 10, Y: 10},
			},
			wantedVector: vector.Vector2{X: 0, Y: 5},
		},
		{
			name: "box2 should be moved right when it's colliding with the box on the left",
			box1: Box{
				Position: vector.Vector2{X: 0, Y: 0},
				Size:     vector.Vector2{X: 10, Y: 10},
			},
			box2: Box{
				Position: vector.Vector2{X: 5, Y: 0},
				Size:     vector.Vector2{X: 10, Y: 10},
			},
			wantedVector: vector.Vector2{X: 5, Y: 0},
		},
		{
			name: "box2 should be moved left when it's colliding with the box on the right",
			box1: Box{
				Position: vector.Vector2{X: 0, Y: 0},
				Size:     vector.Vector2{X: 10, Y: 10},
			},
			box2: Box{
				Position: vector.Vector2{X: -5, Y: 0},
				Size:     vector.Vector2{X: 10, Y: 10},
			},
			wantedVector: vector.Vector2{X: -5, Y: 0},
		},
		{
			name: "should return a correct displacement vector when boxes are not colliding",
			box1: Box{
				Position: vector.Vector2{X: 0, Y: 0},
				Size:     vector.Vector2{X: 10, Y: 10},
			},
			box2: Box{
				Position: vector.Vector2{X: 15, Y: 15},
				Size:     vector.Vector2{X: 10, Y: 10},
			},
			wantedVector: vector.Vector2{X: 0, Y: 0},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			displacementVector := tc.box1.DisplacementVector(tc.box2)
			if !cmp.Equal(displacementVector, tc.wantedVector) {
				t.Errorf("DisplacementVector() = %v, wanted %v", displacementVector, tc.wantedVector)
			}
		})
	}
}

func TestCollidesWith(t *testing.T) {
	testCases := []struct {
		name     string
		box1     Box
		box2     Box
		expected bool
	}{
		{
			name: "should return true when boxes are colliding",
			box1: Box{
				Position: vector.Vector2{X: 0, Y: 0},
				Size:     vector.Vector2{X: 10, Y: 10},
			},
			box2: Box{
				Position: vector.Vector2{X: 5, Y: 5},
				Size:     vector.Vector2{X: 10, Y: 10},
			},
			expected: true,
		},
		{
			name: "should return false when boxes are not colliding",
			box1: Box{
				Position: vector.Vector2{X: 0, Y: 0},
				Size:     vector.Vector2{X: 10, Y: 10},
			},
			box2: Box{
				Position: vector.Vector2{X: 15, Y: 15},
				Size:     vector.Vector2{X: 10, Y: 10},
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.box1.CollidesWith(tc.box2)
			if actual != tc.expected {
				t.Errorf("CollidesWith() = %v, expected %v", actual, tc.expected)
			}
		})
	}
}

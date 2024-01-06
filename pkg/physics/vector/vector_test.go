package vector

import (
	"testing"
)

func TestVector2_Add(t *testing.T) {
	testCases := []struct {
		name     string
		vector   Vector2
		vector2  Vector2
		expected Vector2
	}{
		{
			name:     "should add two vectors",
			vector:   Vector2{X: 1, Y: 1},
			vector2:  Vector2{X: 2, Y: 2},
			expected: Vector2{X: 3, Y: 3},
		},
		{
			name:     "should add two vectors with negative values",
			vector:   Vector2{X: -1, Y: -1},
			vector2:  Vector2{X: -2, Y: -2},
			expected: Vector2{X: -3, Y: -3},
		},
		{
			name:     "should add two vectors with different values",
			vector:   Vector2{X: 1, Y: 2},
			vector2:  Vector2{X: 3, Y: 4},
			expected: Vector2{X: 4, Y: 6},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.vector.Add(tc.vector2)
			if tc.vector != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, tc.vector)
			}
		})
	}
}

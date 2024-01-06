package vector

import (
	"math"
	"testing"
)

func TestSmoothDamp1D(t *testing.T) {
	testCases := []struct {
		name            string
		current         float64
		target          float64
		currentVelocity float64
		smoothTime      float64
		deltaTime       float64
		wantErr         bool
	}{
		{
			name:            "should approach the target value",
			current:         10.0,
			target:          20.0,
			currentVelocity: 0.0,
			smoothTime:      1.0,
			deltaTime:       0.1,
			wantErr:         false,
		},
		{
			name:            "shouldn't move if the target value is reached",
			current:         20.0,
			target:          20.0,
			currentVelocity: 0.0,
			smoothTime:      1.0,
			deltaTime:       0.1,
			wantErr:         false,
		},
		{
			name:            "shouldn't work for smoothTime <= 0",
			current:         10.0,
			target:          20.0,
			currentVelocity: 0.0,
			smoothTime:      0.0,
			deltaTime:       0.1,
			wantErr:         true,
		},
		{
			name:            "shouldn't work for deltaTime <= 0",
			current:         10.0,
			target:          20.0,
			currentVelocity: 0.0,
			smoothTime:      1.0,
			deltaTime:       0.0,
			wantErr:         true,
		},
	}

	// smoothDamp1D is a function which doesn't need to return an exact value.
	// It's more 'art' than 'science' and is likely to change so we will only
	// check the most basic properties of it.
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			position, velocity, err := smoothDamp1D(tc.current, tc.target, tc.currentVelocity, tc.smoothTime, tc.deltaTime)
			if (err != nil) != tc.wantErr {
				t.Errorf("smoothDamp1D() error = %v, wantErr %v", err, tc.wantErr)
			}

			if tc.wantErr {
				return
			}

			// The position should be between the current and target values.
			if position < math.Min(tc.current, tc.target) || position > math.Max(tc.current, tc.target) {
				t.Errorf("smoothDamp1D() position should be between current and target: position = %v, current = %v, target = %v", position, tc.current, tc.target)
			}

			// The velocity should be 0 when the target is reached and not 0 otherwise.
			if (tc.target == position) != (velocity == 0.0) {
				t.Errorf("smoothDamp1D() invalid velocity: velocity = %v, target = %v, position = %v", velocity, tc.target, position)
			}
		})
	}
}

func TestClamp(t *testing.T) {
	testCases := []struct {
		name string
		f    float64
		low  float64
		high float64
		want float64
	}{
		{
			name: "should return the same value if within the range",
			f:    5.0,
			low:  0.0,
			high: 10.0,
			want: 5.0,
		},
		{
			name: "should return the low value if f is less than low",
			f:    -5.0,
			low:  0.0,
			high: 10.0,
			want: 0.0,
		},
		{
			name: "should return the high value if f is greater than high",
			f:    15.0,
			low:  0.0,
			high: 10.0,
			want: 10.0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := clamp(tc.f, tc.low, tc.high)
			if got != tc.want {
				t.Errorf("clamp() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestSmoothDamp(t *testing.T) {
	testCases := []struct {
		name            string
		current         Vector2
		target          Vector2
		currentVelocity Vector2
		smoothTime      float64
		deltaTime       float64
		wantErr         bool
	}{
		{
			name:            "should approach the target value",
			current:         Vector2{X: 10.0, Y: 10.0},
			target:          Vector2{X: 20.0, Y: 20.0},
			currentVelocity: Vector2{X: 0.0, Y: 0.0},
			smoothTime:      1.0,
			deltaTime:       0.1,
			wantErr:         false,
		},
		{
			name:            "shouldn't move if the target value is reached",
			current:         Vector2{X: 20.0, Y: 20.0},
			target:          Vector2{X: 20.0, Y: 20.0},
			currentVelocity: Vector2{X: 0.0, Y: 0.0},
			smoothTime:      1.0,
			deltaTime:       0.1,
			wantErr:         false,
		},
		{
			name:            "shouldn't work for smoothTime <= 0",
			current:         Vector2{X: 10.0, Y: 10.0},
			target:          Vector2{X: 20.0, Y: 20.0},
			currentVelocity: Vector2{X: 0.0, Y: 0.0},
			smoothTime:      0.0,
			deltaTime:       0.1,
			wantErr:         true,
		},
		{
			name:            "shouldn't work for deltaTime <= 0",
			current:         Vector2{X: 10.0, Y: 10.0},
			target:          Vector2{X: 20.0, Y: 20.0},
			currentVelocity: Vector2{X: 0.0, Y: 0.0},
			smoothTime:      1.0,
			deltaTime:       0.0,
			wantErr:         true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			position, velocity, err := SmoothDamp(tc.current, tc.target, tc.currentVelocity, tc.smoothTime, tc.deltaTime)
			if (err != nil) != tc.wantErr {
				t.Errorf("SmoothDamp() error = %v, wantErr %v", err, tc.wantErr)
			}

			if tc.wantErr {
				return
			}

			// The position should be between the current and target values.
			if position.X < math.Min(tc.current.X, tc.target.X) || position.X > math.Max(tc.current.X, tc.target.X) {
				t.Errorf("SmoothDamp() position.X should be between current.X and target.X: position.X = %v, current.X = %v, target.X = %v", position.X, tc.current.X, tc.target.X)
			}
			if position.Y < math.Min(tc.current.Y, tc.target.Y) || position.Y > math.Max(tc.current.Y, tc.target.Y) {
				t.Errorf("SmoothDamp() position.Y should be between current.Y and target.Y: position.Y = %v, current.Y = %v, target.Y = %v", position.Y, tc.current.Y, tc.target.Y)
			}

			// The velocity should be 0 when the target is reached and not 0 otherwise.
			if (tc.target == position) != (velocity == Vector2{X: 0.0, Y: 0.0}) {
				t.Errorf("SmoothDamp() invalid velocity: velocity = %v, target = %v, position = %v", velocity, tc.target, position)
			}
		})
	}
}

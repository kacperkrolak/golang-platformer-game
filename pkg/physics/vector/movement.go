package vector

import (
	"fmt"
	"math"
)

func clamp(f, low, high float64) float64 {
	if f < low {
		return low
	}
	if f > high {
		return high
	}
	return f
}

// https://stackoverflow.com/questions/61372498/how-does-mathf-smoothdamp-work-what-is-it-algorithm
func smoothDamp1D(current, target, currentVelocity, smoothTime, deltaTime float64) (float64, float64, error) {
	if deltaTime <= 0.0 {
		return 0.0, 0.0, fmt.Errorf("deltaTime must be positive")
	}

	if smoothTime <= 0.0 {
		return 0.0, 0.0, fmt.Errorf("smoothTime must be positive")
	}

	maxSpeed := math.MaxFloat64

	smoothTime = math.Max(0.0001, smoothTime)
	omega := 2.0 / smoothTime

	x := omega * deltaTime
	exp := 1.0 / (1.0 + x + 0.48*x*x + 0.235*x*x*x)
	change := current - target
	originalTo := target

	maxChange := maxSpeed * smoothTime
	change = clamp(change, -maxChange, maxChange)

	target = current - change

	temp := (currentVelocity + omega*change) * deltaTime
	currentVelocity = (currentVelocity - omega*temp) * exp
	output := target + (change+temp)*exp

	if (originalTo-current > 0.0) == (output > originalTo) {
		output = originalTo
		currentVelocity = (output - originalTo) / deltaTime
	}

	return output, currentVelocity, nil
}

func SmoothDamp(current, target, currentVelocity Vector2, smoothTime, deltaTime float64) (Vector2, Vector2, error) {
	posX, velX, err := smoothDamp1D(current.X, target.X, currentVelocity.X, smoothTime, deltaTime)
	if err != nil {
		return Vector2{}, Vector2{}, err
	}

	posY, velY, err := smoothDamp1D(current.Y, target.Y, currentVelocity.Y, smoothTime, deltaTime)
	if err != nil {
		return Vector2{}, Vector2{}, err
	}

	return Vector2{X: posX, Y: posY}, Vector2{X: velX, Y: velY}, nil
}

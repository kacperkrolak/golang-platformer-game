package input

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// Return -1 for left, 1 for right, 0 for none.
// It is float to leave space for implementing joystick devices.
func GetHorizontal() float64 {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		return -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		return 1
	}
	return 0
}

// Return true if user doesn't want to move horizontally.
func IsHorizontalIdle() bool {
	return math.Abs(GetHorizontal()) < 0.01
}
func IsJumpPressed() bool {
	return ebiten.IsKeyPressed(ebiten.KeySpace) || ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW)
}

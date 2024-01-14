/*
Package input encapsulates user input handling.
*/
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
//
// On keyboard, it's when left and right arrow keys are not pressed.
// On joystick, it's when the stick is in the middle or almost in the middle.
func IsHorizontalIdle() bool {
	return math.Abs(GetHorizontal()) < 0.01
}

// Check if any of the common jump keys is pressed.
//
// For now, it's space, up arrow and W.
func IsJumpPressed() bool {
	return ebiten.IsKeyPressed(ebiten.KeySpace) || ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW)
}

/*
Package box represents a rectangular area.

This package assumes that the Y-axis is pointing down.
*/
package box

import (
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"
)

type Box struct {
	Position vector.Vector2
	Size     vector.Vector2
}

func (b Box) Left() float64 {
	return b.Position.X
}

func (b Box) Right() float64 {
	return b.Position.X + b.Size.X
}

func (b Box) Top() float64 {
	return b.Position.Y
}

func (b Box) Bottom() float64 {
	return b.Position.Y + b.Size.Y
}

func (b Box) Center() vector.Vector2 {
	return vector.Vector2{
		X: b.Position.X + b.Size.X/2,
		Y: b.Position.Y + b.Size.Y/2,
	}
}

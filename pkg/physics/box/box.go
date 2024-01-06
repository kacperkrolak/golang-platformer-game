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

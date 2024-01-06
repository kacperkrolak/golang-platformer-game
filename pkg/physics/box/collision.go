package box

import (
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"
	"math"
)

func (b Box) CollidesWith(b2 Box) bool {
	if b.Top() >= b2.Bottom() || b.Bottom() <= b2.Top() {
		return false
	}

	if b.Left() >= b2.Right() || b.Right() <= b2.Left() {
		return false
	}

	return true
}

// Vector by which b2 should be moved to not collide with b.
func (b Box) DisplacementVector(b2 Box) vector.Vector2 {
	displacement := vector.Vector2{X: 0, Y: 0}
	if !b.CollidesWith(b2) {
		return displacement
	}

	if b.Top() < b2.Top() {
		displacement.Y = b.Bottom() - b2.Top()
	} else {
		displacement.Y = -(b2.Bottom() - b.Top())
	}

	if b.Left() < b2.Left() {
		displacement.X = b.Right() - b2.Left()
	} else {
		displacement.X = -(b2.Right() - b.Left())
	}

	// We don't want to change both X and Y at the same time, so we choose the smaller one
	if math.Abs(displacement.X) < math.Abs(displacement.Y) {
		displacement.Y = 0
	} else {
		displacement.X = 0
	}

	return displacement
}

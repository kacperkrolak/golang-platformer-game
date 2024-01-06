package camera

import (
	"fmt"
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"
)

type Camera struct {
	Position   vector.Vector2
	Velocity   vector.Vector2
	Target     Targetable
	SmoothTime float64
}

func (c *Camera) Update(tps float64) error {
	target := c.Target
	if target == nil {
		return fmt.Errorf("camera target is nil")
	}

	targetPosition := target.GetRigidbody().Hitbox.Position

	position, velocity, err := vector.SmoothDamp(c.Position, targetPosition, c.Velocity, c.SmoothTime, tps/60)
	if err != nil {
		return err
	}

	c.Position, c.Velocity = position, velocity

	return nil
}

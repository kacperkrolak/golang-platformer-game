package camera

import (
	"kacperkrolak/golang-platformer-game/pkg/physics/rigidbody"
)

type Targetable interface {
	GetRigidbody() rigidbody.Rigidbody
}

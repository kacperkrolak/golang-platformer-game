package mock

import (
	"kacperkrolak/golang-platformer-game/pkg/physics/rigidbody"
)

type Target struct {
	Rigidbody rigidbody.Rigidbody
}

func (t Target) GetRigidbody() rigidbody.Rigidbody {
	return t.Rigidbody
}

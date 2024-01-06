package particle

import (
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"

	"github.com/hajimehoshi/ebiten/v2"
)

type Particle interface {
	Update()
	IsDead() bool
	Draw(screen *ebiten.Image, cameraOffset vector.Vector2)
}

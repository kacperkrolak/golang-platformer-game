package gamemap

import (
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"

	"github.com/hajimehoshi/ebiten/v2"
)

func (m Map) Draw(screen *ebiten.Image, cameraOffset vector.Vector2, img *ebiten.Image, tileSize int) {
	for i, row := range m.Blocks {
		for j, block := range row {
			position := vector.Vector2{X: float64(j * tileSize), Y: float64(i * tileSize)}
			block.Draw(screen, cameraOffset.Added(position), img, tileSize)
		}
	}
}

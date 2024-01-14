package tilemap

import (
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"

	"github.com/hajimehoshi/ebiten/v2"
)

// Draw each tile in the map in positions relative to the camera.
func (m Map) Draw(screen *ebiten.Image, cameraOffset vector.Vector2, img *ebiten.Image, tileSize int) {
	for i, t := range m.Tiles {
		for j, tile := range t {
			position := vector.Vector2{X: float64(j * tileSize), Y: float64(i * tileSize)}
			tile.Draw(screen, cameraOffset.Added(position), img, tileSize)
		}
	}
}

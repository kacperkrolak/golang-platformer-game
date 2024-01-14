package mock

import (
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"
	"kacperkrolak/golang-platformer-game/pkg/world/tilemap/tile"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/stretchr/testify/mock"
)

type MockTile struct {
	mock.Mock
}

func (m *MockTile) Draw(screen *ebiten.Image, cameraOffset vector.Vector2, tileSheet *ebiten.Image, tileSize int) {
	m.Called(screen, cameraOffset, tileSheet, tileSize)
}

func (m *MockTile) GetGroup() string {
	return ""
}

func (m *MockTile) UpdateVariant(arg0 [4]tile.Tile) {
	m.Called(arg0)
}

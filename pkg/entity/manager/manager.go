package manager

import (
	"kacperkrolak/golang-platformer-game/pkg/entity"
	"kacperkrolak/golang-platformer-game/pkg/entity/coin"
	"kacperkrolak/golang-platformer-game/pkg/physics/box"
	"kacperkrolak/golang-platformer-game/pkg/physics/rigidbody"
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"
	"kacperkrolak/golang-platformer-game/pkg/visuals/particle"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Manager struct {
	particleSystem *particle.ParticleSystem
	entities       []entity.Entity
	tilesImage     *ebiten.Image
}

func NewManager(particleSystem *particle.ParticleSystem, tilesImage *ebiten.Image) *Manager {
	return &Manager{
		particleSystem: particleSystem,
		tilesImage:     tilesImage,
	}
}

const (
	COIN = iota
)

func (m *Manager) Spawn(position vector.Vector2, entityType int) error {
	switch entityType {
	case COIN:
		m.entities = append(m.entities, &coin.Coin{
			Rigidbody: rigidbody.Rigidbody{
				Hitbox: box.Box{
					Position: position,
					Size:     coin.COIN_SIZE,
				},
			},
			TileSheet:      m.tilesImage,
			ParticleSystem: m.particleSystem,
		})
	}
	return nil
}

func (m *Manager) Update(deltaTime time.Duration, playerRigidbody *rigidbody.Rigidbody) {
	for i := len(m.entities) - 1; i >= 0; i-- {
		e := m.entities[i]
		e.Update(deltaTime)
		if playerRigidbody.Hitbox.CollidesWith(e.GetRigidbody().Hitbox) {
			e.OnCollisionWithPlayer()
		}
		if e.IsDead() {
			m.entities = append(m.entities[:i], m.entities[i+1:]...)
		}
	}
}

func (m *Manager) Draw(screen *ebiten.Image, cameraOffset vector.Vector2) {
	for _, e := range m.entities {
		e.Draw(screen, cameraOffset)
	}
}

func (m *Manager) Reset() {
	m.entities = []entity.Entity{}
}

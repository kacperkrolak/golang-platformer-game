package particle

import (
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"

	"github.com/hajimehoshi/ebiten/v2"
)

type ParticleSystem struct {
	particles []Particle
}

func (system *ParticleSystem) AddParticles(particles []Particle) {
	system.particles = append(system.particles, particles...)
}

func (system *ParticleSystem) Update() {
	for i := len(system.particles) - 1; i >= 0; i-- {
		system.particles[i].Update()
		if system.particles[i].IsDead() {
			system.particles = append(system.particles[:i], system.particles[i+1:]...)
		}
	}
}

func (system *ParticleSystem) Draw(screen *ebiten.Image, cameraOffset vector.Vector2) {
	for _, p := range system.particles {
		p.Draw(screen, cameraOffset)
	}
}

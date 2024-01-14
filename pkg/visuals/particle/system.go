/*
Package particle implements particle system.

It can be used to create effects like smoke, fire, explosions, etc.
*/
package particle

import (
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"

	"github.com/hajimehoshi/ebiten/v2"
)

// ParticleSystem manages storing, updating and drawing particles.
type ParticleSystem struct {
	particles []Particle
}

// Add Particles to the system.
func (system *ParticleSystem) AddParticles(particles []Particle) {
	system.particles = append(system.particles, particles...)
}

// Trigger updates on all particles and removes dead ones.
func (system *ParticleSystem) Update() {
	for i := len(system.particles) - 1; i >= 0; i-- {
		system.particles[i].Update()
		if system.particles[i].IsDead() {
			system.particles = append(system.particles[:i], system.particles[i+1:]...)
		}
	}
}

// Draw all particles.
func (system *ParticleSystem) Draw(screen *ebiten.Image, cameraOffset vector.Vector2) {
	for _, p := range system.particles {
		p.Draw(screen, cameraOffset)
	}
}

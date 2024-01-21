// Package confetti creates and manages confetti particle effects.
package confetti

import (
	"image/color"
	"kacperkrolak/golang-platformer-game/pkg/physics/box"
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"
	"kacperkrolak/golang-platformer-game/pkg/visuals/particle"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Particle struct {
	Box      box.Box
	Velocity vector.Vector2
	Life     int
	MaxLife  int
	Color    color.Color
}

func (p *Particle) Update() {
	p.Box.Position.Add(p.Velocity)
	p.Life--
}

func (p *Particle) IsDead() bool {
	return p.Life <= 0
}

func (p *Particle) Draw(screen *ebiten.Image, cameraOffset vector.Vector2) {
	image := ebiten.NewImage(int(p.Box.Size.X), int(p.Box.Size.Y))
	image.Fill(p.Color)
	op := &ebiten.DrawImageOptions{}

	alivePercent := float32(p.Life) / float32(p.MaxLife)
	op.GeoM.Rotate(math.Atan2(p.Velocity.Y, p.Velocity.X) + float64(alivePercent))
	op.GeoM.Scale(float64(alivePercent), float64(alivePercent))
	op.GeoM.Translate(p.Box.Position.X+cameraOffset.X, p.Box.Position.Y+cameraOffset.Y)

	op.ColorScale.ScaleAlpha(alivePercent * alivePercent)
	screen.DrawImage(image, op)
}

func CreateEffect(position vector.Vector2, count uint, radius int, speed float64, life int) []particle.Particle {
	particles := make([]particle.Particle, count)
	for i := uint(0); i < count; i++ {
		velocity := vector.Vector2{
			X: rand.Float64() - 0.5,
			Y: rand.Float64() - 0.5,
		}

		position := vector.Vector2{
			X: position.X + rand.Float64()*float64(radius),
			Y: position.Y + rand.Float64()*float64(radius),
		}

		color := color.RGBA{
			R: uint8(rand.Intn(155) + 100),
			G: uint8(rand.Intn(155) + 100),
			B: uint8(rand.Intn(155) + 100),
			A: 255,
		}

		p := Particle{
			Box: box.Box{
				Position: position,
				Size:     vector.Vector2{X: 10, Y: 3},
			},
			Velocity: velocity.Normalized().Scaled(speed),
			Life:     life,
			MaxLife:  life,
			Color:    color,
		}
		particles[i] = &p
	}

	return particles
}

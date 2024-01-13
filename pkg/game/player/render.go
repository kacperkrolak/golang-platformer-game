package player

import (
	"image"
	"image/color"
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var RUN_1 = vector.Vector2{X: 9, Y: 42}
var RUN_2 = vector.Vector2{X: 41, Y: 41}

func (p *Player) Draw(screen *ebiten.Image, offsetX float64, offsetY float64, characterImage *ebiten.Image, tileSize int) {
	isMoving := math.Abs(p.Rigidbody.Velocity.X) > 0.05
	frame := RUN_1
	if p.State == 1 && isMoving {
		frame = RUN_2
	}

	characterFrame := characterImage.SubImage(image.Rect(int(frame.X), int(frame.Y), int(frame.X)+14, int(frame.Y)+21)).(*ebiten.Image)

	op := &ebiten.DrawImageOptions{}
	if !p.FacingRight {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(14, 0)
	}

	op.GeoM.Translate(offsetX+p.Rigidbody.Hitbox.Left(), offsetY+p.Rigidbody.Hitbox.Top())

	screen.DrawImage(characterFrame, op)

	surfaceDet := p.SurfaceDetector()

	// Draw surface detector
	ebitenutil.DrawRect(screen, offsetX+surfaceDet.Left(), offsetY+surfaceDet.Top(), surfaceDet.Size.X, surfaceDet.Size.Y, color.RGBA{255, 0, 0, 255})
}

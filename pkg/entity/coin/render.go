package coin

import (
	"image"
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"

	"github.com/hajimehoshi/ebiten/v2"
)

// Position of texture for each frame of rotating animation.
var FRAMES = []vector.Vector2{
	{X: 7 * 16, Y: 7 * 16},
	{X: 8 * 16, Y: 7 * 16},
	{X: 9 * 16, Y: 7 * 16},
	{X: 10 * 16, Y: 7 * 16},
}

var COIN_SIZE = vector.Vector2{X: 16, Y: 16}

func (c *Coin) Draw(screen *ebiten.Image, cameraOffset vector.Vector2) {
	time := int(c.TimePassed.Milliseconds() / 125)
	frame := FRAMES[time%len(FRAMES)]

	coinTexture := c.TileSheet.SubImage(image.Rect(int(frame.X), int(frame.Y), int(frame.X)+int(COIN_SIZE.X), int(frame.Y)+int(COIN_SIZE.Y))).(*ebiten.Image)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(cameraOffset.X+c.Rigidbody.Hitbox.Left(), cameraOffset.Y+c.Rigidbody.Hitbox.Top())

	screen.DrawImage(coinTexture, op)
}

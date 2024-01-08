package ground

import (
	"image"
	"kacperkrolak/golang-platformer-game/pkg/gamemap/tile"
	"kacperkrolak/golang-platformer-game/pkg/physics/box"
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"

	"github.com/hajimehoshi/ebiten/v2"
)

type Tile struct {
	variant uint8
	hitbox  box.Box
}

func (t Tile) IsCollidable() bool {
	return true
}

func (t Tile) IsSolid() bool {
	return true
}

var TilePositions map[uint8]vector.Vector2 = map[uint8]vector.Vector2{
	tile.TOP:               {X: 176, Y: 0},
	tile.BOTTOM:            {X: 208, Y: 0},
	tile.TOP_BOTTOM:        {X: 176, Y: 0},
	tile.LEFT:              {X: 160, Y: 16},
	tile.RIGHT:             {X: 192, Y: 16},
	tile.LEFT_RIGHT:        {X: 208, Y: 0},
	tile.SURROUNDED:        {X: 208, Y: 0},
	tile.ALONE:             {X: 176, Y: 16},
	tile.LEFT_TOP:          {X: 160, Y: 0},
	tile.LEFT_BOTTOM:       {X: 160, Y: 16},
	tile.RIGHT_TOP:         {X: 192, Y: 0},
	tile.RIGHT_BOTTOM:      {X: 192, Y: 16},
	tile.LEFT_TOP_RIGHT:    {X: 208, Y: 0},
	tile.BOTTOM_LEFT_RIGHT: {X: 208, Y: 0},
	tile.BOTTOM_LEFT_TOP:   {X: 160, Y: 0},
	tile.BOTTOM_RIGHT_TOP:  {X: 192, Y: 0},
}

func (t Tile) Draw(screen *ebiten.Image, screenPosition vector.Vector2, tileSheet *ebiten.Image, tileSize int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(screenPosition.X, screenPosition.Y)

	tilePos := TilePositions[t.variant]
	tileImg := tileSheet.SubImage(image.Rect(int(tilePos.X), int(tilePos.Y), int(tilePos.X)+tileSize, int(tilePos.Y)+tileSize)).(*ebiten.Image)

	screen.DrawImage(tileImg, op)
}

func (t *Tile) UpdateVariant(neighbours [4]tile.Tile) {
	variant := uint8(0)
	if neighbours[0] != nil && !neighbours[0].IsSolid() {
		variant |= tile.LEFT
	}
	if neighbours[1] != nil && !neighbours[1].IsSolid() {
		variant |= tile.TOP
	}
	if neighbours[2] != nil && !neighbours[2].IsSolid() {
		variant |= tile.RIGHT
	}
	if neighbours[3] != nil && !neighbours[3].IsSolid() {
		variant |= tile.BOTTOM
	}

	t.variant = variant
}

func (t Tile) Hitbox() box.Box {
	return t.hitbox
}

func (t *Tile) SetHitbox(hitbox box.Box) {
	t.hitbox = hitbox
}

func (t Tile) IsDeadly() bool {
	return false
}

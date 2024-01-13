package spikes

import (
	"image"
	"kacperkrolak/golang-platformer-game/pkg/physics/box"
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"
	"kacperkrolak/golang-platformer-game/pkg/world/tilemap/tile"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Tile struct {
	variant uint8
	hitbox  box.Box
}

var TilePosition vector.Vector2 = vector.Vector2{X: 16, Y: 47}

func (t Tile) Draw(screen *ebiten.Image, screenPosition vector.Vector2, tileSheet *ebiten.Image, tileSize int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(screenPosition.X, screenPosition.Y)

	tileImg := tileSheet.SubImage(image.Rect(int(TilePosition.X), int(TilePosition.Y), int(TilePosition.X)+tileSize, int(TilePosition.Y)+tileSize)).(*ebiten.Image)
	tileImg = t.preprocessTileImage(tileImg, tileSize)

	screen.DrawImage(tileImg, op)
}

func (t Tile) preprocessTileImage(tileImg *ebiten.Image, tileSize int) *ebiten.Image {
	op := &ebiten.DrawImageOptions{}
	if t.variant&tile.TOP != 0 && t.variant&tile.BOTTOM == 0 {
		return tileImg
	} else if t.variant&tile.TOP == 0 && t.variant&tile.BOTTOM != 0 {
		op.GeoM.Scale(1, -1)
		op.GeoM.Translate(0, float64(tileSize))
	} else if t.variant&tile.LEFT == 0 && t.variant&tile.RIGHT != 0 {
		op.GeoM.Rotate(math.Pi / 2)
		op.GeoM.Translate(float64(tileSize), 0)
	} else if t.variant&tile.RIGHT == 0 && t.variant&tile.LEFT != 0 {
		op.GeoM.Rotate(-math.Pi / 2)
		op.GeoM.Translate(0, float64(tileSize))
	} else {
		return tileImg
	}

	rotatedTileImg := ebiten.NewImage(tileSize, tileSize)
	rotatedTileImg.DrawImage(tileImg, op)
	return rotatedTileImg
}

func (t *Tile) UpdateVariant(neighbours [4]tile.Tile) {
	variant := uint8(0)
	if neighbours[0] != nil && neighbours[0].GetGroup() != "solid" {
		variant |= tile.LEFT
	}
	if neighbours[1] != nil && neighbours[1].GetGroup() != "solid" {
		variant |= tile.TOP
	}
	if neighbours[2] != nil && neighbours[2].GetGroup() != "solid" {
		variant |= tile.RIGHT
	}
	if neighbours[3] != nil && neighbours[3].GetGroup() != "solid" {
		variant |= tile.BOTTOM
	}

	t.variant = variant
}

func (t Tile) GetGroup() string {
	return "obstacle"
}

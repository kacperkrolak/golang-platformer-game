package gamemap

import (
	"fmt"
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Position struct {
	X int
	Y int
}

var Dirt map[uint8]Position = map[uint8]Position{
	TOP:               {X: 176, Y: 0},
	BOTTOM:            {X: 208, Y: 0},
	TOP_BOTTOM:        {X: 176, Y: 0},
	LEFT:              {X: 160, Y: 16},
	RIGHT:             {X: 192, Y: 16},
	LEFT_RIGHT:        {X: 208, Y: 0},
	SURROUNDED:        {X: 208, Y: 0},
	ALONE:             {X: 176, Y: 16},
	LEFT_TOP:          {X: 160, Y: 0},
	LEFT_BOTTOM:       {X: 160, Y: 16},
	RIGHT_TOP:         {X: 192, Y: 0},
	RIGHT_BOTTOM:      {X: 192, Y: 16},
	LEFT_TOP_RIGHT:    {X: 208, Y: 0},
	BOTTOM_LEFT_RIGHT: {X: 208, Y: 0},
	BOTTOM_LEFT_TOP:   {X: 160, Y: 0},
	BOTTOM_RIGHT_TOP:  {X: 192, Y: 0},
}

var Spikes Position = Position{X: 16, Y: 47}

func GetTileTexture(tile Tile, img *ebiten.Image, tileSize int) (*ebiten.Image, error) {
	var pos Position
	switch {
	case tile.Type == DIRT:
		pos = Dirt[tile.Variant]
	case tile.Type == SPIKES:
		pos = Spikes
	default:
		return nil, fmt.Errorf("no texture for this tile type")
	}

	tileImg := img.SubImage(image.Rect(pos.X, pos.Y, pos.X+tileSize, pos.Y+tileSize)).(*ebiten.Image)

	if tile.Type == SPIKES {
		op := &ebiten.DrawImageOptions{}
		if tile.Variant&TOP != 0 && tile.Variant&BOTTOM == 0 {
			return tileImg, nil
		} else if tile.Variant&TOP == 0 && tile.Variant&BOTTOM != 0 {
			op.GeoM.Scale(1, -1)
			op.GeoM.Translate(0, float64(tileSize))
		} else if tile.Variant&LEFT == 0 && tile.Variant&RIGHT != 0 {
			op.GeoM.Rotate(math.Pi / 2)
			op.GeoM.Translate(float64(tileSize), 0)
		} else if tile.Variant&RIGHT == 0 && tile.Variant&LEFT != 0 {
			op.GeoM.Rotate(-math.Pi / 2)
			op.GeoM.Translate(0, float64(tileSize))
		} else {
			return tileImg, nil
		}

		rotatedTileImg := ebiten.NewImage(tileSize, tileSize)
		rotatedTileImg.DrawImage(tileImg, op)
		return rotatedTileImg, nil
	}

	return tileImg, nil
}

func (m Map) Draw(screen *ebiten.Image, offsetX float64, offsetY float64, img *ebiten.Image, tileSize int) {
	for i, t := range m.Tiles {
		for j, tile := range t {
			op := &ebiten.DrawImageOptions{}
			xOnMap, yOnMap := float64(j*tileSize), float64(i*tileSize)
			op.GeoM.Translate(xOnMap+offsetX, yOnMap+offsetY)

			img, err := GetTileTexture(tile, img, tileSize)
			if err == nil {
				screen.DrawImage(img, op)
			}
		}
	}
}

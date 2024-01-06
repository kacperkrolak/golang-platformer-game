package gamemap

import (
	"fmt"
	"image"

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

func GetTileTexture(tile Tile, img *ebiten.Image, tileSize int) (*ebiten.Image, error) {
	var sx, sy int
	if tile.Type == DIRT {
		pos := Dirt[tile.Variant]
		sx = pos.X
		sy = pos.Y
	} else {
		return nil, fmt.Errorf("no texture for this tile type")
	}

	return img.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image), nil
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

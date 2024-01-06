package game

import (
	"fmt"
	"kacperkrolak/golang-platformer-game/pkg/game/camera"
	"kacperkrolak/golang-platformer-game/pkg/game/player"
	"kacperkrolak/golang-platformer-game/pkg/gamemap"
	"kacperkrolak/golang-platformer-game/pkg/physics/box"
	"kacperkrolak/golang-platformer-game/pkg/physics/rigidbody"
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 520
)

type Game struct {
	gameMap        gamemap.Map
	tilesImage     *ebiten.Image
	characterImage *ebiten.Image
	tileSize       int
	player         *player.Player
	camera         camera.Camera
}

func MakeGame(mapFile string, textureFile string, characterFile string) Game {
	gameMap, err := loadGameMap(mapFile)
	if err != nil {
		log.Fatal(err)
	}

	tilesImage, err := loadTilesImage(textureFile)
	if err != nil {
		log.Fatal(err)
	}

	characterImage, err := loadTilesImage(characterFile)
	if err != nil {
		log.Fatal(err)
	}

	player := player.Player{
		Rigidbody: rigidbody.Rigidbody{
			Hitbox: box.Box{
				Position: vector.Vector2{X: 50, Y: 0},
				Size:     vector.Vector2{X: 14, Y: 21},
			},
		},
		Speed: 10 * 16,
	}
	return Game{
		gameMap:        gameMap,
		tilesImage:     tilesImage,
		characterImage: characterImage,
		tileSize:       16,
		player:         &player,
		camera: camera.Camera{
			Position:   vector.Vector2{X: 0, Y: 0},
			Velocity:   vector.Vector2{X: 0, Y: 0},
			Target:     &player,
			SmoothTime: 15,
		},
	}
}

func (g *Game) Update() error {
	// Update is run 60 times a second by default
	tps := float64(60)

	surfaceDetector := g.player.SurfaceDetector()
	g.player.Grounded = false
	for _, row := range g.gameMap.Tiles {
		for _, t := range row {
			if !t.IsCollidable() {
				continue
			}

			if surfaceDetector.CollidesWith(t.Hitbox) {
				g.player.Grounded = true
				break
			}
		}
	}

	g.player.Update(tps, g.tileSize)

	for _, row := range g.gameMap.Tiles {
		for _, t := range row {
			if !t.IsCollidable() {
				continue
			}

			g.player.Rigidbody.MoveOutOfBox(t.Hitbox)
		}
	}

	err := g.camera.Update(tps)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// getScreenPosition converts map position to screen position
func (g *Game) getScreenPosition(x, y float64) (float64, float64) {
	return ScreenWidth/2 - g.camera.Position.X, ScreenHeight/2 - g.camera.Position.Y
}

func (g *Game) Draw(screen *ebiten.Image) {
	offsetX, offsetY := g.getScreenPosition(0, 0)
	g.gameMap.Draw(screen, offsetX, offsetY, g.tilesImage, g.tileSize)
	g.player.Draw(screen, offsetX, offsetY, g.characterImage, g.tileSize)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("POS: %f %f Camera: %f %f", g.player.Rigidbody.Hitbox.Left(), g.player.Rigidbody.Hitbox.Top(), g.camera.Position.X, g.camera.Position.Y))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

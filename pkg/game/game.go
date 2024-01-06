package game

import (
	"fmt"
	"kacperkrolak/golang-platformer-game/pkg/game/camera"
	"kacperkrolak/golang-platformer-game/pkg/game/player"
	"kacperkrolak/golang-platformer-game/pkg/gamemap"
	"kacperkrolak/golang-platformer-game/pkg/physics/box"
	"kacperkrolak/golang-platformer-game/pkg/physics/rigidbody"
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"
	"kacperkrolak/golang-platformer-game/pkg/visuals/particle"
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
	particleSystem *particle.ParticleSystem
}

const TileSize = 16

func MakeGame(mapFile string, textureFile string, characterFile string) Game {
	gameMap, err := loadGameMap(mapFile, TileSize)
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
		Speed: 10 * TileSize,
	}
	return Game{
		gameMap:        gameMap,
		tilesImage:     tilesImage,
		characterImage: characterImage,
		tileSize:       TileSize,
		player:         &player,
		camera: camera.Camera{
			Position:   vector.Vector2{X: 0, Y: 0},
			Velocity:   vector.Vector2{X: 0, Y: 0},
			Target:     &player,
			SmoothTime: 15,
		},
		particleSystem: &particle.ParticleSystem{},
	}
}

func (g *Game) Update() error {
	// Update is run 60 times a second by default
	tps := float64(60)

	g.player.Grounded = false
	if g.gameMap.CollidesWith(g.player.SurfaceDetector()) {
		g.player.Grounded = true
	}

	g.player.Update(tps, g.tileSize, g.particleSystem)

	for _, row := range g.gameMap.Tiles {
		for _, t := range row {
			if !t.IsCollidable() {
				continue
			}

			g.player.Rigidbody.MoveOutOfBox(t.Hitbox)
		}
	}

	g.particleSystem.Update()

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
	g.particleSystem.Draw(screen, vector.Vector2{X: offsetX, Y: offsetY})
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %f", ebiten.ActualTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

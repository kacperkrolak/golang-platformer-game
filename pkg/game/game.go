/*
Package game contains the main game logic.

It heavily relies on the other packages. Mainly:
- pkg/world
- pkg/physics
- pkg/visuals
- ebitten
*/
package game

import (
	"fmt"
	"image/color"
	"kacperkrolak/golang-platformer-game/pkg/entity/manager"
	"kacperkrolak/golang-platformer-game/pkg/game/camera"
	"kacperkrolak/golang-platformer-game/pkg/game/player"
	"kacperkrolak/golang-platformer-game/pkg/physics/box"
	"kacperkrolak/golang-platformer-game/pkg/physics/rigidbody"
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"
	"kacperkrolak/golang-platformer-game/pkg/visuals/particle"
	"kacperkrolak/golang-platformer-game/pkg/world/gamemap"
	"kacperkrolak/golang-platformer-game/pkg/world/tilemap"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 520
)

// Game contains all the game logic.
type Game struct {
	gameMap        gamemap.Map
	background     tilemap.Map
	tilesImage     *ebiten.Image
	characterImage *ebiten.Image
	tileSize       int
	player         *player.Player
	camera         camera.Camera
	particleSystem *particle.ParticleSystem
	spawnPoint     vector.Vector2
	entityManager  *manager.Manager
	coinPositions  []vector.Vector2
}

const TileSize = 16

// Initialize the game.
func MakeGame(mapFile string, textureFile string, characterFile string) Game {
	tilesImage, err := loadTilesImage(textureFile)
	if err != nil {
		log.Fatal(err)
	}

	characterImage, err := loadTilesImage(characterFile)
	if err != nil {
		log.Fatal(err)
	}

	gameMap, tileMap, coinPositions, spawnPoint, err := loadWorldMap(mapFile, TileSize, tilesImage)
	if err != nil {
		log.Fatal(err)
	}

	particleSystem := particle.ParticleSystem{}
	player := player.Player{
		Rigidbody: rigidbody.Rigidbody{
			Hitbox: box.Box{
				Position: vector.Vector2{X: spawnPoint.X, Y: spawnPoint.Y},
				Size:     vector.Vector2{X: 14, Y: 21},
			},
		},
		Speed:          8.5 * TileSize,
		ParticleSystem: &particleSystem,
	}

	game := Game{
		gameMap:        gameMap,
		background:     tileMap,
		tilesImage:     tilesImage,
		characterImage: characterImage,
		tileSize:       TileSize,
		player:         &player,
		camera: camera.Camera{
			Position:   vector.Vector2{X: spawnPoint.X, Y: spawnPoint.Y},
			Velocity:   vector.Vector2{X: 0, Y: 0},
			Target:     &player,
			SmoothTime: 15,
		},
		spawnPoint:     spawnPoint,
		particleSystem: &particleSystem,
		coinPositions:  coinPositions,
		entityManager:  manager.NewManager(&particleSystem, tilesImage),
	}

	game.Reset()

	return game
}

func (g *Game) Reset() {
	g.player.Rigidbody.Hitbox.Position = g.spawnPoint
	g.player.Rigidbody.Velocity = vector.Vector2{X: 0, Y: 0}

	g.entityManager.Reset()
	for _, position := range g.coinPositions {
		g.entityManager.Spawn(position, manager.COIN)
	}
}

// Update the game and all its components.
//
// By default, Update is run 60 times a second.
func (g *Game) Update() error {
	tps := float64(60)
	deltaTime := time.Duration(1 / tps * float64(time.Second))

	if collidedWith := g.gameMap.CollidesWith(g.player.SurfaceDetector()); len(collidedWith) > 0 {
		groundedStatus := false
		for _, t := range collidedWith {
			if t.IsSolid() {
				groundedStatus = true
			}
			t.OnCollision(&g.player.Rigidbody)
		}
		g.player.UpdateGroundedState(groundedStatus)
	} else {
		g.player.UpdateGroundedState(false)
	}

	if collidedWith := g.gameMap.CollidesWith(g.player.WallDetector()); len(collidedWith) > 0 {
		g.player.UpdateWallSlidingState(true)
	} else {
		g.player.UpdateWallSlidingState(false)
	}

	g.player.Update(deltaTime, g.tileSize)

	if g.player.Rigidbody.Hitbox.Bottom() > float64(g.gameMap.Height()) {
		g.Reset()
		return nil
	}

	if collidedWith := g.gameMap.CollidesWith(g.player.Rigidbody.Hitbox); len(collidedWith) > 0 {
		for _, t := range collidedWith {
			if !t.IsCollidable() {
				continue
			}

			if t.IsDeadly() {
				// Make the game more fair by allowing player to touch 1 pixel of spikes
				displacementVector := g.player.Rigidbody.Hitbox.DisplacementVector(t.Hitbox())
				if displacementVector.Length() > 2 {
					g.Reset()
					return nil
				}
				continue
			}

			displacement := t.Hitbox().DisplacementVector(g.player.Rigidbody.Hitbox)
			g.player.OnBumping(displacement)
			g.player.Rigidbody.Hitbox.Position.Add(displacement)
		}
	}

	g.particleSystem.Update()
	g.entityManager.Update(deltaTime, &g.player.Rigidbody)

	err := g.camera.Update(tps)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// Calculate the offset of the camera.
func (g *Game) getScreenPosition(x, y float64) (float64, float64) {
	return ScreenWidth/2 - g.camera.Position.X, ScreenHeight/2 - g.camera.Position.Y
}

// Draw the game and all its components.
func (g *Game) Draw(screen *ebiten.Image) {
	// 6b978d
	background := color.RGBA{R: 107, G: 151, B: 141, A: 255}
	screen.Fill(background)

	offsetX, offsetY := g.getScreenPosition(0, 0)
	cameraOffset := vector.Vector2{X: offsetX, Y: offsetY}
	g.background.Draw(screen, cameraOffset, g.tilesImage, g.tileSize)
	g.gameMap.Draw(screen, cameraOffset, g.tilesImage, g.tileSize)

	g.entityManager.Draw(screen, cameraOffset)

	g.player.Draw(screen, offsetX, offsetY, g.characterImage, g.tileSize)
	g.particleSystem.Draw(screen, cameraOffset)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %f, FPS: %f", ebiten.ActualTPS(), ebiten.ActualFPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

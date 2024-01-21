package parser

import (
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"
	"kacperkrolak/golang-platformer-game/pkg/world/gamemap/block"
	"kacperkrolak/golang-platformer-game/pkg/world/tilemap/tile"
)

type ParsedData struct {
	Blocks        [][]block.Block
	Tiles         [][]tile.Tile
	CoinPositions []vector.Vector2
	SpawnPoint    vector.Vector2
}

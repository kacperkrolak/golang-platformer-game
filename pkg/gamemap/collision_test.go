package gamemap

import (
	"kacperkrolak/golang-platformer-game/pkg/gamemap/tile"
	"kacperkrolak/golang-platformer-game/pkg/gamemap/tile/mock"
	"kacperkrolak/golang-platformer-game/pkg/physics/box"
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"
	"testing"
)

func TestCollidesWith(t *testing.T) {
	testCases := []struct {
		name     string
		m        Map
		box      box.Box
		expected bool
	}{
		{
			name: "should return true when the box collides with collidable tiles",
			m: Map{
				Tiles: [][]tile.Tile{
					{mock.NewEmptyTile(), mock.NewCollidableTile(), mock.NewEmptyTile()},
					{mock.NewCollidableTile(), mock.NewCollidableTile(), mock.NewCollidableTile()},
					{mock.NewEmptyTile(), mock.NewCollidableTile(), mock.NewEmptyTile()},
				},
				TileSize: 10,
			},
			box: box.Box{
				Position: vector.Vector2{X: 15, Y: 15},
				Size:     vector.Vector2{X: 10, Y: 10},
			},
			expected: true,
		},
		{
			name: "should return false when the box does not collide with any collidable tiles",
			m: Map{
				Tiles: [][]tile.Tile{
					{mock.NewEmptyTile(), mock.NewEmptyTile(), mock.NewEmptyTile()},
					{mock.NewEmptyTile(), mock.NewEmptyTile(), mock.NewEmptyTile()},
					{mock.NewEmptyTile(), mock.NewEmptyTile(), mock.NewEmptyTile()},
				},
				TileSize: 10,
			},
			box: box.Box{
				Position: vector.Vector2{X: 15, Y: 15},
				Size:     vector.Vector2{X: 10, Y: 10},
			},
			expected: false,
		},
		{
			name: "should return true when the box overlaps multiple collidable tiles",
			m: Map{
				Tiles: [][]tile.Tile{
					{mock.NewCollidableTile(), mock.NewCollidableTile(), mock.NewCollidableTile()},
					{mock.NewCollidableTile(), mock.NewCollidableTile(), mock.NewCollidableTile()},
					{mock.NewCollidableTile(), mock.NewCollidableTile(), mock.NewCollidableTile()},
				},
				TileSize: 10,
			},
			box: box.Box{
				Position: vector.Vector2{X: 5, Y: 5},
				Size:     vector.Vector2{X: 20, Y: 20},
			},
			expected: true,
		},
		{
			name: "should return false when the box is outside of the map",
			m: Map{
				Tiles: [][]tile.Tile{
					{mock.NewEmptyTile(), mock.NewEmptyTile(), mock.NewEmptyTile()},
					{mock.NewEmptyTile(), mock.NewEmptyTile(), mock.NewEmptyTile()},
					{mock.NewEmptyTile(), mock.NewEmptyTile(), mock.NewEmptyTile()},
				},
				TileSize: 10,
			},
			box: box.Box{
				Position: vector.Vector2{X: 25, Y: 25},
				Size:     vector.Vector2{X: 10, Y: 10},
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.m.CollidesWith(tc.box)
			if (len(actual) > 0) != tc.expected {
				t.Errorf("CollidesWith() = %v, expected %v", actual, tc.expected)
			}
		})
	}
}

package gamemap

import (
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
				Tiles: [][]Tile{
					{{Type: EMPTY}, {Type: DIRT}, {Type: EMPTY}},
					{{Type: DIRT}, {Type: DIRT}, {Type: DIRT}},
					{{Type: EMPTY}, {Type: DIRT}, {Type: EMPTY}},
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
				Tiles: [][]Tile{
					{{Type: EMPTY}, {Type: EMPTY}, {Type: EMPTY}},
					{{Type: EMPTY}, {Type: EMPTY}, {Type: EMPTY}},
					{{Type: EMPTY}, {Type: EMPTY}, {Type: EMPTY}},
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
				Tiles: [][]Tile{
					{{Type: DIRT}, {Type: DIRT}, {Type: DIRT}},
					{{Type: DIRT}, {Type: DIRT}, {Type: DIRT}},
					{{Type: DIRT}, {Type: DIRT}, {Type: DIRT}},
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
				Tiles: [][]Tile{
					{{Type: EMPTY}, {Type: EMPTY}, {Type: EMPTY}},
					{{Type: EMPTY}, {Type: EMPTY}, {Type: EMPTY}},
					{{Type: EMPTY}, {Type: EMPTY}, {Type: EMPTY}},
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
			if actual != tc.expected {
				t.Errorf("CollidesWith() = %v, expected %v", actual, tc.expected)
			}
		})
	}
}

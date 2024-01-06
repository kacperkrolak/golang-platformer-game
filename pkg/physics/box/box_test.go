package box

import (
	"kacperkrolak/golang-platformer-game/pkg/physics/vector"
	"testing"
)

func TestBox_Left(t *testing.T) {
	box := Box{
		Position: vector.Vector2{X: 10, Y: 20},
		Size:     vector.Vector2{X: 1, Y: 2},
	}

	if box.Left() != 10 {
		t.Errorf("expected 10, got %v", box.Left())
	}
}

func TestBox_Right(t *testing.T) {
	box := Box{
		Position: vector.Vector2{X: 10, Y: 20},
		Size:     vector.Vector2{X: 1, Y: 2},
	}

	if box.Right() != 11 {
		t.Errorf("expected 11, got %v", box.Right())
	}
}

func TestBox_Top(t *testing.T) {
	box := Box{
		Position: vector.Vector2{X: 10, Y: 20},
		Size:     vector.Vector2{X: 1, Y: 2},
	}

	if box.Top() != 20 {
		t.Errorf("expected 20, got %v", box.Top())
	}
}

func TestBox_Bottom(t *testing.T) {
	box := Box{
		Position: vector.Vector2{X: 10, Y: 20},
		Size:     vector.Vector2{X: 1, Y: 2},
	}

	if box.Bottom() != 22 {
		t.Errorf("expected 22, got %v", box.Bottom())
	}
}

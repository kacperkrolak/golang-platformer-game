package vector

type Vector2 struct {
	X float64
	Y float64
}

func (v *Vector2) Add(v2 Vector2) Vector2 {
	v.X += v2.X
	v.Y += v2.Y
	return *v
}

func (v *Vector2) Scale(s float64) Vector2 {
	v.X *= s
	v.Y *= s
	return *v
}

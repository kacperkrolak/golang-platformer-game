package vector

import "math"

type Vector2 struct {
	X float64
	Y float64
}

func (v *Vector2) Add(v2 Vector2) *Vector2 {
	v.X += v2.X
	v.Y += v2.Y
	return v
}

func (v Vector2) Added(v2 Vector2) Vector2 {
	return *v.Add(v2)
}

func (v *Vector2) Scale(s float64) *Vector2 {
	v.X *= s
	v.Y *= s
	return v
}

func (v Vector2) Scaled(s float64) Vector2 {
	return *v.Scale(s)
}

func (v *Vector2) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vector2) Normalize() *Vector2 {
	length := v.Length()
	v.X /= length
	v.Y /= length
	return v
}

func (v Vector2) Normalized() Vector2 {
	return *v.Normalize()
}

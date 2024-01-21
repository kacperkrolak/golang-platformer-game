/*
Package vector provides a simple 2D vector representation and operations on it.

Although most of the operations are universal, the package is designed to be used
with a coordinate system where Y axis points down and X axis points right.
For example, method `Down()` return a vector with Y=1.
*/
package vector

import "math"

// Vector2 is a simple 2D vector representation.
type Vector2 struct {
	X float64
	Y float64
}

// Vector with both components set to 0.
func Zero() Vector2 {
	return Vector2{X: 0, Y: 0}
}

// Vector pointing right.
func Right() Vector2 {
	return Vector2{X: 1, Y: 0}
}

// Vector pointing left.
func Left() Vector2 {
	return Vector2{X: -1, Y: 0}
}

// Vector pointing up.
func Up() Vector2 {
	return Vector2{X: 0, Y: -1}
}

// Vector pointing down.
func Down() Vector2 {
	return Vector2{X: 0, Y: 1}
}

// Add second vector to the first one.
func (v *Vector2) Add(v2 Vector2) *Vector2 {
	v.X += v2.X
	v.Y += v2.Y
	return v
}

// Return a new vector which is a sum of two vectors.
func (v Vector2) Added(v2 Vector2) Vector2 {
	return *v.Add(v2)
}

// Scale the vector by a given factor.
func (v *Vector2) Scale(s float64) *Vector2 {
	v.X *= s
	v.Y *= s
	return v
}

// Return a new vector which is a scaled version of the original one.
func (v Vector2) Scaled(s float64) Vector2 {
	return *v.Scale(s)
}

// Return the length of the vector.
func (v *Vector2) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Changes the vector so that its length is 1, but the direction is the same.
func (v *Vector2) Normalize() *Vector2 {
	length := v.Length()
	if length == 0 {
		return v
	}

	v.X /= length
	v.Y /= length
	return v
}

// Return a new vector with the same direction, but length of 1.
func (v Vector2) Normalized() Vector2 {
	return *v.Normalize()
}

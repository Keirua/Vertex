package main

import "math"

type Vertex struct {
	X, Y, Z float64
}

func (v *Vertex) Scale(s float64) *Vertex {
	v.X = v.X * s
	v.Y = v.Y * s
	v.Z = v.Z * s

	return v
}

func (v Vertex) Mulf(s float64) Vertex {
	return Vertex{v.X * s, v.Y * s, v.Z * s}
}

func (v *Vertex) Normalize() *Vertex {
	var l = v.Length()
	if l == 0.0 {
		panic("C'est foutu !")
	}
	var invLength = 1.0 / l
	v = v.Scale(invLength)
	return v
}

func (v Vertex) LengthSq() float64 {
	return v.Dot(v)
}

func (v Vertex) Length() float64 {
	return math.Sqrt(v.LengthSq())
}

func (v Vertex) Dot(v2 Vertex) float64 {
	return (v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z)
}

func (v Vertex) Add(v2 Vertex) Vertex {
	return Vertex{v.X + v2.X, v.Y + v2.Y, v.Z + v2.Z}
}

func (v Vertex) Substract(v2 Vertex) Vertex {
	return Vertex{v.X - v2.X, v.Y - v2.Y, v.Z - v2.Z}
}

package main

type Ray struct {
	Origin    Vertex
	Direction Vertex
}

func (r Ray) VertexAt(t float64) Vertex {
	return Vertex{
		(r.Origin.X + r.Direction.X*t),
		(r.Origin.Y + r.Direction.Y*t),
		(r.Origin.Z + r.Direction.Z*t)}
}

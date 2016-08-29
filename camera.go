package main

import "math"

type Camera struct {
	width       int
	height      int
	fov         float64
	angle       float64
	aspectRatio float64
}

func (camera *Camera) Initialize(width int, height int, fov float64) {
	camera.width = width
	camera.height = height
	camera.fov = fov
	camera.aspectRatio = float64(width) / float64(height)
	camera.angle = math.Tan(0.5 * camera.fov)
}

func (camera Camera) ComputeRayDirection(x float64, y float64) Ray {
	var xx float64 = float64(2*((float64(x)+0.5)/float64(camera.width))-1) * camera.angle * camera.aspectRatio
	var yy float64 = float64(1.0-2.0*((float64(y)+0.5)/float64(camera.height))) * camera.angle

	var direction = Vertex{xx, yy, -1}
	direction.Normalize()

	return Ray{Vertex{}, direction}
}

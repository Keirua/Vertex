package main

import (
	"math/rand"
)

type Light struct {
	Position Vertex
	Color    Color01
}


func (currLight Light) generateLightRay (intersectionPoint Vertex, softShadowStrength float64) Ray {
	var lightRay Ray

	lightRay.Origin = intersectionPoint
	lightRay.Origin.X += (rand.Float64()*2.0 - 1.0) * softShadowStrength
	lightRay.Origin.Y += (rand.Float64()*2.0 - 1.0) * softShadowStrength
	lightRay.Origin.Z += (rand.Float64()*2.0 - 1.0) * softShadowStrength
	
	lightRay.Direction = currLight.Position.Substract(intersectionPoint)
	lightRay.Direction.Normalize()

	return lightRay
}
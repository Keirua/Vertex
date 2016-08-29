package main

import "math"

type Sphere struct {
	Center   Vertex
	Radius   float64
	Material Material
}

func (sphere Sphere) GetMaterial() Material {
    return sphere.Material;
}

func (sphere Sphere) ComputeNormalAtPoint(pointHit Vertex) Vertex {
    return pointHit.Substract(sphere.Center)
}


func (sphere Sphere) Intersect(ray Ray, info *IntersectionInfo) bool {
	var l Vertex = sphere.Center.Substract(ray.Origin)
	var tca = l.Dot(ray.Direction)
	if tca < 0 {
		return false
	}

	var d2 = l.Dot(l) - (tca * tca)
	if d2 > (sphere.Radius * sphere.Radius) {
		return false
	}

	var thc float64 = math.Sqrt(sphere.Radius*sphere.Radius - d2)

	info.T = math.Min(tca-thc, tca+thc)

	return true
}

/*
    The normal must be normalized
*/
func (sphere Sphere) ComputeUV(normal Vertex) (float64, float64) {
    /*
    u = 0.5 + arctan2(dz, dx) / (2*pi)
    v = 0.5 - arcsin(dy) / pi
    */
    var u = 0.5 + (math.Atan2(normal.Z, normal.X) / (2*math.Pi))
    var v = 0.5 - (math.Asin(normal.Y) / math.Pi)

    return u, v;
}

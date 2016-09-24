package main

import "math"

type Cylinder struct {
	Center   Vertex
	Radius   float64
	Material *Material
}

func (cylinder Cylinder) GetMaterial() *Material {
	return cylinder.Material
}

func (cylinder Cylinder) ComputeNormalAtIntersectionPoint(info *IntersectionInfo) Vertex {
	var normal = info.IntersectionPoint.Substract(cylinder.Center)
	normal.Z = 0;
	return normal
}

func (cylinder Cylinder) Intersect(ray Ray, info *IntersectionInfo) bool {
	var origin = ray.Origin.Substract(cylinder.Center)

	var a float64 = ray.Direction.X*ray.Direction.X + ray.Direction.Y*ray.Direction.Y
	var b float64 = 2*origin.X*ray.Direction.X+2*origin.Y*ray.Direction.Y
	var c float64 = origin.X*origin.X + origin.Y*origin.Y - 1.0

	if (a != 0) {
		var sqrtDelta = math.Sqrt(b*b - 4.0*a*c)
		var t1 float64 = -(b - sqrtDelta)/(2*a)
		var t2 float64 = -(b + sqrtDelta)/(2*a)

		//info.T = math.Min (t1, t2)
		info.T = math.Min (math.Max(0, t1), math.Max(0, t2))

		return info.T > 0
	}
	return false
}

/*
   The normal must be normalized
*/
func (cylinder Cylinder) ComputeUV(info IntersectionInfo) (float64, float64) {
	// Cylindrical texture mapping
	var u = math.Acos(info.IntersectionPoint.X/cylinder.Radius) / (2*math.Pi)
	var v = math.Acos(info.IntersectionPoint.Y/cylinder.Radius) / (2*math.Pi)

	return u, v
}

package main

type Hittable interface {
	ComputeNormalAtIntersectionPoint(info *IntersectionInfo) Vertex
	Intersect(ray Ray, info *IntersectionInfo) bool
	GetMaterial() *Material
	ComputeUV(info IntersectionInfo) (float64, float64)
}

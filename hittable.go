package main

type Hittable interface {
    ComputeNormalAtIntersectionPoint (info *IntersectionInfo) Vertex
    Intersect (ray Ray, info* IntersectionInfo) bool
    GetMaterial() *Material;
    ComputeUV(point Vertex) (float64, float64);
}
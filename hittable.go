package main

type Hittable interface {
    ComputeNormalAtPoint (pointHit Vertex) Vertex
    Intersect (ray Ray, info* IntersectionInfo) bool
    GetMaterial() Material;
    ComputeUV(point Vertex) (float64, float64);
}
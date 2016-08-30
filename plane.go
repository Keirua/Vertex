package main

import "math"

type Plane struct {
    Point   Vertex          // A point on the place
    Normal  Vertex
    Material *Material
}

func (plane Plane) GetMaterial() *Material {
    return plane.Material;
}

func (plane Plane) ComputeNormalAtIntersectionPoint (info *IntersectionInfo) Vertex {
    /*if (plane.Normal.Dot(info.Ray.Direction) < 0.0) {
        return plane.Normal.Mulf(-1)
    }*/
    return plane.Normal;
}


func (plane Plane) Intersect(ray Ray, info *IntersectionInfo) bool {
    var normalAndDirectionDotProduct = plane.Normal.Dot(ray.Direction)
    if (normalAndDirectionDotProduct <= 0.0) {
        return false;
    }

    var D = -plane.Normal.Dot(plane.Point)
    var rayO = ray.Origin.Substract (plane.Point);

    info.T = -((plane.Normal.Dot(rayO) + D) / (normalAndDirectionDotProduct))

    return true;
}

func (plane Plane) ComputeUV(normal Vertex) (float64, float64) {
    var u = 0.5 + (math.Atan2(normal.Z, normal.X) / (2*math.Pi))
    var v = 0.5 - (math.Asin(normal.Y) / math.Pi)

    return u, v;
}

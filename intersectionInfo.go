package main

import "math"

type IntersectionInfo struct {
	T                 float64
	ObjectHit         *Hittable
	Ray               *Ray
	IntersectionPoint Vertex
	Normal            Vertex
}

/*
   Finds, among all the objects in the scene, with which one there is the closest intersection (if any)
*/
func (intersectionInfo *IntersectionInfo) GetIntersectionInfo(ray Ray, scene *[]Hittable) {
	intersectionInfo.T = math.MaxFloat64

	// Finds if there is an object along the ray
	for index, object := range *scene {
		var currentIntersectionInfo IntersectionInfo
		if object.Intersect(ray, &currentIntersectionInfo) {
			if currentIntersectionInfo.T < intersectionInfo.T {
				intersectionInfo.T = currentIntersectionInfo.T
				intersectionInfo.ObjectHit = &(*scene)[index]
			}
		}
	}

	// If so, computes intersection point and normal
	if intersectionInfo.ObjectHit != nil {
		intersectionInfo.IntersectionPoint = ray.VertexAt(intersectionInfo.T)
		intersectionInfo.Ray = &ray
		var normal = (*intersectionInfo.ObjectHit).ComputeNormalAtIntersectionPoint(intersectionInfo)
		normal.Normalize()
		intersectionInfo.Normal = normal

	}
}

package main

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"
	"math3d"
)

const (
	width  = 640
	height = 480
	fov    = 30.0
)

var whiteMaterial = math3d.Material{math3d.Color01{0.8, 0.8, 0.8}}
var redMaterial = math3d.Material{math3d.Color01{1, 0, 0}}
var blueMaterial = math3d.Material{math3d.Color01{0, .5, 1}}
var purpleMaterial = math3d.Material{math3d.Color01{0.65, .2, 0.97}}
var greenMaterial = math3d.Material{math3d.Color01{0.3, 0.9, 0.2}}

var sphereFloor = math3d.Sphere{math3d.Vertex{0, 10003, -20}, 10000.0, whiteMaterial}
var sphere1 = math3d.Sphere{math3d.Vertex{4.0, -1, -5}, 2.0, redMaterial}
var sphere2 = math3d.Sphere{math3d.Vertex{0.0, 0, -15}, 4, greenMaterial}
var sphere3 = math3d.Sphere{math3d.Vertex{3.0, 2, -10}, 3, blueMaterial}
var sphere4 = math3d.Sphere{math3d.Vertex{-5.5, 0, -8}, 3, purpleMaterial}

//var lightSphere = math3d.Sphere{math3d.Vertex{3.0, -3, -10}, 0.5, purpleMaterial}
var light = math3d.Light{math3d.Vertex{3.0, -3, -10}, math3d.Color01{0.65, .2, 0.97}}

var g_Spheres = []math3d.Sphere{sphere2, sphere1, sphereFloor, sphere3, sphere4/*, lightSphere*/}
var g_Lights = []math3d.Light{light}
var g_Camera math3d.Camera

/*
	Finds, among all the objects in the scene, with which one there is the closest intersection (if any)
*/
func getIntersectionInfo(ray math3d.Ray) math3d.IntersectionInfo {
	var intersectionInfo = math3d.IntersectionInfo{math.MaxFloat64, -1}
	for index, sph := range g_Spheres {
		var currentIntersectionInfo math3d.IntersectionInfo
		if sph.Intersect(ray, &currentIntersectionInfo) {
			if currentIntersectionInfo.T < intersectionInfo.T {
				intersectionInfo.T = currentIntersectionInfo.T
				intersectionInfo.ObjectIndex = index
			}
		}
	}

	return intersectionInfo;
}

func trace(ray math3d.Ray) math3d.Color01 {
	var finalColor math3d.Color01

	var intersectionInfo = getIntersectionInfo(ray)

	if intersectionInfo.ObjectIndex != -1 {
		var objectHit = g_Spheres[intersectionInfo.ObjectIndex]

		/*var intersectionPoint = ray.VertexAt(intersectionInfo.T)
		var normal = objectHit.ComputeNormalAtPoint(intersectionPoint)
		normal.Normalize()*/

		finalColor = objectHit.Material.SurfaceColor
	}

	return finalColor
}

func computeColorAtXY(x int, y int) color.RGBA {
	var finalColor math3d.Color01

	// 2x2 anti aliasing : for every point, we send 4 ray
	// each one contributing to 1/4th of the final pixel color
	// much slower since we launch 4x more rays per pixels
	/*for i := float64(x); i < float64(x)+1.0; i += 0.5 {
		for j := float64(y); j < float64(y)+1.0; j += 0.5 {
			var ray = g_Camera.ComputeRayDirection(i, j)
			var tracedColor = trace(ray)

			finalColor = finalColor.Add(tracedColor.Mul(0.25))
		}
	}*/

	var ray = g_Camera.ComputeRayDirection(float64(x), float64(y))
	finalColor = trace(ray)

	return finalColor.ToRGBA()
}

func main() {
	rand.Seed(42)
	g_Camera.Initialize(width, height, fov)

	image := generateImage(width, height, computeColorAtXY)
	SavePNG(image, "out.png")
	fmt.Println("Success !")
}

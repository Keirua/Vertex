package main

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"
	"math3d"
)

const (
	width             = 640
	height            = 480
	fov               = 30.0
	antiAliasingLevel = 1 // minimum 1
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
var light2 = math3d.Light{math3d.Vertex{0, -5, 0}, math3d.Color01{0.87, 0.33, 0.97}}

var g_Spheres = []math3d.Sphere{sphere2, sphere1, sphereFloor, sphere3, sphere4 /*, lightSphere*/}
var g_Lights = []math3d.Light{ /*light, */ light2}
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

	return intersectionInfo
}

func trace(ray math3d.Ray) math3d.Color01 {
	var finalColor math3d.Color01
	var coef float64 = 1.0
	var intersectionInfo = getIntersectionInfo(ray)

	if intersectionInfo.ObjectIndex != -1 {
		var objectHit = g_Spheres[intersectionInfo.ObjectIndex]
		var intersectionPoint = ray.VertexAt(intersectionInfo.T)
		var normal = objectHit.ComputeNormalAtPoint(intersectionPoint)
		normal.Normalize()

		for _, currLight := range g_Lights {
			var lightRay math3d.Ray
			lightRay.Origin = intersectionPoint
			lightRay.Direction = currLight.Position.Substract(intersectionPoint)
			lightRay.Direction.Normalize()
			if lightRay.Direction.Dot(normal) <= 0.0 {
				continue
			}

			var isInShadow bool = false
			for _, currObject := range g_Spheres {
				var shadowIntersectionInfo math3d.IntersectionInfo
				if currObject.Intersect(lightRay, &shadowIntersectionInfo) {
					isInShadow = true
					break
				}
			}

			if !isInShadow {
				// lambert contribution
				var lambert float64 = lightRay.Direction.Dot(normal) * coef

				// finalColor = finalColor + lambert * currentLight * currentMaterial
				finalColor = finalColor.AddColor(objectHit.Material.SurfaceColor.MulColor(currLight.Color).MulFloat(lambert))
			}
		}
		//finalColor = objectHit.Material.SurfaceColor
	}

	return finalColor
}

func computeColorAtXY(x int, y int) color.RGBA {
	var finalColor math3d.Color01

	var steps float64 = 1.0 / antiAliasingLevel

	// with 2x2 anti aliasing, for every point, we send 4 ray
	// each one contributing to 1/4th of the final pixel color
	// much slower since we launch 4x more rays per pixels
	var i float64
	var j float64

	for i = 0.0; i < antiAliasingLevel; i++ {
		for j = 0.0; j < antiAliasingLevel; j++ {
			var ray = g_Camera.ComputeRayDirection(float64(x)+i*steps, float64(y)+j*steps)
			var tracedColor = trace(ray)

			finalColor = finalColor.AddColor(tracedColor.MulFloat(steps * steps))
		}
	}


	return finalColor.ToRGBA()
}

func main() {
	rand.Seed(42)
	g_Camera.Initialize(width, height, fov)

	image := generateImage(width, height, computeColorAtXY)
	SavePNG(image, "out.png")
	fmt.Println("Success !")
}

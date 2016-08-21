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

var light = math3d.Sphere{math3d.Vertex{3.0, -3, -10}, 0.5, purpleMaterial}

var g_Spheres = []math3d.Sphere{sphere2, sphere1, sphereFloor, sphere3, sphere4, light}
var g_Camera math3d.Camera

func trace(ray math3d.Ray) math3d.Color01 {
	var intersectionInfo = math3d.IntersectionInfo{math.MaxFloat64, math3d.Material{}}
	var finalColor math3d.Color01

	for _, sph := range g_Spheres {
		var currentIntersectionInfo math3d.IntersectionInfo
		if sph.Intersect(ray, &currentIntersectionInfo) {
			if currentIntersectionInfo.T < intersectionInfo.T {
				intersectionInfo.T = currentIntersectionInfo.T
				intersectionInfo.Material = sph.Material
			}
		}
	}

	finalColor = intersectionInfo.Material.SurfaceColor

	return finalColor
}

func computeColorAtXY(x int, y int) color.RGBA {
	var ray = g_Camera.ComputeRayDirection(x, y)

	var tracedColor = trace(ray)

	return tracedColor.ToRGBA()
}

func main() {
	rand.Seed(42)
	g_Camera.Initialize(width, height, fov)

	image := generateImage(width, height, computeColorAtXY)
	SavePNG(image, "out.png")
	fmt.Println("Success !")
}

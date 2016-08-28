package main

import (
	"fmt"
	"image/color"
	"math"
	"math3d"
)

const (
	DEFAULT_WIDTH              = 640
	DEFAULT_HEIGHT             = 480
	DEFAULT_FOV                = 30.0
	DEFAULT_ANTIALIASING_LEVEL = 1 // minimum
	MAX_DEPTH                  = 2
)

var checkboardTexture16 = math3d.CheckboardTexture{16,16}
var imageTexture = math3d.FileTexture{"images/cropped-P1120606-small.jpg",nil}

var whiteMaterial = math3d.Material{math3d.Color01{0.8, 0.8, 0.8}, 0.5, nil}
var redMaterial = math3d.Material{math3d.Color01{1, 0, 0}, 0.5, nil}
var blueMaterial = math3d.Material{math3d.Color01{0, .5, 1}, 0.5, nil}
var purpleMaterial = math3d.Material{math3d.Color01{0.65, .2, 0.97}, 0.5, &imageTexture}
var greenMaterial = math3d.Material{math3d.Color01{0.3, 0.9, 0.2}, 0.5, &checkboardTexture16}

var sphereFloor = math3d.Sphere{math3d.Vertex{0, 10003, -20}, 10000.0, whiteMaterial}
var sphere1 = math3d.Sphere{math3d.Vertex{4.0, -1, -5}, 2.0, redMaterial}
var sphere2 = math3d.Sphere{math3d.Vertex{0.0, 0, -15}, 4, greenMaterial}
var sphere3 = math3d.Sphere{math3d.Vertex{3.0, 2, -10}, 3, blueMaterial}
var sphere4 = math3d.Sphere{math3d.Vertex{-5.5, 0, -8}, 3, purpleMaterial}

var light = math3d.Light{math3d.Vertex{3.0, -10, -10}, math3d.Color01{0.65, .6, 0.97}}
var light2 = math3d.Light{math3d.Vertex{0, -5, 0}, math3d.Color01{0.87, 0.8, 0.97}}

var g_VisibleObjects []math3d.Hittable;

var g_Lights = []math3d.Light{light, light2}
var g_Camera math3d.Camera
var g_Options Options


func trace(ray math3d.Ray, contributionCoef float64, depth int) math3d.Color01 {
	var finalColor math3d.Color01
	// Find the closest object the ray intersects																														
	var intersectionInfo math3d.IntersectionInfo
	intersectionInfo.GetIntersectionInfo(ray, g_VisibleObjects);

	if intersectionInfo.ObjectHit != nil {
		var objectHit = *intersectionInfo.ObjectHit
		var normal = intersectionInfo.Normal
		var intersectionPoint = intersectionInfo.IntersectionPoint

		// Compute color at the surface
		var colorOnSurface = objectHit.GetMaterial().SurfaceColor
		// If there is a texture, we add it's contribution
		if (objectHit.GetMaterial().Texture != nil) {
			var u, v = objectHit.ComputeUV(normal)
			var colorAtUV = objectHit.GetMaterial().Texture.GetColor01AtUV(u, v)
			colorOnSurface = colorOnSurface.MulColor(colorAtUV)
		}

		// Add Reflection		
		var reflectionRefractionColorMix math3d.Color01;
		var reflectionContributionCoef = contributionCoef * objectHit.GetMaterial().ReflectionCoef;
		
		// Computes the reflection ray
		var reflet float64 = 2.0 * (ray.Direction.Dot(normal));
		var reflectedRay math3d.Ray;
		reflectedRay.Origin = intersectionPoint.Add(normal.Mulf(1e-4))
		reflectedRay.Direction = ray.Direction.Substract(normal.Mulf(reflet))

		if ((objectHit.GetMaterial().ReflectionCoef >0 ) && depth < g_Options.MaxDepth){

			var reflectionColor = trace(reflectedRay, reflectionContributionCoef, depth+1)

			reflectionRefractionColorMix = reflectionColor;
		}

		// Add Lighting		
		for _, currLight := range g_Lights {
			var lightRay math3d.Ray
			lightRay.Origin = intersectionPoint
			lightRay.Direction = currLight.Position.Substract(intersectionPoint)
			lightRay.Direction.Normalize()
			if lightRay.Direction.Dot(normal) <= 0.0 {
				continue
			}
			// Throw shadow rays to see if objects are blocking the light
			var isInShadow bool = false
			for _, currObject := range g_VisibleObjects {
				var shadowIntersectionInfo math3d.IntersectionInfo
				if currObject.Intersect(lightRay, &shadowIntersectionInfo) {
					isInShadow = true
					break
				}
			}

			if !isInShadow {
				// lambert contribution
				var lambert float64 = lightRay.Direction.Dot(normal) * contributionCoef

				// finalColor = finalColor + lambert * currentLight * currentMaterial
				finalColor = finalColor.AddColor(colorOnSurface.MulColor(currLight.Color).MulFloat(lambert))
			} else {
				// soften the shadow. Total hack, no solid mathematical foundation
				finalColor = finalColor.AddColor(colorOnSurface.MulFloat(0.1))
			}
		}
		//finalColor = colorOnSurface
		finalColor = finalColor.AddColor(reflectionRefractionColorMix.MulFloat(reflectionContributionCoef))
	}


	return finalColor
}

func clampColor(c math3d.Color01) math3d.Color01 {
	return math3d.Color01{
		math.Min(1.0, math.Max(0.0, c.R)),
		math.Min(1.0, math.Max(0.0, c.G)),
		math.Min(1.0, math.Max(0.0, c.B))}
}

func computeColorAtXY(x int, y int) color.RGBA {
	var finalColor math3d.Color01

	var steps float64 = 1.0 / float64(g_Options.AntiAliasingLevel)
	var rayContributionCoefficient float64 = 1.0 / float64(g_Options.AntiAliasingLevel*g_Options.AntiAliasingLevel);

	// with 2x2 anti aliasing, for every point, we send 4 ray
	// each one contributing to 1/4th of the final pixel color
	// much slower since we launch 4x more rays per pixels
	var i float64
	var j float64

	for i = 0.0; i < float64(g_Options.AntiAliasingLevel); i++ {
		for j = 0.0; j < float64(g_Options.AntiAliasingLevel); j++ {
			var ray = g_Camera.ComputeRayDirection(float64(x)+i*steps, float64(y)+j*steps)
			var tracedColor = trace(ray, 1.0, 0)

			finalColor = finalColor.AddColor(tracedColor.MulFloat(rayContributionCoefficient))
		}
	}

	return clampColor(finalColor).ToRGBA()
}


func init() {
	g_Options.ParseCommandLineOptions();

	g_Camera.Initialize(g_Options.Width, g_Options.Height, g_Options.Fov)

	g_VisibleObjects = append(g_VisibleObjects, sphere2, sphere1, sphereFloor, sphere3, sphere4 /*, lightSphere*/)
	imageTexture.Load()
}

func main() {
	image := generateImage(g_Options.Width, g_Options.Height, computeColorAtXY)
	fmt.Println(g_Options)
	SavePNG(image, g_Options.OutputFilename)

	fmt.Println("Success !")
}

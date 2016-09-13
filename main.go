package main

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"os"
	"runtime/pprof"
)

const (
	DEFAULT_WIDTH              = 640
	DEFAULT_HEIGHT             = 480
	DEFAULT_FOV                = 30.0
	DEFAULT_ANTIALIASING_LEVEL = 1 // minimum
	MAX_DEPTH                  = 2
)

var checkboardTexture16 = CheckboardTexture{16, 16}
var imageTexture = FileTexture{"images/cropped-P1120606-small.jpg", nil}

var turbulence = NewTurbulence(32)
var marble5 = NewMarble(5, 10, 5, 32)
var marble1 = NewMarble(5, 10, 1, 32)
var wood = NewWood(12, 0.1, 32)

var whiteMaterial = Material{Color01{0.8, 0.8, 0.8}, 0.5, nil, Color01{0.8, 0.8, 0.8}, 60}
var redMaterial = Material{Color01{1, 0, 0}, 0.5, nil, Color01{0.8, 0.8, 0.8}, 50}
var blueMaterial = Material{Color01{0, .5, 1}, 0.5, marble5, Color01{0.3, 0.3, 0.3}, 60}
var purpleMaterial = Material{Color01{0.65, .2, 0.97}, 0.5, &checkboardTexture16, Color01{0.8, 0.8, 0.8}, 60}
var greenMaterial = Material{Color01{0.3, 0.9, 0.2}, 0.5, turbulence, Color01{0.8, 0.8, 0.8}, 60}

var sphereFloor = Sphere{Vertex{0, 10003, -20}, 10000.0, &whiteMaterial}
var planeFloor = Plane{Vertex{0, 1, 0}, Vertex{0, -1, 0}, &whiteMaterial}
var sphere1 = Sphere{Vertex{4.0, -1, -5}, 2.0, &redMaterial}
var sphere2 = Sphere{Vertex{0.0, 0, -15}, 4, &greenMaterial}
var sphere3 = Sphere{Vertex{3.0, 2, -10}, 3, &blueMaterial}
var sphere4 = Sphere{Vertex{-5.5, 0, -8}, 3, &purpleMaterial}

var light = Light{Vertex{3.0, -10, -10}, Color01{0.65, .6, 0.97}}
var light2 = Light{Vertex{0, -5, 0}, Color01{0.87, 0.8, 0.97}}

var g_VisibleObjects []Hittable

var g_Lights = []Light{light, light2}
var g_Camera Camera
var g_Options Options

func trace(ray Ray, contributionCoef float64, depth int) Color01 {
	//fmt.Println(ray)
	var finalColor Color01
	// Find the closest object the ray intersects
	var intersectionInfo IntersectionInfo
	intersectionInfo.GetIntersectionInfo(ray, &g_VisibleObjects)

	if intersectionInfo.ObjectHit != nil {
		var objectHit = *intersectionInfo.ObjectHit
		var normal = intersectionInfo.Normal
		var intersectionPoint = intersectionInfo.IntersectionPoint

		// Compute color at the surface
		var colorOnSurface = objectHit.GetMaterial().SurfaceColor

		// If there is a texture, we add it's contribution
		if objectHit.GetMaterial().Texture != nil {
			var u, v = objectHit.ComputeUV(normal)
			var colorAtUV = objectHit.GetMaterial().Texture.GetColor01AtUV(u, v)
			colorOnSurface = colorOnSurface.MulColor(colorAtUV)
		}

		// Add Reflection
		var reflectionRefractionColorMix Color01
		var reflectionContributionCoef = contributionCoef * objectHit.GetMaterial().ReflectionCoef

		// Computes the reflection ray
		var reflet float64 = 2.0 * (ray.Direction.Dot(normal))
		var reflectedRay Ray
		reflectedRay.Origin = intersectionPoint.Add(normal.Mulf(1e-4))
		reflectedRay.Direction = ray.Direction.Substract(normal.Mulf(reflet))

		if (objectHit.GetMaterial().ReflectionCoef > 0) && depth < g_Options.MaxDepth {

			var reflectionColor = trace(reflectedRay, reflectionContributionCoef, depth+1)

			reflectionRefractionColorMix = reflectionColor
		}

		// Add Lighting
		for _, currLight := range g_Lights {
			var lightRay Ray
			lightRay.Origin = intersectionPoint
			lightRay.Direction = currLight.Position.Substract(intersectionPoint)
			lightRay.Direction.Normalize()
			if lightRay.Direction.Dot(normal) <= 0.0 {
				continue
			}
			// Throw shadow rays to see if objects are blocking the light
			var isInShadow bool = false
			for _, currObject := range g_VisibleObjects {
				var shadowIntersectionInfo IntersectionInfo
				if currObject.Intersect(lightRay, &shadowIntersectionInfo) {
					isInShadow = true
					break
				}
			}

			if !isInShadow {
				// blinn-phong contribution (for the specular highlights)
				var blinnDirection = lightRay.Direction.Substract(ray.Direction)
				blinnDirection.Normalize()

				var blinnCoef = math.Max(0.0, blinnDirection.Dot(normal))
				var blinn = objectHit.GetMaterial().SpecularColor.MulFloat(math.Pow(blinnCoef, objectHit.GetMaterial().SpecularPower) * contributionCoef)
				finalColor = finalColor.AddColor(blinn.MulColor(currLight.Color))

				// lambert contribution
				var lambert float64 = lightRay.Direction.Dot(normal) * contributionCoef

				// finalColor = finalColor + lambert * currentLight * currentMaterial
				finalColor = finalColor.AddColor(colorOnSurface.MulColor(currLight.Color).MulFloat(lambert))

			} else {
				// soften the shadow. Total hack, no solid mathematical foundation
				finalColor = finalColor.AddColor(colorOnSurface.MulFloat(0.1))
			}
		}

		finalColor = finalColor.AddColor(reflectionRefractionColorMix.MulFloat(reflectionContributionCoef))
	}

	return finalColor
}

func clampMinMax(v float64) float64{
	return math.Min(1.0, math.Max(0.0, v))
}

func clampExp(v float64) float64{
	var coef float64 = -2.0
	return 1-math.Exp(v*coef)
}

func clampColor(c Color01) Color01 {
	return Color01{
		clampExp(c.R),
		clampExp(c.G),
		clampExp(c.B)}
}

func computeColorAtXY(x int, y int) color.RGBA {
	var finalColor Color01

	var steps float64 = 1.0 / float64(g_Options.AntiAliasingLevel)
	var rayContributionCoefficient float64 = 1.0 / float64(g_Options.AntiAliasingLevel*g_Options.AntiAliasingLevel)

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
	g_Options.ParseCommandLineOptions()

	g_Camera.Initialize(g_Options.Width, g_Options.Height, g_Options.Fov)

	g_VisibleObjects = append(g_VisibleObjects, sphere2, sphere1, sphereFloor, sphere3, sphere4 /*, lightSphere*/)
	//g_VisibleObjects = append(g_VisibleObjects, sphere2, sphere1, planeFloor, sphere3, sphere4 /*, lightSphere*/)
	//g_VisibleObjects = append(g_VisibleObjects, sphere2, planeFloor)
	//g_VisibleObjects = append(g_VisibleObjects, sphere2, sphereFloor)
	imageTexture.Load()
}

func main() {
	//var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

	if g_Options.CpuProfileFilename != "" {
		f, err := os.Create(g_Options.CpuProfileFilename)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	image := generateImage(g_Options.Width, g_Options.Height, computeColorAtXY)
	fmt.Println(g_Options)
	SavePNG(image, g_Options.OutputFilename)

	fmt.Println("Success !")
}

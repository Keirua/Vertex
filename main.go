package main

import (
	"flag"
	"fmt"
	"image/color"
	"math"
	"math/rand"
	"math3d"
)

const (
	DEFAULT_WIDTH              = 640
	DEFAULT_HEIGHT             = 480
	DEFAULT_FOV                = 30.0
	DEFAULT_ANTIALIASING_LEVEL = 1 // minimum
	MAX_DEPTH                  = 3
)

var whiteMaterial = math3d.Material{math3d.Color01{0.8, 0.8, 0.8}, 0.5, 0.5}
var redMaterial = math3d.Material{math3d.Color01{1, 0, 0}, 0.5, 0.5}
var blueMaterial = math3d.Material{math3d.Color01{0, .5, 1}, 0.5, 0.5}
var purpleMaterial = math3d.Material{math3d.Color01{0.65, .2, 0.97}, 0.5,0.5}
var greenMaterial = math3d.Material{math3d.Color01{0.3, 0.9, 0.2}, 0.5,0.5}

var sphereFloor = math3d.Sphere{math3d.Vertex{0, 10003, -20}, 10000.0, whiteMaterial}
var sphere1 = math3d.Sphere{math3d.Vertex{4.0, -1, -5}, 2.0, redMaterial}
var sphere2 = math3d.Sphere{math3d.Vertex{0.0, 0, -15}, 4, greenMaterial}
var sphere3 = math3d.Sphere{math3d.Vertex{3.0, 2, -10}, 3, blueMaterial}
var sphere4 = math3d.Sphere{math3d.Vertex{-5.5, 0, -8}, 3, purpleMaterial}

//var lightSphere = math3d.Sphere{math3d.Vertex{3.0, -3, -10}, 0.5, purpleMaterial}
var light = math3d.Light{math3d.Vertex{3.0, -10, -10}, math3d.Color01{0.65, .6, 0.97}}
var light2 = math3d.Light{math3d.Vertex{0, -5, 0}, math3d.Color01{0.87, 0.8, 0.97}}

var g_Spheres = []math3d.Sphere{sphere2, sphere1, sphereFloor, sphere3, sphere4 /*, lightSphere*/}
var g_Lights = []math3d.Light{/*light, */light2}
var g_Camera math3d.Camera
var g_Options Options

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

func trace(ray math3d.Ray, contributionCoef float64, depth int) math3d.Color01 {
	var finalColor math3d.Color01
	// Find the closest object the ray intersects																														
	var intersectionInfo = getIntersectionInfo(ray)

	if intersectionInfo.ObjectIndex != -1 {
		var objectHit = g_Spheres[intersectionInfo.ObjectIndex]
		var intersectionPoint = ray.VertexAt(intersectionInfo.T)
		var normal = objectHit.ComputeNormalAtPoint(intersectionPoint)
		normal.Normalize()

		// Add Reflection		
		/*var reflectionContributionCoef = contributionCoef * objectHit.Material.ReflectionCoef;
		
		// Computes the reflection ray
		var reflet float64 = 2.0 * (ray.Direction.Dot(normal));
		var reflectedRay math3d.Ray;
		reflectedRay.Origin = intersectionPoint.Add(normal.Mulf(1e-4))
		reflectedRay.Direction = ray.Direction.Substract(normal.Mulf(reflet))
/*
		if ((objectHit.Material.ReflectionCoef >0 || objectHit.Material.Transparency > 0.0 ) && depth < MAX_DEPTH){

			var reflectionColor = trace(reflectedRay, reflectionContributionCoef, depth+1)
			var refractionColor = math3d.Color01{}

			if (objectHit.Material.Transparency > 0) { 
				// Computes the refraction ray
				var refractionRay math3d.Ray;
	            var refractionIndex float64 = 1.1	// indice of refraction ! /!\ should be part of material
	            var eta float64 = 1.0 / refractionIndex
	            var cosi float64 = -normal.Dot(ray.Direction);  // requires normalized vectors
	            var k float64 = 1 - eta * eta * (1 - cosi * cosi);

	            refractionRay.Origin = intersectionPoint.Add(normal.Mulf(1e-4))
	            refractionRay.Direction = ray.Direction.Mulf(eta).Add(normal.Mulf(eta *  cosi - math.Sqrt(k))); 
	            refractionRay.Direction.Normalize(); 

	            refractionColor = trace(refractionRay, reflectionContributionCoef, depth+1)
	        }

			var fresnelCoef = 0.8;
			var reflectionRefractionColorMix = reflectionColor.MulFloat(fresnelCoef).AddColor(refractionColor.MulFloat(1-fresnelCoef));

			finalColor = objectHit.Material.SurfaceColor.AddColor(reflectionRefractionColorMix.MulFloat(reflectionContributionCoef))
		}*/

		// Add Lighting		
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
				var lambert float64 = lightRay.Direction.Dot(normal) * contributionCoef

				// finalColor = finalColor + lambert * currentLight * currentMaterial
				//finalColor = finalColor.AddColor(objectHit.Material.SurfaceColor.MulColor(currLight.Color).MulFloat(lambert))
				finalColor = finalColor.AddColor(objectHit.Material.SurfaceColor.MulColor(currLight.Color).MulFloat(lambert))
			} else {
				// soften the shadow. Total hack, no solid mathematical foundation
				//finalColor = finalColor.AddColor(objectHit.Material.SurfaceColor.MulFloat(0.1))
			}
		}
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

type Options struct {
	Width             int
	Height            int
	Fov               float64
	AntiAliasingLevel int
	OutputFilename    string
}

func init() {
	rand.Seed(42)

	flag.IntVar(&g_Options.Width, "width", DEFAULT_WIDTH, "The output image's width")
	flag.IntVar(&g_Options.Height, "height", DEFAULT_HEIGHT, "The output image's height")
	flag.StringVar(&g_Options.OutputFilename, "output", "out.png", "The output filename")
	flag.IntVar(&g_Options.AntiAliasingLevel, "as", DEFAULT_ANTIALIASING_LEVEL, "AntiAliasingLevel")
	flag.Float64Var(&g_Options.Fov, "fov", DEFAULT_FOV, "FOV, in degre")

	flag.Parse()

	g_Camera.Initialize(g_Options.Width, g_Options.Height, g_Options.Fov)
}

func main() {

	image := generateImage(g_Options.Width, g_Options.Height, computeColorAtXY)
	fmt.Println(g_Options)
	SavePNG(image, g_Options.OutputFilename)

	fmt.Println("Success !")
}

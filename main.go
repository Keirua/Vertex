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

/*
(Sphere(Vec3f( 0.0, -10004, -20), 10000, Vec3f(0.20, 0.20, 0.20), 0, 0.0));
(Sphere(Vec3f( 0.0,      0, -20),     4, Vec3f(1.00, 0.32, 0.36), 1, 0.5));
(Sphere(Vec3f( 5.0,     -1, -15),     2, Vec3f(0.90, 0.76, 0.46), 1, 0.0));
(Sphere(Vec3f( 5.0,      0, -25),     3, Vec3f(0.65, 0.77, 0.97), 1, 0.0));
(Sphere(Vec3f(-5.5,      0, -15),     3, Vec3f(0.90, 0.90, 0.90), 1, 0.0));*/

var whiteMaterial = math3d.Material{math3d.Color01{1, 1, 1}}
var redMaterial = math3d.Material{math3d.Color01{1, 0, 0}}
var blueMaterial = math3d.Material{math3d.Color01{0, .5, 1}}
var darkGrayMaterial = math3d.Material{math3d.Color01{0.65, .77, 0.97}}
var greenMaterial = math3d.Material{math3d.Color01{0.3, 0.9, 0.2}}

var sphereFloor = math3d.Sphere{math3d.Vertex{0, 10003, -20}, 10000.0, whiteMaterial}
var sphere1 = math3d.Sphere{math3d.Vertex{4.0, -1, -5}, 2.0, redMaterial}
var sphere2 = math3d.Sphere{math3d.Vertex{0.0, 0, -15}, 4, blueMaterial}
var sphere3 = math3d.Sphere{math3d.Vertex{3.0, 2, -10}, 3, darkGrayMaterial}
var sphere4 = math3d.Sphere{math3d.Vertex{-5.5, 0, -15}, 3, greenMaterial}

var g_Spheres = []math3d.Sphere{sphere2, sphere1, sphereFloor, sphere3, sphere4}
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
	SaveJPEG(image, "out.jpg")
	fmt.Println("Success !")
}

/*
void render(const std::vector<Sphere> &spheres)
{
    unsigned width = 640, height = 480;
    Vec3f *image = new Vec3f[width * height], *pixel = image;
    float invWidth = 1 / float(width), invHeight = 1 / float(height);
    float fov = 30, aspectratio = width / float(height);
    float angle = tan(M_PI * 0.5 * fov / 180.);
    // Trace rays
    for (unsigned y = 0; y < height; ++y) {
        for (unsigned x = 0; x < width; ++x, ++pixel) {
            float xx = (2 * ((x + 0.5) * invWidth) - 1) * angle * aspectratio;
            float yy = (1 - 2 * ((y + 0.5) * invHeight)) * angle;
            Vec3f raydir(xx, yy, -1);
            raydir.normalize();
            *pixel = trace(Vec3f(0), raydir, spheres, 0);
        }
    }
    // Save result to a PPM image (keep these flags if you compile under Windows)

    delete [] image;
}


int main(int argc, char **argv)
{
    srand48(13);
    std::vector<Sphere> spheres;
    // position, radius, surface color, reflectivity, transparency, emission color

    // light
    spheres.push_back(Sphere(Vec3f( 0.0,     20, -30),     3, Vec3f(0.00, 0.00, 0.00), 0, 0.0, Vec3f(3)));
    render(spheres);

    return 0;
}
*/

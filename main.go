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
	fov    = 30
)

func computeRayDirection(x int, y int) math3d.Ray {
	invWidth := 1.0 / float64(width)
	invHeight := 1.0 / float64(height)
	aspectratio := width / float64(height)
	angle := math.Tan(math.Pi * 0.5 * fov * 180.0)

	xx := float64(2*((float64(x)+0.5)*invWidth)-1) * angle * aspectratio
	yy := float64(1-2*((float64(y)+0.5)*invHeight)) * angle

	var direction = math3d.Vertex{xx, yy, -1}
	direction.Normalize()

	return math3d.Ray{math3d.Vertex{}, direction}
}

func trace(ray math3d.Ray) bool {
	//svar sph math3d.Sphere = math3d.Sphere{math3d.Vertex{5.0, -1, -15}, 2.0, math3d.Vertex{0.90, 0.76, 0.76}}
	return (bool(rand.Intn(2) > 0))
	//return sph.Intersect(ray)
}

func computeColorAtXY(x int, y int) color.RGBA {
	var ray = computeRayDirection(x, y)

	var value uint8 = 0
	if trace(ray) {
		value = 0xFF
	}

	return color.RGBA{value, value, value, 0xFF}
}

func main() {
    rand.Seed(42)
	image := generateImage(width, height, computeColorAtXY)
	saveImage(image, "out.jpg")
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
    std::ofstream ofs("./untitled.ppm", std::ios::out | std::ios::binary);
    ofs << "P6\n" << width << " " << height << "\n255\n";
    for (unsigned i = 0; i < width * height; ++i) {
        ofs << (unsigned char)(std::min(float(1), image[i].x) * 255) <<
               (unsigned char)(std::min(float(1), image[i].y) * 255) <<
               (unsigned char)(std::min(float(1), image[i].z) * 255);
    }
    ofs.close();
    delete [] image;
}


int main(int argc, char **argv)
{
    srand48(13);
    std::vector<Sphere> spheres;
    // position, radius, surface color, reflectivity, transparency, emission color
    spheres.push_back(Sphere(Vec3f( 0.0, -10004, -20), 10000, Vec3f(0.20, 0.20, 0.20), 0, 0.0));
    spheres.push_back(Sphere(Vec3f( 0.0,      0, -20),     4, Vec3f(1.00, 0.32, 0.36), 1, 0.5));
    spheres.push_back(Sphere(Vec3f( 5.0,     -1, -15),     2, Vec3f(0.90, 0.76, 0.46), 1, 0.0));
    spheres.push_back(Sphere(Vec3f( 5.0,      0, -25),     3, Vec3f(0.65, 0.77, 0.97), 1, 0.0));
    spheres.push_back(Sphere(Vec3f(-5.5,      0, -15),     3, Vec3f(0.90, 0.90, 0.90), 1, 0.0));
    // light
    spheres.push_back(Sphere(Vec3f( 0.0,     20, -30),     3, Vec3f(0.00, 0.00, 0.00), 0, 0.0, Vec3f(3)));
    render(spheres);

    return 0;
}
*/

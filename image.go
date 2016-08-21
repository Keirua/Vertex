package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
)

func generateImage(width int, height int, computePixel func(i int, j int) color.RGBA) *image.RGBA {
	m := image.NewRGBA(image.Rect(0, 0, width, height))
	b := m.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			colorValue := computePixel(x, y)
			m.Set(x, y, colorValue)
		}
	}
	return m
}

func SaveJPEG(img *image.RGBA, filename string) {
	file, err := os.Create(filename)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	jpeg.Encode(file, img, &jpeg.Options{100})
}

func SavePPM(img *image.RGBA, filename string) {
	file, err := os.Create(filename)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	// jpeg.Encode(file, img, &jpeg.Options{100})
}

/*
std::ofstream ofs("./untitled.ppm", std::ios::out | std::ios::binary);
ofs << "P6\n" << width << " " << height << "\n255\n";
for (unsigned i = 0; i < width * height; ++i) {
    ofs << (unsigned char)(std::min(float(1), image[i].x) * 255) <<
           (unsigned char)(std::min(float(1), image[i].y) * 255) <<
           (unsigned char)(std::min(float(1), image[i].z) * 255);
}
ofs.close();
*/

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

func saveImage(img *image.RGBA, filename string) {
	file, err := os.Create(filename)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	jpeg.Encode(file, img, &jpeg.Options{100})
}

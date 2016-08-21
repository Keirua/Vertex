package main

import (
    "image"
    "image/jpeg"
    "image/color"
    "log"
    "os"
)

func saveImage (img *image.RGBA, filename string){
    file, err := os.Create(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    jpeg.Encode(file, img, &jpeg.Options{80})
}

func createImage (width int, height int) *image.RGBA {
    m := image.NewRGBA(image.Rect(0, 0, width, height))
    b := m.Bounds()
    for y := b.Min.Y; y < b.Max.Y; y++ {
        for x := b.Min.X; x < b.Max.X; x++ {
            value := uint8(x*y)
            m.Set(x, y, color.RGBA{value, value, value, 0xFF})
        }
    }
    return m
}

func generateImage (width int, height int, computePixel func(i int, j int) color.RGBA) *image.RGBA {
    m := image.NewRGBA(image.Rect(0, 0, width, height))
    b := m.Bounds()
    for y := b.Min.Y; y < b.Max.Y; y++ {
        for x := b.Min.X; x < b.Max.X; x++ {
            colorValue := computePixel(x, y)
            //value := uint8(x*y)
            m.Set(x, y, colorValue)
        }
    }
    return m
}
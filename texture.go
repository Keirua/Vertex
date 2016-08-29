package main

import (
    "os"
    "image"
    // We want to init the decoders  for png/jpg so that image file loading
    // works out of the box
    _ "image/png"
    _ "image/jpeg"
)

type Texture interface {
    GetColor01AtUV(u float64, v float64) Color01
}

type CheckboardTexture struct {
    W int
    H int
}

func (checkboardTexture CheckboardTexture) GetColor01AtUV(u float64, v float64) Color01 {
    var xAtU int = int(u*float64(checkboardTexture.W))
    var yAtV int = int(v*float64(checkboardTexture.H))

    if (((xAtU + yAtV)%2) != 0){
        return Color01{1,1,1}
    }

    return Color01{0,0,0}
}

type FileTexture struct {
    Filename string
    Image image.Image
}

func (fileTexture *FileTexture) Load() {
    // https://golang.org/pkg/os/#Open
    fImg1, _ := os.Open(fileTexture.Filename)
    defer fImg1.Close()
    // https://golang.org/pkg/image/#Decode
    img1, _, _ := image.Decode(fImg1)

    fileTexture.Image = img1
}

func (fileTexture FileTexture) GetColor01AtUV(u float64, v float64) Color01 {
    var bounds = fileTexture.Image.Bounds()
    var width = bounds.Max.X-bounds.Min.X
    var height = bounds.Max.Y-bounds.Min.Y

    var xAtU int = int(u*float64(width))
    var yAtV int = int(v*float64(height))

    var colorRGBA = fileTexture.Image.At(xAtU, yAtV)
    var R,G,B,_ = colorRGBA.RGBA()

    return Color01{float64(R)/(255.0*255.0),float64(G)/(255.0*255.0),float64(B)/(255.0*255.0)}
}
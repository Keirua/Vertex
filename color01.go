package main

import "image/color"

type Color01 struct {
    R float64
    G float64
    B float64
}

func (col Color01) ToRGBA() color.RGBA {
    return color.RGBA{uint8(col.R * 255), uint8(col.G * 255), uint8(col.B * 255), 255}
}

func (col Color01) AddColor(col2 Color01) Color01 {
    col.R = col.R + col2.R
    col.G = col.G + col2.G
    col.B = col.B + col2.B
    return col
}

// f should be >= 0 and <= 1
func (col Color01) MulFloat(f float64) Color01 {
    col.R = col.R * f
    col.G = col.G * f
    col.B = col.B * f

    return col
}

func (col Color01) MulColor(col2 Color01) Color01 {
    col.R = col.R * col2.R
    col.G = col.G * col2.G
    col.B = col.B * col2.B

    return col
}
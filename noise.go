package main

// The code in this file is mostly a port of the code on this page :
// http://lodev.org/cgtutor/randomnoise.html
// The article contains very good explanations regarding the techniques involved !

import (
	/*"image"
	  "image/color"
	  "image/png"
	  "log"
	  "os"*/
	"math"
	"math/rand"
)

/*
Here are some good values for the following objects :

var turbulence = NewTurbulence(32)
var marble5 = NewMarble(5, 10, 5, 32)
var marble1 = NewMarble(5, 10, 1, 32)
var wood = NewWood(12, 0.1, 32)
*/

type Noise struct {
	noise [256][256]float64
}

func (noise *Noise) generateNoise() {
	for i := 0; i < 256; i++ {
		for j := 0; j < 256; j++ {
			noise.noise[j][i] = rand.Float64()
		}
	}
}

func NewNoise() *Noise {
	var n = Noise{}
	n.generateNoise()
	return &n
}

func (noise *Noise) SmoothNoise(x float64, y float64) float64 {
	var noiseWidth = 256
	var noiseHeight = 256

	//get fractional part of x and y
	var fractX float64 = float64(x - math.Floor(x))
	var fractY float64 = float64(y - math.Floor(y))

	//wrap around
	var x1 int = (int(x) + noiseWidth) % noiseWidth
	var y1 int = (int(y) + noiseHeight) % noiseHeight

	//neighbor values
	var x2 int = (x1 + noiseWidth - 1) % noiseWidth
	var y2 int = (y1 + noiseHeight - 1) % noiseHeight

	//smooth the noise with bilinear interpolation
	var value float64 = 0.0
	value += fractX * fractY * noise.noise[y1][x1]
	value += (1 - fractX) * fractY * noise.noise[y1][x2]
	value += fractX * (1 - fractY) * noise.noise[y2][x1]
	value += (1 - fractX) * (1 - fractY) * noise.noise[y2][x2]

	return value
}

func (noise *Noise) At(i float64, j float64) float64 {
	//return noise.SmoothNoise(i,j)
	return noise.noise[int(j)][int(i)]
}

func (noise *Noise) GetColor01AtUV(u float64, v float64) Color01 {
	var c = noise.At(u*256.0, v*256.0)
	return Color01{c, c, c}
}

type Turbulence struct {
	size  float64
	noise *Noise
}

func NewTurbulence(size float64) *Turbulence {
	var n = Noise{}
	n.generateNoise()
	var turbulence = Turbulence{size, &n}
	return &turbulence
}

func (turbulence Turbulence) At(x float64, y float64) float64 {
	var value float64 = 0.0
	var size = turbulence.size
	var initialSize = size

	for size >= 1 {
		value += turbulence.noise.SmoothNoise(x/size, y/size) * size
		size *= 0.5
	}

	return (0.5 * value / initialSize)
}

func (turbulence *Turbulence) GetColor01AtUV(u float64, v float64) Color01 {
	var c = turbulence.At(u*256.0, v*256)
	return Color01{c, c, c}
}

type Marble struct {
	xPeriod         float64
	yPeriod         float64
	turbulencePower float64
	turbulenceSize  float64
	turbulence      *Turbulence
}

func NewMarble(xPeriod, yPeriod, turbulencePower, turbulenceSize float64) *Marble {
	var t = NewTurbulence(turbulenceSize)
	var marble = Marble{xPeriod, yPeriod, turbulencePower, turbulenceSize, t}

	return &marble
}

func (marble *Marble) At(x, y float64) float64 {
	var noiseWidth = 256
	var noiseHeight = 256
	var xyValue float64 = x*marble.xPeriod/float64(noiseWidth) + y*marble.yPeriod/float64(noiseHeight) + marble.turbulencePower*marble.turbulence.At(x, y)
	var sineValue float64 = math.Abs(math.Sin(xyValue * math.Pi))

	return sineValue
}

func (marble *Marble) GetColor01AtUV(u float64, v float64) Color01 {
	var c = marble.At(u*256.0, v*256)
	return Color01{c, c, c}
}

type Wood struct {
	xyPeriod        float64
	turbulencePower float64
	turbulenceSize  float64
	turbulence      *Turbulence
}

func NewWood(xyPeriod, turbulencePower, turbulenceSize float64) *Wood {
	var t = NewTurbulence(turbulenceSize)
	var wood = Wood{xyPeriod, turbulencePower, turbulenceSize, t}

	return &wood
}

func (wood *Wood) At(x, y float64) float64 {
	var noiseWidth = 256
	var noiseHeight = 256

	var xValue float64 = (x - 0.5*float64(noiseWidth)) / float64(noiseWidth)
	var yValue float64 = (y - 0.5*float64(noiseHeight)) / float64(noiseHeight)
	var distValue float64 = math.Sqrt(xValue*xValue + yValue*yValue)
	var displacement = wood.turbulencePower * wood.turbulence.At(x, y)
	var sineValue float64 = math.Abs(math.Sin(2 * wood.xyPeriod * (distValue + displacement) * math.Pi))

	return sineValue
}

func (wood *Wood) GetColor01AtUV(u float64, v float64) Color01 {
	var c = wood.At(u*256.0, v*256.0)
	return Color01{c, c, c}
}

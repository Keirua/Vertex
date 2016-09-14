package main

import (
	"flag"
)

type Options struct {
	Width              int
	Height             int
	Fov                float64
	AntiAliasingLevel  int
	MaxDepth           int
	OutputFilename     string
	CpuProfileFilename string
	// if no soft shadows : num = 1, strength = 0
	// if soft shadows : num = 16, strength = 0.2 for instance
	NumSoftShadowRays	int
	SoftShadowStrength	float64
}

func (options *Options) ParseCommandLineOptions() {
	flag.IntVar(&options.Width, "width", DEFAULT_WIDTH, "The output image's width")
	flag.IntVar(&options.Height, "height", DEFAULT_HEIGHT, "The output image's height")
	flag.StringVar(&options.OutputFilename, "output", "out.png", "The output filename")
	flag.IntVar(&options.AntiAliasingLevel, "as", DEFAULT_ANTIALIASING_LEVEL, "Antialiasing level")
	flag.Float64Var(&options.Fov, "fov", DEFAULT_FOV, "FOV, in degre")
	flag.IntVar(&options.MaxDepth, "depth", MAX_DEPTH, "max recursion")
	flag.Float64Var(&options.SoftShadowStrength, "softShadowStrength", 0.0, "amount of displacement for the soft shadow rays")
	flag.IntVar(&options.NumSoftShadowRays, "nbSoftShadowRays", 1, "nb of rays to soft for soft shadows (16 is good)")
	flag.StringVar(&options.CpuProfileFilename, "cpuprofile", "", "cpu profile for debug")

	flag.Parse()
}

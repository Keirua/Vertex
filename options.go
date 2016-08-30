package main

import (
    "flag"
)

type Options struct {
    Width             int
    Height            int
    Fov               float64
    AntiAliasingLevel int
    MaxDepth int
    OutputFilename    string
    CpuProfileFilename string
}

func (options *Options) ParseCommandLineOptions() {
    flag.IntVar(&options.Width, "width", DEFAULT_WIDTH, "The output image's width")
    flag.IntVar(&options.Height, "height", DEFAULT_HEIGHT, "The output image's height")
    flag.StringVar(&options.OutputFilename, "output", "out.png", "The output filename")
    flag.IntVar(&options.AntiAliasingLevel, "as", DEFAULT_ANTIALIASING_LEVEL, "Antialiasing level")
    flag.Float64Var(&options.Fov, "fov", DEFAULT_FOV, "FOV, in degre")
    flag.IntVar(&options.MaxDepth, "depth", MAX_DEPTH, "max recursion")
    flag.StringVar(&options.CpuProfileFilename, "cpuprofile", "", "cpu profile for debug")

    flag.Parse()
}
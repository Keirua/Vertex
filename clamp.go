package main

import (
	"math"
)

// Clamp method to ensure the final colors are in the [0;1] range
type Clampable interface {
	Clamp(v float64) float64
}

type ClampMinMax struct {
}

func (c *ClampMinMax) Clamp(v float64) float64 {
	return math.Min(1.0, math.Max(0.0, v))
}

type ClampExponential struct {
	Coef float64
}

func (c ClampExponential) Clamp(v float64) float64 {
	// var coef float64 = -2.0
	return 1 - math.Exp(v*c.Coef)
}

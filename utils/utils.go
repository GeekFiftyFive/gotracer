package utils

import (
	"math"
	"math/rand/v2"
)

// Utility functions

func DegreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func RandomFloat() float64 {
	return rand.Float64()
}

func RandomRange(min, max float64) float64 {
	return min + RandomFloat()*(max-min)
}

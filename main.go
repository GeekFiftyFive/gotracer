package main

import (
	"gotracer/camera"
	"gotracer/color"
	"gotracer/material"
	"gotracer/sphere"
	"gotracer/vector"
	"math"
)

func main() {
	// World
	world := material.HittableList{}

	r := math.Cos(math.Pi / 4.0)

	materialLeft := material.NewLambertian(color.NewColor(0, 0, 1))
	materialRight := material.NewLambertian(color.NewColor(1, 0, 0))

	sphereLeft := sphere.NewSphere(vector.NewVector3(-r, 0, -1), r, &materialLeft)
	sphereRight := sphere.NewSphere(vector.NewVector3(r, 0, -1), r, &materialRight)

	world.Add(&sphereLeft)
	world.Add(&sphereRight)

	cam := camera.Camera{AspectRatio: 16.0 / 9.0, ImageWidth: 400, SamplesPerPixel: 100, MaxDepth: 50, Fov: 90}
	cam.Render(&world)
}

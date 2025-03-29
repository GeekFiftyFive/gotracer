package main

import (
	"gotracer/camera"
	"gotracer/hittable"
	"gotracer/sphere"
	"gotracer/vector"
)

func main() {
	// World
	world := hittable.HittableList{}
	sphere1 := sphere.NewSphere(vector.NewVector3(0, 0, -1), 0.5)
	sphere2 := sphere.NewSphere(vector.NewVector3(0, -100.5, -1), 100)
	world.Add(&sphere1)
	world.Add(&sphere2)
	cam := camera.Camera{AspectRatio: 16.0 / 9.0, ImageWidth: 400, SamplesPerPixel: 100}
	cam.Render(&world)
}

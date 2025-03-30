package main

import (
	"gotracer/camera"
	"gotracer/color"
	"gotracer/material"
	"gotracer/sphere"
	"gotracer/vector"
)

func main() {
	// World
	world := material.HittableList{}

	materialGround := material.NewLambertian(color.NewColor(0.8, 0.8, 0.0))
	materialCenter := material.NewLambertian(color.NewColor(0.1, 0.2, 0.5))
	materialLeft := material.NewDialectric(1.5)
	materialBubble := material.NewDialectric(1.0 / 1.5)
	materialRight := material.NewMetal(color.NewColor(0.8, 0.6, 0.2), 1.0)

	sphereGround := sphere.NewSphere(vector.NewVector3(0.0, -100.5, -1.0), 100.0, &materialGround)
	sphereCenter := sphere.NewSphere(vector.NewVector3(0.0, 0.0, -1.2), 0.5, &materialCenter)
	sphereLeft := sphere.NewSphere(vector.NewVector3(-1.0, 0.0, -1.0), 0.5, &materialLeft)
	sphereBubble := sphere.NewSphere(vector.NewVector3(-1.0, 0.0, -1.0), 0.4, &materialBubble)
	sphereRight := sphere.NewSphere(vector.NewVector3(1.0, 0.0, -1.0), 0.5, &materialRight)

	world.Add(&sphereGround)
	world.Add(&sphereCenter)
	world.Add(&sphereLeft)
	world.Add(&sphereBubble)
	world.Add(&sphereRight)

	cam := camera.Camera{
		AspectRatio:     16.0 / 9.0,
		ImageWidth:      400,
		SamplesPerPixel: 100,
		MaxDepth:        50,
		Fov:             20,
		LookFrom:        vector.NewVector3(-2, 2, 1),
		LookAt:          vector.NewVector3(0, 0, -1),
		Vup:             vector.NewVector3(0, 1, 0),
	}
	cam.Render(&world)
}

package main

import (
	"gotracer/camera"
	"gotracer/color"
	"gotracer/geometry"
	"gotracer/material"
	"gotracer/utils"
	"gotracer/vector"
)

func main() {
	// World
	world := material.HittableList{}

	groundMaterial := material.NewLambertian(color.NewColor(0.5, 0.5, 0.5))
	geometry.NewCuboid(vector.NewVector3(0, 0, 0), vector.NewVector3(10, 10, 10), &groundMaterial)
	groundSphere := geometry.NewSphere(vector.NewVector3(0, -1000, 0), 1000, &groundMaterial)
	world.Add(&groundSphere)

	for i := range 22 {
		a := float64(i - 11)
		for j := range 22 {
			b := float64(j - 11)
			chooseMat := utils.RandomFloat()
			center := vector.NewVector3(a+0.9*utils.RandomFloat(), 0.2, b+0.9*utils.RandomFloat())

			if center.SubtractVector(vector.NewVector3(4, 0.2, 0)).Length() > 0.9 {
				var sphereMaterial material.Material
				if chooseMat < 0.9 {
					// Diffuse
					albedo := color.Random().MultiplyVector(color.Random())
					lambertian := material.NewLambertian(albedo)
					sphereMaterial = &lambertian
				} else if chooseMat < 0.95 {
					// Metal
					albedo := color.RandomRange(0.5, 1)
					fuzz := utils.RandomRange(0, 0.5)
					metal := material.NewMetal(albedo, fuzz)
					sphereMaterial = &metal
				} else {
					// Glass
					glass := material.NewDialectric(1.5)
					sphereMaterial = &glass
				}
				sphere := geometry.NewSphere(center, 0.2, sphereMaterial)
				world.Add(&sphere)
			}
		}
	}

	material1 := material.NewDialectric(1.5)
	sphere1 := geometry.NewSphere(vector.NewVector3(0, 1, 0), 1.0, &material1)
	world.Add(&sphere1)

	material2 := material.NewLambertian(color.NewColor(0.4, 0.2, 0.1))
	sphere2 := geometry.NewSphere(vector.NewVector3(-4, 1, 0), 1.0, &material2)
	world.Add(&sphere2)

	material3 := material.NewMetal(color.NewColor(0.7, 0.6, 0.5), 0.0)
	sphere3 := geometry.NewSphere(vector.NewVector3(4, 1, 0), 1.0, &material3)
	world.Add(&sphere3)

	cam := camera.Camera{
		AspectRatio:     16.0 / 9.0,
		ImageWidth:      400,
		SamplesPerPixel: 10,
		MaxDepth:        50,
		Fov:             20,
		LookFrom:        vector.NewVector3(13, 2, 3),
		LookAt:          vector.NewVector3(0, 0, 0),
		Vup:             vector.NewVector3(0, 1, 0),
		DefocusAngle:    0.6,
		FocusDist:       10.0,
	}
	cam.Render(&world)
}

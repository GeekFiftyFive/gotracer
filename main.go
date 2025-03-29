package main

import (
	"bytes"
	"fmt"
	"gotracer/color"
	"gotracer/ray"
	"gotracer/vector"
	"log"
	"os"
)

func hitSphere(center vector.Point3, radius float64, r ray.Ray) bool {
	oc := center.SubtractVector(r.Origin())
	a := r.Direction().Dot(r.Direction())
	b := r.Direction().Dot(oc) * -2.0
	c := oc.Dot(oc) - (radius * radius)
	discriminant := b*b - 4*a*c
	return discriminant >= 0
}

func rayColor(r ray.Ray) color.Color {
	if hitSphere(vector.NewVector3(0, 0, -1), 0.5, r) {
		return color.NewColor(1, 0, 0)
	}

	unitDirection := r.Direction().UnitVector()
	a := 0.5 * (unitDirection.Y() + 1.0)
	return color.NewColor(1.0, 1.0, 1.0).
		MultiplyFloat(1.0 - a).
		AddVector(color.NewColor(0.5, 0.7, 1.0).MultiplyFloat(a))
}

func main() {
	// Logger
	logger := log.New(os.Stderr, "", 0)

	// Set image parameters
	aspectRatio := 16.0 / 9.0
	imageWidth := 400

	// Calculate image height
	imageHeight := max(int(float64(imageWidth)/aspectRatio), 1)

	// Setup camera
	focalLength := 1.0
	viewportHeight := 2.0
	viewportWidth := viewportHeight * (float64(imageWidth) / float64(imageHeight))
	var cameraCenter vector.Point3 = vector.NewVector3(0, 0, 0)

	// Calculate the vectors across the horizontal and down the vertical viewport edges
	viewportU := vector.NewVector3(viewportWidth, 0, 0)
	viewportV := vector.NewVector3(0, -viewportHeight, 0)

	// Calculate the horizontal and vertical delta vectors from pixel to pixel
	pixelDeltaU := viewportU.DivideFloat(float64(imageWidth))
	pixelDeltaV := viewportV.DivideFloat(float64(imageHeight))

	// Calculate the location of the upper left pixel
	viewportUpperLeft := cameraCenter.
		SubtractVector(vector.NewVector3(0, 0, focalLength)).
		SubtractVector(viewportU.DivideFloat(2.0)).
		SubtractVector(viewportV.DivideFloat(2.0))

	pixel00Loc := viewportUpperLeft.AddVector(pixelDeltaU.AddVector(pixelDeltaV).MultiplyFloat(0.5))

	// Render
	var writer bytes.Buffer
	writer.Write(fmt.Appendf(nil, "P3\n%d %d\n255\n", imageWidth, imageHeight))

	for j := range imageHeight {
		logger.Printf("Scanlines remaining: %d\n", imageHeight-j)
		for i := range imageWidth {
			pixelCenter := pixel00Loc.
				AddVector(pixelDeltaU.MultiplyFloat(float64(i))).
				AddVector(pixelDeltaV.MultiplyFloat(float64(j)))
			rayDirection := pixelCenter.SubtractVector(cameraCenter)

			r := ray.NewRay(cameraCenter, rayDirection)

			pixelColor := rayColor(r)

			color.WriteColor(pixelColor, &writer)
		}
	}

	fmt.Print(writer.String())
}

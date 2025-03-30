package camera

import (
	"bytes"
	"fmt"
	"gotracer/color"
	"gotracer/interval"
	"gotracer/material"
	"gotracer/ray"
	"gotracer/utils"
	"gotracer/vector"
	"log"
	"math"
	"os"
)

type Camera struct {
	AspectRatio       float64
	ImageWidth        int
	SamplesPerPixel   int
	MaxDepth          int
	Fov               float64
	LookFrom          vector.Point3
	LookAt            vector.Point3
	Vup               vector.Vector3
	imageHeight       int
	center            vector.Point3
	pixel00Loc        vector.Point3
	pixelDeltaU       vector.Vector3
	pixelDeltaV       vector.Vector3
	u                 vector.Vector3
	v                 vector.Vector3
	w                 vector.Vector3
	pixelSamplesScale float64
}

func (c *Camera) Render(world material.Hittable) {
	logger := log.New(os.Stderr, "", 0) // TODO: Move logger into its own package
	c.initialize()

	var writer bytes.Buffer
	writer.Write(fmt.Appendf(nil, "P3\n%d %d\n255\n", c.ImageWidth, c.imageHeight))

	for j := range c.imageHeight {
		logger.Printf("Scanlines remaining: %d\n", c.imageHeight-j)
		for i := range c.ImageWidth {
			pixelColor := color.NewColor(0, 0, 0)
			for range c.SamplesPerPixel {
				r := c.getRay(i, j)
				pixelColor = pixelColor.AddVector(c.rayColor(r, c.MaxDepth, world))
			}

			color.WriteColor(pixelColor.MultiplyFloat(c.pixelSamplesScale), &writer)
		}
	}

	fmt.Print(writer.String())
}

func (c *Camera) initialize() {
	// Calculate image height
	c.imageHeight = max(int(float64(c.ImageWidth)/c.AspectRatio), 1)

	c.pixelSamplesScale = 1.0 / float64(c.SamplesPerPixel)

	c.center = c.LookFrom

	// Determine viewport dimensions
	focalLength := c.LookFrom.SubtractVector(c.LookAt).Length()
	theta := utils.DegreesToRadians(c.Fov)
	h := math.Tan(theta / 2.0)
	viewportHeight := 2.0 * h * focalLength
	viewportWidth := viewportHeight * (float64(c.ImageWidth) / float64(c.imageHeight))

	// Calculate the u,v,w unit basis vectors for the camera coordinate frame.
	w := c.LookFrom.SubtractVector(c.LookAt).UnitVector()
	u := c.Vup.Cross(w).UnitVector()
	v := w.Cross(u)

	// Calculate the vectors across the horizontal and down the vertical viewport edges
	viewportU := u.MultiplyFloat(viewportWidth)
	viewportV := v.MultiplyFloat(-1 * viewportHeight)

	// Calculate the horizontal and vertical delta vectors from pixel to pixel
	c.pixelDeltaU = viewportU.DivideFloat(float64(c.ImageWidth))
	c.pixelDeltaV = viewportV.DivideFloat(float64(c.imageHeight))

	// Calculate the location of the upper left pixel
	viewportUpperLeft := c.center.
		SubtractVector(w.MultiplyFloat(focalLength)).
		SubtractVector(viewportU.DivideFloat(2)).
		SubtractVector(viewportV.DivideFloat(2))

	c.pixel00Loc = viewportUpperLeft.AddVector(c.pixelDeltaU.AddVector(c.pixelDeltaV).MultiplyFloat(0.5))
}

func (c *Camera) rayColor(r ray.Ray, depth int, world material.Hittable) color.Color {
	if depth <= 0 {
		return color.NewColor(0, 0, 0)
	}
	isHit, rec := world.Hit(r, interval.Interval{Min: 0.001, Max: math.Inf(+1)})
	if isHit {
		isScatter, scattered, attenuation := rec.Mat.Scatter(r, rec)
		if isScatter {
			return c.rayColor(scattered, depth-1, world).MultiplyVector(attenuation)
		}
		return color.NewColor(0, 0, 0)
	}
	unitDirection := r.Direction().UnitVector()
	a := 0.5 * (unitDirection.Y() + 1.0)
	return color.NewColor(1.0, 1.0, 1.0).
		MultiplyFloat(1.0 - a).
		AddVector(color.NewColor(0.5, 0.7, 1.0).MultiplyFloat(a))
}

func (c *Camera) getRay(i, j int) ray.Ray {
	// Construct a camera ray originating from the origin and directed at randomly sampled
	// point around the pixel location i, j.
	offset := sampleSquare()
	pixelSample := c.pixel00Loc.
		AddVector(c.pixelDeltaU.MultiplyFloat(float64(i) + (offset.X()))).
		AddVector(c.pixelDeltaV.MultiplyFloat(float64(j) + offset.Y()))
	rayOrigin := c.center
	rayDirection := pixelSample.SubtractVector(rayOrigin)
	return ray.NewRay(rayOrigin, rayDirection)
}

func sampleSquare() vector.Vector3 {
	return vector.NewVector3(utils.RandomFloat()-0.5, utils.RandomFloat()-0.5, 0)
}

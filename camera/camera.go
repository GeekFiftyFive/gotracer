package camera

import (
	"bytes"
	"fmt"
	"gotracer/color"
	"gotracer/hittable"
	"gotracer/interval"
	"gotracer/ray"
	"gotracer/vector"
	"log"
	"math"
	"os"
)

type Camera struct {
	AspectRatio float64
	ImageWidth  int
	imageHeight int
	center      vector.Point3
	pixel00Loc  vector.Point3
	pixelDeltaU vector.Vector3
	pixelDeltaV vector.Vector3
}

func (c *Camera) Render(world hittable.Hittable) {
	logger := log.New(os.Stderr, "", 0) // TODO: Move logger into its own package
	c.initialize()

	var writer bytes.Buffer
	writer.Write(fmt.Appendf(nil, "P3\n%d %d\n255\n", c.ImageWidth, c.imageHeight))

	for j := range c.imageHeight {
		logger.Printf("Scanlines remaining: %d\n", c.imageHeight-j)
		for i := range c.ImageWidth {
			pixelCenter := c.pixel00Loc.
				AddVector(c.pixelDeltaU.MultiplyFloat(float64(i))).
				AddVector(c.pixelDeltaV.MultiplyFloat(float64(j)))
			rayDirection := pixelCenter.SubtractVector(c.center)

			r := ray.NewRay(c.center, rayDirection)

			pixelColor := c.rayColor(r, world)

			color.WriteColor(pixelColor, &writer)
		}
	}

	fmt.Print(writer.String())
}

func (c *Camera) initialize() {
	// Calculate image height
	c.imageHeight = max(int(float64(c.ImageWidth)/c.AspectRatio), 1)

	c.center = vector.NewVector3(0, 0, 0)

	// Determine viewport dimensions
	focalLength := 1.0
	viewportHeight := 2.0
	viewportWidth := viewportHeight * (float64(c.ImageWidth) / float64(c.imageHeight))
	var cameraCenter vector.Point3 = vector.NewVector3(0, 0, 0)

	// Calculate the vectors across the horizontal and down the vertical viewport edges
	viewportU := vector.NewVector3(viewportWidth, 0, 0)
	viewportV := vector.NewVector3(0, -viewportHeight, 0)

	// Calculate the horizontal and vertical delta vectors from pixel to pixel
	c.pixelDeltaU = viewportU.DivideFloat(float64(c.ImageWidth))
	c.pixelDeltaV = viewportV.DivideFloat(float64(c.imageHeight))

	// Calculate the location of the upper left pixel
	viewportUpperLeft := cameraCenter.
		SubtractVector(vector.NewVector3(0, 0, focalLength)).
		SubtractVector(viewportU.DivideFloat(2.0)).
		SubtractVector(viewportV.DivideFloat(2.0))

	c.pixel00Loc = viewportUpperLeft.AddVector(c.pixelDeltaU.AddVector(c.pixelDeltaV).MultiplyFloat(0.5))
}

func (c *Camera) rayColor(r ray.Ray, world hittable.Hittable) color.Color {
	isHit, rec := world.Hit(r, interval.Interval{Min: 0, Max: math.Inf(+1)})
	if isHit {
		return rec.Normal.AddVector(color.NewColor(1, 1, 1)).MultiplyFloat(0.5)
	}
	unitDirection := r.Direction().UnitVector()
	a := 0.5 * (unitDirection.Y() + 1.0)
	return color.NewColor(1.0, 1.0, 1.0).
		MultiplyFloat(1.0 - a).
		AddVector(color.NewColor(0.5, 0.7, 1.0).MultiplyFloat(a))
}

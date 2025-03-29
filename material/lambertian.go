package material

import (
	"gotracer/color"
	"gotracer/ray"
	"gotracer/vector"
)

type Lambertian struct {
	albedo color.Color
}

func NewLambertian(albedo color.Color) Lambertian {
	return Lambertian{albedo}
}

func (l *Lambertian) Scatter(in ray.Ray, rec *HitRecord) (isScatter bool, scattered ray.Ray, attenuation color.Color) {
	isScatter = true
	scatterDirection := rec.Normal.AddVector(vector.RandomUnitVector())

	// Catch degenerate scatter direction
	if scatterDirection.NearZero() {
		scatterDirection = rec.Normal
	}

	scattered = ray.NewRay(rec.P, scatterDirection)
	attenuation = l.albedo
	return
}

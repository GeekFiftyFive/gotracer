package material

import (
	"gotracer/color"
	"gotracer/ray"
	"gotracer/utils"
	"gotracer/vector"
	"math"
)

type Dialectric struct {
	refractionIndex float64
}

func NewDialectric(refractionIndex float64) Dialectric {
	return Dialectric{refractionIndex}
}

func (d *Dialectric) Scatter(in ray.Ray, rec *HitRecord) (isScatter bool, scattered ray.Ray, attenuation color.Color) {
	isScatter = true
	attenuation = color.NewColor(1.0, 1.0, 1.0)
	var ri float64
	if rec.FrontFace {
		ri = 1.0 / d.refractionIndex
	} else {
		ri = d.refractionIndex
	}

	unitDirection := in.Direction().UnitVector()
	cosTheta := math.Min(unitDirection.MultiplyFloat(-1).Dot(rec.Normal), 1)
	sinTheta := math.Sqrt(1.0 - cosTheta*cosTheta)

	cannotRefract := ri*sinTheta > 1.0
	var direction vector.Vector3

	if cannotRefract || reflectance(cosTheta, ri) > utils.RandomFloat() {
		direction = vector.Reflect(unitDirection, rec.Normal)
	} else {
		direction = vector.Refract(unitDirection, rec.Normal, ri)
	}

	scattered = ray.NewRay(rec.P, direction)
	return
}

func reflectance(cosine, refractionIndex float64) float64 {
	// Use Schlick's approximation for reflectance
	r0 := (1 - refractionIndex) / (1 + refractionIndex)
	r0 *= r0
	return r0 + (1-r0)*math.Pow((1-cosine), 5)
}

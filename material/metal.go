package material

import (
	"gotracer/color"
	"gotracer/ray"
	"gotracer/vector"
)

type Metal struct {
	albedo color.Color
	fuzz   float64
}

func NewMetal(albedo color.Color, fuzz float64) Metal {
	return Metal{albedo, fuzz}
}

func (m *Metal) Scatter(in ray.Ray, rec *HitRecord) (isScatter bool, scattered ray.Ray, attenuation color.Color) {
	reflected := vector.Reflect(in.Direction(), rec.Normal)
	reflected = reflected.UnitVector().AddVector(vector.RandomUnitVector().MultiplyFloat(m.fuzz))
	scattered = ray.NewRay(rec.P, reflected)
	attenuation = m.albedo
	isScatter = scattered.Direction().Dot(rec.Normal) > 0
	return
}

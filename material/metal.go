package material

import (
	"gotracer/color"
	"gotracer/ray"
	"gotracer/vector"
)

type Metal struct {
	albedo color.Color
}

func NewMetal(albedo color.Color) Metal {
	return Metal{albedo}
}

func (m *Metal) Scatter(in ray.Ray, rec *HitRecord) (isScatter bool, scattered ray.Ray, attenuation color.Color) {
	isScatter = true
	reflected := vector.Reflect(in.Direction(), rec.Normal)
	scattered = ray.NewRay(rec.P, reflected)
	attenuation = m.albedo
	return
}

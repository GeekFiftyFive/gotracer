package material

import (
	"gotracer/color"
	"gotracer/ray"
	"gotracer/vector"
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
	refracted := vector.Refract(unitDirection, rec.Normal, ri)
	scattered = ray.NewRay(rec.P, refracted)
	return
}

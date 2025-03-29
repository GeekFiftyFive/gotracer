package material

import (
	"gotracer/color"
	"gotracer/ray"
)

type Material interface {
	Scatter(in ray.Ray, rec *HitRecord) (isScatter bool, scattered ray.Ray, attenuation color.Color)
}

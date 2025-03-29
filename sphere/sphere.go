package sphere

import (
	"gotracer/hittable"
	"gotracer/interval"
	"gotracer/ray"
	"gotracer/vector"
	"math"
)

type Sphere struct {
	Center vector.Point3
	Radius float64
}

func NewSphere(center vector.Point3, radius float64) Sphere {
	return Sphere{Center: center, Radius: math.Max(0, radius)}
}

func (s *Sphere) Hit(r ray.Ray, rayT interval.Interval) (isHit bool, rec *hittable.HitRecord) {
	oc := s.Center.SubtractVector(r.Origin())
	a := r.Direction().LengthSquared()
	h := r.Direction().Dot(oc)
	c := oc.LengthSquared() - s.Radius*s.Radius

	discriminant := h*h - a*c
	if discriminant < 0 {
		return
	}

	sqrtd := math.Sqrt(discriminant)
	root := (h - sqrtd) / a
	if !rayT.Surrounds(root) {
		root = (h + sqrtd) / a
		if !rayT.Surrounds(root) {
			return
		}
	}

	isHit = true
	rec = &hittable.HitRecord{}
	rec.T = root
	rec.P = r.At(rec.T)
	outwardNormal := rec.P.SubtractVector(s.Center).DivideFloat(s.Radius)
	rec.SetFaceNormal(r, outwardNormal)
	return
}

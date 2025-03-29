package sphere

import (
	"gotracer/interval"
	"gotracer/material"
	"gotracer/ray"
	"gotracer/vector"
	"math"
)

type Sphere struct {
	center vector.Point3
	radius float64
	mat    material.Material
}

func NewSphere(center vector.Point3, radius float64, mat material.Material) Sphere {
	return Sphere{center: center, radius: math.Max(0, radius), mat: mat}
}

func (s *Sphere) Hit(r ray.Ray, rayT interval.Interval) (isHit bool, rec *material.HitRecord) {
	oc := s.center.SubtractVector(r.Origin())
	a := r.Direction().LengthSquared()
	h := r.Direction().Dot(oc)
	c := oc.LengthSquared() - s.radius*s.radius

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
	rec = &material.HitRecord{}
	rec.T = root
	rec.P = r.At(rec.T)
	outwardNormal := rec.P.SubtractVector(s.center).DivideFloat(s.radius)
	rec.SetFaceNormal(r, outwardNormal)
	rec.Mat = s.mat
	return
}

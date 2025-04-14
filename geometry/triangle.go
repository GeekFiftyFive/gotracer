package geometry

import (
	"gotracer/interval"
	"gotracer/material"
	"gotracer/ray"
	"gotracer/vector"
	"math"
)

type Triangle struct {
	v1  vector.Point3
	v2  vector.Point3
	v3  vector.Point3
	mat material.Material
}

func NewTriangle(v1, v2, v3 vector.Point3, mat material.Material) Triangle {
	return Triangle{v1, v2, v3, mat}
}

func (tri *Triangle) Hit(r ray.Ray, rayT interval.Interval) (isHit bool, rec *material.HitRecord) {
	epsilon := 0.001

	direction := r.Direction()
	edge1 := tri.v2.SubtractVector(tri.v1)
	edge2 := tri.v3.SubtractVector(tri.v1)
	edge3 := tri.v3.SubtractVector(tri.v2)
	rayCrossE2 := direction.Cross(edge2)
	det := edge1.Dot(rayCrossE2)

	if det > -epsilon && det < epsilon {
		return
	}

	invDet := 1.0 / det
	s := r.Origin().SubtractVector(tri.v1)
	u := s.Dot(rayCrossE2) * invDet

	if (u < 0 && math.Abs(u) > epsilon) || (u > 1 && math.Abs(u-1) > epsilon) {
		return
	}

	sCrossE1 := s.Cross(edge1)
	v := invDet * direction.Dot(sCrossE1)

	if (v < 0 && math.Abs(v) > epsilon) || (u+v > 1 && math.Abs(u+v-1) > epsilon) {
		return
	}

	t := invDet * edge2.Dot(sCrossE1)

	if t <= epsilon || t > rayT.Max || t < rayT.Min {
		return
	}

	normal := edge1.Cross(edge3).UnitVector()

	isHit = true
	rec = &material.HitRecord{}
	rec.T = t
	rec.P = r.At(rec.T)
	rec.SetNormal(normal)
	rec.Mat = tri.mat
	return
}

func (tri *Triangle) SetWindingSign(sign int) {
	ab := tri.v1.SubtractVector(tri.v2)
	ac := tri.v1.SubtractVector(tri.v3)
	cross := ab.Cross(ac)

	if sign < 0 && cross.Z() > 0 || sign > 0 && cross.Z() < 0 {
		swap := tri.v1
		tri.v1 = tri.v3
		tri.v3 = swap
	}
}

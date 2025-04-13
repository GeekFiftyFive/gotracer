package geometry

import (
	"gotracer/interval"
	"gotracer/material"
	"gotracer/matrix"
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
	normal := tri.v1.SubtractVector(tri.v2).Cross(tri.v3.SubtractVector(tri.v1))
	dot := r.Direction().Dot(normal)
	if math.Abs(dot) < 0.001 { // Account for floating point error
		// Ray and Plane tri sits on are parallel, no intersection
		return
	}

	matA := matrix.NewMatrix3(r.Direction().MultiplyFloat(-1), tri.v2.SubtractVector(tri.v1), tri.v3.SubtractVector(tri.v1))
	detA := matA.Determinant()

	u := matrix.NewMatrix3(r.Direction().MultiplyFloat(-1), r.Origin().SubtractVector(tri.v1), tri.v3.SubtractVector(tri.v1)).Determinant() / detA
	v := matrix.NewMatrix3(r.Direction().MultiplyFloat(-1), tri.v2.SubtractVector(tri.v1), r.Origin().SubtractVector(tri.v1)).Determinant() / detA

	if (v < 0 && math.Abs(v) > 0.001) || (u+v > 1 && math.Abs(u+v-1) > 0.001) {
		return
	}

	t := matrix.NewMatrix3(r.Origin().SubtractVector(tri.v1), tri.v2.SubtractVector(tri.v1), tri.v3.SubtractVector(tri.v1)).Determinant() / detA

	if t < 0.001 {
		return
	}
	isHit = true
	rec = &material.HitRecord{}
	rec.T = t
	rec.P = r.At(t)
	rec.SetFaceNormal(r, normal)
	rec.Mat = tri.mat
	return
}

func pointInTriangle(point vector.Point3, triangle Triangle) bool {
	p := point
	a := triangle.v1
	b := triangle.v2
	c := triangle.v3

	a = a.SubtractVector(p)
	b = b.SubtractVector(p)
	c = c.SubtractVector(p)

	u := b.Cross(c)
	v := c.Cross(a)
	w := a.Cross(b)

	if u.Dot(v) < 0 || u.Dot(w) < 0 {
		return false
	}

	return true
}

package geometry

import (
	"gotracer/interval"
	"gotracer/material"
	"gotracer/ray"
	"gotracer/vector"
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

func (t *Triangle) Hit(r ray.Ray, rayT interval.Interval) (isHit bool, rec *material.HitRecord) {
	// TODO: Implement
	return
}

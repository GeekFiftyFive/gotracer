package ray

import "gotracer/vector"

type ray struct {
	orig vector.Point3
	dir  vector.Vector3
}

type Ray interface {
	Origin() vector.Point3
	Direction() vector.Vector3
	At(t float64) vector.Point3
}

func NewRay(origin vector.Point3, direction vector.Vector3) Ray {
	return ray{orig: origin, dir: direction}
}

func (r ray) Origin() vector.Point3 {
	return r.orig
}

func (r ray) Direction() vector.Vector3 {
	return r.dir
}

func (r ray) At(t float64) vector.Point3 {
	return r.orig.AddVector(r.Direction().MultiplyFloat(t))
}

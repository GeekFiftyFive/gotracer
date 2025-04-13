package material

import (
	"gotracer/interval"
	"gotracer/ray"
	"gotracer/vector"
)

type HitRecord struct {
	P         vector.Point3
	Normal    vector.Vector3
	Mat       Material
	T         float64
	FrontFace bool
}

type Hittable = interface {
	Hit(r ray.Ray, rayT interval.Interval) (isHit bool, rec *HitRecord)
}

func (rec *HitRecord) SetFaceNormal(r ray.Ray, outwardNormal vector.Vector3) {
	// Sets the hit record normal vector.
	// NOTE: the parameter `outwardNormal` is assumed to have unit length.

	rec.FrontFace = r.Direction().Dot(outwardNormal) < 0.0
	if rec.FrontFace {
		rec.Normal = outwardNormal
	} else {
		rec.Normal = outwardNormal.MultiplyFloat(-1)
	}
}

type HittableList struct {
	objects []Hittable
}

func (hl *HittableList) Clear() {
	hl.objects = []Hittable{}
}

func (hl *HittableList) Add(h Hittable) {
	hl.objects = append(hl.objects, h)
}

func (hl *HittableList) Hit(r ray.Ray, rayT interval.Interval) (hitAnything bool, rec *HitRecord) {
	closestSoFar := rayT.Max
	for _, obj := range hl.objects {
		isHit, tempRec := obj.Hit(r, interval.Interval{Min: rayT.Min, Max: closestSoFar})
		if isHit {
			hitAnything = true
			closestSoFar = tempRec.T
			rec = tempRec
		}
	}
	return
}

package vector

import (
	"gotracer/utils"
	"math"
)

type vec3 struct {
	x float64
	y float64
	z float64
}

type Vector3 interface {
	X() float64
	Y() float64
	Z() float64
	AddVector(Vector3) Vector3
	SubtractVector(Vector3) Vector3
	MultiplyVector(Vector3) Vector3
	DivideVector(Vector3) Vector3
	AddFloat(float64) Vector3
	SubtractFloat(float64) Vector3
	MultiplyFloat(float64) Vector3
	DivideFloat(float64) Vector3
	Length() float64
	LengthSquared() float64
	Dot(Vector3) float64
	Cross(Vector3) Vector3
	UnitVector() Vector3
	NearZero() bool
}

type Point3 = Vector3

func NewVector3(x, y, z float64) Vector3 {
	return vec3{x, y, z}
}

func Random() Vector3 {
	return vec3{x: utils.RandomFloat(), y: utils.RandomFloat(), z: utils.RandomFloat()}
}

func RandomRange(min float64, max float64) Vector3 {
	return vec3{x: utils.RandomRange(min, max), y: utils.RandomRange(min, max), z: utils.RandomRange(min, max)}
}

func RandomUnitVector() Vector3 {
	for {
		p := RandomRange(-1, 1)
		lensq := p.LengthSquared()
		if 1e-160 < lensq && lensq <= 1 {
			return p.DivideFloat(math.Sqrt(lensq))
		}
	}
}

func RandomOnHemisphere(normal Vector3) Vector3 {
	onUnitSphere := RandomUnitVector()
	if onUnitSphere.Dot(normal) > 0.0 { // In the same hemisphere as the normal
		return onUnitSphere
	} else {
		return onUnitSphere.MultiplyFloat(-1)
	}
}

func Reflect(v Vector3, n Vector3) Vector3 {
	return v.SubtractVector(n.MultiplyFloat(v.Dot(n) * 2))
}

func (v vec3) X() float64 {
	return v.x
}

func (v vec3) Y() float64 {
	return v.y
}

func (v vec3) Z() float64 {
	return v.z
}

func (v1 vec3) AddVector(v2 Vector3) Vector3 {
	return vec3{x: v1.x + v2.X(), y: v1.y + v2.Y(), z: v1.z + v2.Z()}
}

func (v1 vec3) SubtractVector(v2 Vector3) Vector3 {
	return vec3{x: v1.x - v2.X(), y: v1.y - v2.Y(), z: v1.z - v2.Z()}
}

func (v1 vec3) MultiplyVector(v2 Vector3) Vector3 {
	return vec3{x: v1.x * v2.X(), y: v1.y * v2.Y(), z: v1.z * v2.Z()}
}

func (v1 vec3) DivideVector(v2 Vector3) Vector3 {
	return vec3{x: v1.x / v2.X(), y: v1.y / v2.Y(), z: v1.z / v2.Z()}
}

func (v vec3) AddFloat(f float64) Vector3 {
	return vec3{x: v.x + f, y: v.y + f, z: v.z + f}
}

func (v vec3) SubtractFloat(f float64) Vector3 {
	return vec3{x: v.x - f, y: v.y - f, z: v.z - f}
}

func (v vec3) MultiplyFloat(f float64) Vector3 {
	return vec3{x: v.x * f, y: v.y * f, z: v.z * f}
}

func (v vec3) DivideFloat(f float64) Vector3 {
	return vec3{x: v.x / f, y: v.y / f, z: v.z / f}
}

func (v vec3) LengthSquared() float64 {
	return v.x*v.x + v.y*v.y + v.z*v.z
}

func (v vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v1 vec3) Dot(v2 Vector3) float64 {
	return v1.X()*v2.X() + v1.Y()*v2.Y() + v1.Z()*v2.Z()
}

func (v1 vec3) Cross(v2 Vector3) Vector3 {
	return vec3{x: v1.Y()*v2.Z() - v1.Z()*v2.Y(), y: v1.Z()*v2.X() - v1.X()*v2.Z(), z: v1.X()*v2.Y() - v1.Y()*v2.X()}
}

func (v vec3) UnitVector() Vector3 {
	return v.DivideFloat(v.Length())
}

func (v vec3) NearZero() bool {
	// Return true if vector is close to 0 in all directions
	s := 1e-8
	return math.Abs(v.x) < s && math.Abs(v.y) < s && math.Abs(v.z) < s
}

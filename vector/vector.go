package vector

import "math"

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
}

func NewVector3(x, y, z float64) Vector3 {
	return vec3{x, y, z}
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
	return v.x*v.x + v.y*v.y + v.z + v.z
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

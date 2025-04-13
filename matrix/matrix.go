package matrix

import "gotracer/vector"

type matrix2 struct {
	a float64
	b float64
	c float64
	d float64
}

func (m matrix2) determinant() float64 {
	return m.a*m.d - m.b*m.c
}

type matrix3 struct {
	a1 vector.Vector3
	a2 vector.Vector3
	a3 vector.Vector3
}

type Matrix3 interface {
	A1() vector.Vector3
	A2() vector.Vector3
	A3() vector.Vector3
	Determinant() float64
}

func NewMatrix3(a1 vector.Vector3, a2 vector.Vector3, a3 vector.Vector3) Matrix3 {
	return matrix3{a1, a2, a3}
}

func (m matrix3) A1() vector.Vector3 {
	return m.a1
}

func (m matrix3) A2() vector.Vector3 {
	return m.a2
}

func (m matrix3) A3() vector.Vector3 {
	return m.a3
}

func (m matrix3) Determinant() float64 {
	aDet := m.a1.X() * matrix2{a: m.a2.Y(), b: m.a3.Y(), c: m.a2.Z(), d: m.a3.Z()}.determinant()
	bDet := m.a2.X() * matrix2{a: m.a1.Y(), b: m.a3.Y(), c: m.a1.Z(), d: m.a3.Z()}.determinant()
	cDet := m.a3.X() * matrix2{a: m.a1.Y(), b: m.a2.Y(), c: m.a1.Z(), d: m.a2.Z()}.determinant()
	return aDet - bDet + cDet
}

package geometry

import (
	"gotracer/material"
	"gotracer/vector"
)

type Cuboid = material.HittableList

// func NewCuboid(v1, v2 vector.Point3, mat material.Material) Cuboid {
// 	cuboid := Cuboid{}
// 	points := []vector.Point3{}

// 	for i := range 8 {
// 		var x, y, z float64
// 		if i&1 > 0 {
// 			x = v2.X()
// 		} else {
// 			x = v1.X()
// 		}

// 		if i&2 > 0 {
// 			y = v2.Y()
// 		} else {
// 			y = v1.Y()
// 		}

// 		if i&4 > 0 {
// 			z = v2.Z()
// 		} else {
// 			z = v1.Z()
// 		}

// 		points = append(points, vector.NewVector3(x, y, z))
// 	}

// 	logger := log.New(os.Stderr, "", 0) // TODO: Move logger into its own package

// 	logger.Print(points)

// 	for idx := range 3 {
// 		u1 := 1 << idx
// 		var u2 int
// 		if u1 == 4 {
// 			u2 = 1
// 		} else {
// 			u2 = u1 << 1
// 		}

// 		addVert := func(i, j, k int) {
// 			vert := NewTriangle(points[i], points[j], points[k], mat)
// 			cuboid.Add(&vert)
// 		}
// 		addVert(u2, u1, 0)
// 		addVert(u1, u2, u1+u2)
// 		addVert(7-u1, 7-u2, 7)
// 		addVert(7-u2, 7-u1, 7-(u1+u2))
// 	}
// 	return cuboid
// }

func NewCuboid(v1, v2 vector.Point3, mat material.Material) Cuboid {
	vertices := []vector.Vector3{
		vector.NewVector3(v1.X(), v1.Y(), v1.Z()),
		vector.NewVector3(v2.X(), v1.Y(), v1.Z()),
		vector.NewVector3(v2.X(), v2.Y(), v1.Z()),
		vector.NewVector3(v1.X(), v2.Y(), v1.Z()),
		vector.NewVector3(v1.X(), v2.Y(), v2.Z()),
		vector.NewVector3(v2.X(), v2.Y(), v2.Z()),
		vector.NewVector3(v2.X(), v1.Y(), v2.Z()),
		vector.NewVector3(v1.X(), v1.Y(), v2.Z()),
	}
	getTri := func(i, j, k int) Triangle {
		return NewTriangle(vertices[i], vertices[j], vertices[k], mat)
	}

	triangles := []Triangle{
		getTri(0, 2, 1),
		getTri(0, 3, 2),
		getTri(2, 3, 4),
		getTri(2, 4, 5),
		getTri(1, 2, 5),
		getTri(1, 5, 6),
		getTri(0, 7, 4),
		getTri(0, 4, 3),
		getTri(5, 4, 7),
		getTri(5, 7, 6),
		getTri(0, 6, 7),
		getTri(0, 1, 6),
	}

	cuboid := Cuboid{}

	for _, triangle := range triangles {
		cuboid.Add(&triangle)
	}

	return cuboid
}

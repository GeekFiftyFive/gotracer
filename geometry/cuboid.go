package geometry

import (
	"fmt"
	"gotracer/material"
	"gotracer/vector"
	"log"
	"os"
)

type Cuboid = material.HittableList

func NewCuboid(v1, v2 vector.Point3, mat material.Material) Cuboid {
	logger := log.New(os.Stderr, "", 0) // TODO: Move logger into its own package
	cuboid := Cuboid{}
	points := []vector.Point3{}
	faces := make(map[string][]vector.Point3)

	for i := range 8 {
		var x, y, z float64
		if i&1 > 0 {
			x = v2.X()
		} else {
			x = v1.X()
		}

		if i&2 > 0 {
			y = v2.Y()
		} else {
			y = v1.Y()
		}

		if i&4 > 0 {
			z = v2.Z()
		} else {
			z = v1.Z()
		}

		points = append(points, vector.NewVector3(x, y, z))
	}

	for _, point := range points {
		faceKeys := getFaceKeys(v1, v2, point)
		for _, faceKey := range faceKeys {
			face := faces[faceKey]
			if face == nil {
				face = []vector.Point3{}
			}
			face = append(face, point)
			faces[faceKey] = face
			// This is a complete face, calculate the tris
			if len(face) == 4 {
				tri1 := NewTriangle(face[0], face[1], face[2], mat)
				tri2 := NewTriangle(face[1], face[3], face[2], mat)
				logger.Print(tri1, tri2)
				cuboid.Add(&tri1)
				cuboid.Add(&tri2)
			}
		}
	}

	return cuboid
}

func getFaceKeys(v1, v2, u vector.Point3) (faceKeys []string) {
	for i, point := range []vector.Point3{v1, v2} {
		if u.X() == point.X() {
			faceKeys = append(faceKeys, fmt.Sprintf("x%d", i+1))
		}

		if u.Y() == point.Y() {
			faceKeys = append(faceKeys, fmt.Sprintf("y%d", i+1))
		}

		if u.Z() == point.Z() {
			faceKeys = append(faceKeys, fmt.Sprintf("z%d", i+1))
		}
	}

	return
}

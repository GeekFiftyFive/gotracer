package color

import (
	"bytes"
	"fmt"
	"gotracer/vector"
)

type Color = vector.Vector3

const multiplier = 255.999

func NewColor(r float64, g float64, b float64) Color {
	return vector.NewVector3(r, g, b)
}

func WriteColor(color Color, buffer *bytes.Buffer) {
	r := color.X()
	g := color.Y()
	b := color.Z()

	rbyte := int(multiplier * r)
	gbyte := int(multiplier * g)
	bbyte := int(multiplier * b)

	buffer.Write(fmt.Appendf(nil, "%d %d %d\n", rbyte, gbyte, bbyte))
}

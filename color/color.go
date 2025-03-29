package color

import (
	"bytes"
	"fmt"
	"gotracer/interval"
	"gotracer/vector"
	"math"
)

type Color = vector.Vector3

const multiplier = 256

func NewColor(r float64, g float64, b float64) Color {
	return vector.NewVector3(r, g, b)
}

func linearToGamma(linearComponent float64) float64 {
	if linearComponent > 0 {
		return math.Sqrt(linearComponent)
	}

	return 0
}

func WriteColor(color Color, buffer *bytes.Buffer) {
	r := linearToGamma(color.X())
	g := linearToGamma(color.Y())
	b := linearToGamma(color.Z())

	intensity := interval.Interval{Min: 0.0, Max: 0.999}
	rbyte := int(multiplier * intensity.Clamp(r))
	gbyte := int(multiplier * intensity.Clamp(g))
	bbyte := int(multiplier * intensity.Clamp(b))

	buffer.Write(fmt.Appendf(nil, "%d %d %d\n", rbyte, gbyte, bbyte))
}

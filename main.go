package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Image
	imageWidth, imageHeight := 256, 256

	//Render
	var writer bytes.Buffer
	writer.Write(fmt.Appendf(nil, "P3\n%d %d\n255\n", imageWidth, imageHeight))

	for j := range imageHeight {
		for i := range imageWidth {
			r := float64(i) / (float64(imageWidth) - 1)
			g := float64(j) / (float64(imageHeight) - 1)
			b := float64(0)

			const multiplier float64 = 255.999

			ir := int(multiplier * r)
			ig := int(multiplier * g)
			ib := int(multiplier * b)

			writer.Write(fmt.Appendf(nil, "%d %d %d\n", ir, ig, ib))
		}
	}

	fmt.Print(writer.String())
}

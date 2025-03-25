package main

import (
	"bytes"
	"fmt"
	"gotracer/color"
	"log"
	"os"
)

func main() {
	// Logger
	logger := log.New(os.Stderr, "", 0)

	// Image
	imageWidth, imageHeight := 256, 256

	// Render
	var writer bytes.Buffer
	writer.Write(fmt.Appendf(nil, "P3\n%d %d\n255\n", imageWidth, imageHeight))

	for j := range imageHeight {
		logger.Printf("Scanlines remaining: %d\n", imageHeight-j)
		for i := range imageWidth {
			r := float64(i) / (float64(imageWidth) - 1)
			g := float64(j) / (float64(imageHeight) - 1)
			b := float64(0)

			pixelColor := color.NewColor(r, g, b)

			color.WriteColor(pixelColor, &writer)
		}
	}

	fmt.Print(writer.String())
}

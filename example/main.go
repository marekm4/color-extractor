package main

import (
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	"io"
	"os"

	"github.com/marekm4/color-extractor"
)

func main() {
	image, _, _ := image.Decode(os.Stdin)
	colors := color_extractor.ExtractColors(image)
	createPalette(os.Stdout, colors)
}

func createPalette(w io.Writer, colors []color.Color) {
	squareSize := 40
	image := image.NewRGBA(image.Rect(0, 0, len(colors)*squareSize, squareSize))
	for i, color := range colors {
		for j := 0; j < squareSize; j++ {
			for k := 0; k < squareSize; k++ {
				image.Set(i*squareSize+j, k, color)
			}
		}
	}
	png.Encode(w, image)
}

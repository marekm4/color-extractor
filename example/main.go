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
	img, _, _ := image.Decode(os.Stdin)
	colors := color_extractor.ExtractColors(img)

	createPalette(os.Stdout, colors)
}

func createPalette(w io.Writer, colors []color.Color) {
	box := 40
	img := image.NewRGBA(image.Rect(0, 0, len(colors)*box, box))
	for i, color := range colors {
		for j := 0; j < box; j++ {
			for k := 0; k < box; k++ {
				img.Set(i*box+j, k, color)
			}
		}
	}
	png.Encode(w, img)
}

package main

import (
	"image"
	"image/color"
	_ "image/jpeg"
	"os"
	"testing"

	"github.com/marekm4/color-extractor"
)

func TestMain(t *testing.T) {
	file := "Fotolia_45549559_320_480.jpg"
	imageFile, _ := os.Open(file)
	defer imageFile.Close()

	image, _, _ := image.Decode(imageFile)
	colors := color_extractor.ExtractColors(image)

	expectedColors := []color.Color{
		color.RGBA{231, 230, 227, 255},
		color.RGBA{57, 58, 10, 255},
		color.RGBA{204, 51, 24, 255},
		color.RGBA{190, 177, 55, 255},
		color.RGBA{104, 152, 11, 255},
	}
	if !color_extractor.TestColorsEqual(expectedColors, colors) {
		t.Fatalf("Image %s: %v expected, got %v", file, expectedColors, colors)
	}
}

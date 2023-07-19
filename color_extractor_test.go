package color_extractor

import (
	"image"
	"image/color"
	_ "image/jpeg"
	"os"
	"testing"
)

func TestExtractColors(t *testing.T) {
	white := color.RGBA{R: 225, G: 255, B: 255, A: 255}
	red := color.RGBA{R: 255, G: 0, B: 0, A: 255}
	green := color.RGBA{R: 0, G: 255, B: 0, A: 255}
	transparent := color.RGBA{R: 0, G: 0, B: 0, A: 0}
	semiTransparentRed := color.RGBA{R: 255, G: 0, B: 0, A: 127}

	testCases := map[string]struct {
		Image           image.Image
		ExtractedColors []color.Color
	}{
		"Empty file": {
			Image:           imageFromColors([]color.Color{}),
			ExtractedColors: []color.Color{},
		},
		"Single pixel": {
			Image: imageFromColors([]color.Color{
				red,
			}),
			ExtractedColors: []color.Color{
				red,
			},
		},
		"One color": {
			Image: imageFromColors([]color.Color{
				white,
				white,
				white,
				white,
			}),
			ExtractedColors: []color.Color{
				white,
			},
		},
		"Transparent image": {
			Image: imageFromColors([]color.Color{
				white,
				white,
				white,
				transparent,
			}),
			ExtractedColors: []color.Color{
				white,
			},
		},
		"Semitransparent single pixel": {
			Image: imageFromColors([]color.Color{
				semiTransparentRed,
			}),
			ExtractedColors: []color.Color{
				red,
			},
		},
		"Semitransparent image": {
			Image: imageFromColors([]color.Color{
				semiTransparentRed,
				semiTransparentRed,
				green,
			}),
			ExtractedColors: []color.Color{
				green,
				red,
			},
		},
		"Semitransparent image, bigger semitransparent region": {
			Image: imageFromColors([]color.Color{
				semiTransparentRed,
				semiTransparentRed,
				semiTransparentRed,
				green,
			}),
			ExtractedColors: []color.Color{
				red,
				green,
			},
		},
		"Two colors": {
			Image: imageFromColors([]color.Color{
				red,
				red,
				green,
				green,
				red,
				red,
			}),
			ExtractedColors: []color.Color{
				red,
				green,
			},
		},
		"Mixed colors": {
			Image: imageFromColors([]color.Color{
				red,
				red,
				color.RGBA{R: 245, G: 0, B: 0, A: 255},
				color.RGBA{R: 245, G: 0, B: 0, A: 255},
				green,
				green,
				color.RGBA{R: 0, G: 240, B: 0, A: 255},
			}),
			ExtractedColors: []color.Color{
				color.RGBA{R: 250, G: 0, B: 0, A: 255},
				color.RGBA{R: 0, G: 250, B: 0, A: 255},
			},
		},
		"File": {
			Image: imageFromFile("example/Fotolia_45549559_320_480.jpg"),
			ExtractedColors: []color.Color{
				color.RGBA{R: 232, G: 230, B: 228, A: 255},
				color.RGBA{R: 58, G: 58, B: 10, A: 255},
				color.RGBA{R: 205, G: 51, B: 25, A: 255},
				color.RGBA{R: 191, G: 178, B: 56, A: 255},
				color.RGBA{R: 104, G: 152, B: 12, A: 255},
			},
		},
		"Sub Image": {
			Image: subImage(imageFromFile("example/Fotolia_45549559_320_480.jpg"), image.Rect(25, 120, 35, 130)),
			ExtractedColors: []color.Color{
				color.RGBA{R: 154, G: 1, B: 8, A: 255},
				color.RGBA{R: 117, G: 2, B: 1, A: 255},
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			extractedColors := ExtractColors(testCase.Image)
			if !testColorsEqual(testCase.ExtractedColors, extractedColors) {
				t.Fatalf("%v expected, got %v", testCase.ExtractedColors, extractedColors)
			}
		})
	}
}

func imageFromColors(colors []color.Color) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, len(colors), 1))
	for i, c := range colors {
		img.Set(i, 0, c)
	}
	return img
}

func imageFromFile(filename string) image.Image {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer func() {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}()
	img, _, _ := image.Decode(file)
	return img
}

func subImage(img image.Image, rect image.Rectangle) image.Image {
	return img.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(rect)
}

func testColorsEqual(a, b []color.Color) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

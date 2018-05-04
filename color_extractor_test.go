package color_extractor

import (
	"image"
	"image/color"
	_ "image/jpeg"
	"os"
	"testing"
)

func TestExtractColors(t *testing.T) {
	white := color.RGBA{225, 255, 255, 255}
	red := color.RGBA{255, 0, 0, 255}
	green := color.RGBA{0, 255, 0, 255}
	transparent := color.RGBA{0, 0, 0, 0}

	testCases := []struct {
		Name            string
		Image           image.Image
		ExtractedColors []color.Color
	}{
		{
			Name:            "Empty file",
			Image:           imageFromColors([]color.Color{}),
			ExtractedColors: []color.Color{},
		},
		{
			Name: "Single pixel",
			Image: imageFromColors([]color.Color{
				red,
			}),
			ExtractedColors: []color.Color{
				red,
			},
		},
		{
			Name: "One color",
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
		{
			Name: "Transparent image",
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
		{
			Name: "Two colors",
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
		{
			Name: "Mixed colors",
			Image: imageFromColors([]color.Color{
				red,
				red,
				color.RGBA{245, 0, 0, 255},
				color.RGBA{245, 0, 0, 255},
				green,
				green,
				color.RGBA{0, 240, 0, 255},
			}),
			ExtractedColors: []color.Color{
				color.RGBA{250, 0, 0, 255},
				color.RGBA{0, 250, 0, 255},
			},
		},
		{
			Name:  "File",
			Image: imageFromFile("example/Fotolia_45549559_320_480.jpg"),
			ExtractedColors: []color.Color{
				color.RGBA{231, 230, 227, 255},
				color.RGBA{57, 58, 10, 255},
				color.RGBA{204, 51, 24, 255},
				color.RGBA{190, 177, 55, 255},
				color.RGBA{104, 152, 11, 255},
			},
		},
	}

	for _, testCase := range testCases {
		extractedColors := ExtractColors(testCase.Image)
		if !testColorsEqual(testCase.ExtractedColors, extractedColors) {
			t.Fatalf("TestCase %s: %v expected, got %v", testCase.Name, testCase.ExtractedColors, extractedColors)
		}
	}
}

func imageFromColors(colors []color.Color) image.Image {
	image := image.NewRGBA(image.Rect(0, 0, len(colors), 1))
	for i, color := range colors {
		image.Set(i, 0, color)
	}
	return image
}

func imageFromFile(filename string) image.Image {
	file, _ := os.Open(filename)
	defer file.Close()
	image, _, _ := image.Decode(file)
	return image
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

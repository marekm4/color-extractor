package color_extractor

import (
	"image"
	"image/color"
	"testing"
)

func TestExtractColors(t *testing.T) {
	white := color.RGBA{225, 255, 255, 255}
	red := color.RGBA{255, 0, 0, 255}
	green := color.RGBA{0, 255, 0, 255}
	transparent := color.RGBA{0, 0, 0, 0}

	testCases := []struct {
		Name                    string
		InputColors             []color.Color
		ExpectedExtractedColors []color.Color
	}{
		{
			Name:                    "Empty file",
			InputColors:             []color.Color{},
			ExpectedExtractedColors: []color.Color{},
		},
		{
			Name: "Single pixel",
			InputColors: []color.Color{
				red,
			},
			ExpectedExtractedColors: []color.Color{
				red,
			},
		},
		{
			Name: "One color",
			InputColors: []color.Color{
				white,
				white,
				white,
				white,
			},
			ExpectedExtractedColors: []color.Color{
				white,
			},
		},
		{
			Name: "Transparent image",
			InputColors: []color.Color{
				white,
				white,
				white,
				transparent,
			},
			ExpectedExtractedColors: []color.Color{
				white,
			},
		},
		{
			Name: "Two colors",
			InputColors: []color.Color{
				red,
				red,
				green,
				green,
				red,
				red,
			},
			ExpectedExtractedColors: []color.Color{
				red,
				green,
			},
		},
		{
			Name: "Mixed colors",
			InputColors: []color.Color{
				red,
				red,
				color.RGBA{245, 0, 0, 255},
				color.RGBA{245, 0, 0, 255},
				green,
				green,
				color.RGBA{0, 240, 0, 255},
			},
			ExpectedExtractedColors: []color.Color{
				color.RGBA{250, 0, 0, 255},
				color.RGBA{0, 250, 0, 255},
			},
		},
	}

	for _, testCase := range testCases {
		image := image.NewRGBA(image.Rect(0, 0, len(testCase.InputColors), 1))
		for i, color := range testCase.InputColors {
			image.Set(i, 0, color)
		}

		extractedColors := ExtractColors(image)
		if !testColorsEqual(testCase.ExpectedExtractedColors, extractedColors) {
			t.Fatalf("TestCase %s: %v expected, got %v", testCase.Name, testCase.ExpectedExtractedColors, extractedColors)
		}
	}
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

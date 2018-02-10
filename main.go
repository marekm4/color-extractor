package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"os"

	"github.com/bugra/kmeans"
)

func main() {
	files := []string{
		"example_images/01540_driftwood_1366x768.jpg",
		"example_images/04103_misslibertyii_1920x1080.jpg",
		"example_images/14775211410_42b8d244da_o.jpg",
		"example_images/Fotolia_45549559_320_480.jpg",
		"example_images/giant-panda-shutterstock_86500690.jpg",
		"example_images/mac_os_x_retina_zebras-wallpaper-1920x1080.jpg",
	}

	fmt.Println("<div>")
	for _, file := range files {
		imageFile, _ := os.Open(file)
		defer imageFile.Close()

		image, _, _ := image.Decode(imageFile)
		colors := extractColors(image, 5)

		fmt.Println("<img src=\"" + file + "\" width=\"200\"><br>")
		for _, c := range colors {
			printColor(c)

		}
		fmt.Println("<br><br><br>")
	}
	fmt.Println("</div>")
}

func printColor(c color.Color) {
	r, g, b, _ := c.RGBA()
	fmt.Print("<div style=\"background-color:rgb(", r>>8, ",", g>>8, ",", b>>8, ");display:inline-block;width:40px;height:40px;margin-right:-5px;\"></div>\n")
}

func extractColors(image image.Image, colorsCount int) []color.Color {
	width := image.Bounds().Max.X
	height := image.Bounds().Max.Y

	// calculate downsizing ratio
	step := width/256 + 1

	// load image's pixels into [][]float64
	colorData := [][]float64{}
	for x := 0; x < width; x += step {
		for y := 0; y < height; y += step {
			color := image.At(x, y)
			r, g, b, _ := color.RGBA()
			colorData = append(colorData, []float64{float64(r >> 8), float64(g >> 8), float64(b >> 8)})
		}
	}

	// calculate clusters
	clusters, _ := kmeans.Kmeans(colorData, colorsCount, kmeans.EuclideanDistance, 1)

	// calculate average color for each cluster
	selectedColorsSums := make([][]float64, colorsCount, colorsCount)
	for i := range selectedColorsSums {
		selectedColorsSums[i] = make([]float64, 3, 3)
	}
	selectedColorsCount := make([]float64, colorsCount, colorsCount)
	for idx, cluster := range clusters {
		selectedColorsCount[cluster]++
		selectedColorsSums[cluster][0] += colorData[idx][0]
		selectedColorsSums[cluster][1] += colorData[idx][1]
		selectedColorsSums[cluster][2] += colorData[idx][2]
	}

	// pack average cluster color to color.Color struct
	selectedColors := []color.Color{}
	for i := 0; i < colorsCount; i++ {
		selectedColors = append(selectedColors, color.RGBA{
			R: uint8(selectedColorsSums[i][0] / float64(selectedColorsCount[i])),
			G: uint8(selectedColorsSums[i][1] / float64(selectedColorsCount[i])),
			B: uint8(selectedColorsSums[i][2] / float64(selectedColorsCount[i])),
			A: 255,
		})
	}

	return selectedColors
}

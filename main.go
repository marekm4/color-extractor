package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"math"
	"os"
	"path/filepath"
	"sort"

	"github.com/bugra/kmeans"
)

func main() {
	files, err := filepath.Glob("example_images/*")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		imageFile, _ := os.Open(file)
		defer imageFile.Close()

		image, _, _ := image.Decode(imageFile)
		colors := extractColors(image)

		fmt.Println("<img src=\"" + file + "\" width=\"200\"><br>")
		for _, c := range colors {
			printColor(c)

		}
		fmt.Println("<br><br><br>")
	}
}

func printColor(c color.Color) {
	r, g, b, _ := c.RGBA()
	fmt.Print("<div style=\"background-color:rgb(", r>>8, ",", g>>8, ",", b>>8, ");display:inline-block;width:40px;height:40px;margin-right:-5px;\"></div>\n")
}

// https://en.wikipedia.org/wiki/Elbow_method_(clustering)
func extractColors(image image.Image) []color.Color {
	for i := 1; true; i++ {
		colors, SSE := extractColorsWithCount(image, i)
		if SSE < 2000 || i >= 7 {
			return colors
		}
	}

	return nil
}

// https://en.wikipedia.org/wiki/K-means_clustering
func extractColorsWithCount(image image.Image, colorsCount int) ([]color.Color, float64) {
	width := image.Bounds().Max.X
	height := image.Bounds().Max.Y

	// calculate downsizing ratio
	stepX := int(math.Max(float64(width)/224., 1))
	stepY := int(math.Max(float64(height)/224., 1))

	// load image's pixels into [][]float64
	colorData := [][]float64{}
	for x := 0; x < width; x += stepX {
		for y := 0; y < height; y += stepY {
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
	selectedColorsAverages := make([][]float64, colorsCount, colorsCount)
	for i := range selectedColorsAverages {
		selectedColorsAverages[i] = make([]float64, 3, 3)
	}
	for i := 0; i < colorsCount; i++ {
		selectedColorsAverages[i][0] = selectedColorsSums[i][0] / float64(selectedColorsCount[i])
		selectedColorsAverages[i][1] = selectedColorsSums[i][1] / float64(selectedColorsCount[i])
		selectedColorsAverages[i][2] = selectedColorsSums[i][2] / float64(selectedColorsCount[i])
	}

	// pack average cluster color to SortableColor struct
	selectedColors := []SortableColor{}
	for i := 0; i < colorsCount; i++ {
		selectedColors = append(selectedColors, SortableColor{
			selectedColorsCount[i],
			color.RGBA{
				R: uint8(selectedColorsAverages[i][0]),
				G: uint8(selectedColorsAverages[i][1]),
				B: uint8(selectedColorsAverages[i][2]),
				A: 255,
			},
		})
	}

	// sort colors by cluster size
	sort.Sort(sort.Reverse(ByCount(selectedColors)))

	// calculate SSE
	SSE := 0.
	for i, point := range colorData {
		cluster := clusters[i]
		centroid := selectedColorsAverages[cluster]
		change, _ := kmeans.SquaredEuclideanDistance(centroid, point)
		SSE += change
	}
	SSE /= float64(len(colorData))

	// extract color.Color from SortableColor
	selectedColorsExtracted := []color.Color{}
	for _, sc := range selectedColors {
		selectedColorsExtracted = append(selectedColorsExtracted, sc.Color)
	}

	return selectedColorsExtracted, SSE
}

type SortableColor struct {
	Count float64
	Color color.Color
}

type ByCount []SortableColor

func (c ByCount) Len() int           { return len(c) }
func (c ByCount) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c ByCount) Less(i, j int) bool { return c[i].Count < c[j].Count }

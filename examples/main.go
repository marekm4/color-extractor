package main

import (
	color_extractor "color-extractor"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
)

func main() {
	files, err := filepath.Glob("images/*")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		imageFile, _ := os.Open(file)
		defer imageFile.Close()

		image, _, _ := image.Decode(imageFile)
		colors := color_extractor.ExtractColors(image)

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

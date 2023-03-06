Simple image color extractor written in Go with no external dependencies.

[![Go Reference](https://pkg.go.dev/badge/github.com/marekm4/color-extractor.svg)](https://pkg.go.dev/github.com/marekm4/color-extractor)
[![Go Report Card](https://goreportcard.com/badge/github.com/marekm4/color-extractor)](https://goreportcard.com/report/github.com/marekm4/color-extractor)
[![Go Coverage](https://github.com/marekm4/color-extractor/wiki/coverage.svg)](https://raw.githack.com/wiki/marekm4/color-extractor/coverage.html)

Demo:

https://color-extractor-demo.marekm4.com/

Blog post:

https://medium.com/@marek.michalik/c-vs-rust-vs-go-performance-analysis-945ab749056c

Usage:
```go
package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"os"

	"github.com/marekm4/color-extractor"
)

func main() {
	file := "Fotolia_45549559_320_480.jpg"
	imageFile, _ := os.Open(file)
	defer imageFile.Close()

	image, _, _ := image.Decode(imageFile)
	colors := color_extractor.ExtractColors(image)

	fmt.Println(colors)
}
```

Example image:

![Image](https://raw.githubusercontent.com/marekm4/color-extractor/master/example/Fotolia_45549559_320_480.jpg)

Extracted colors:

![Colors](https://raw.githubusercontent.com/marekm4/color-extractor/master/example/colors.png)

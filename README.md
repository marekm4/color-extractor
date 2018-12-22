[![Build Status](https://travis-ci.org/marekm4/color-extractor.svg?branch=master)](https://travis-ci.org/marekm4/color-extractor)

Simple image color extractor written in Go with no external dependencies.

Demo:

https://color-extractor-demo.herokuapp.com/

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

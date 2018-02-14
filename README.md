Simple color extractor written in Go.

Usage:
```go
import color_extractor "github.com/marekm4/color-extractor"

...

imageFile, _ := os.Open(file)
image, _, _ := image.Decode(imageFile)
colors := color_extractor.ExtractColors(image)
```

Examples:
![Examples](https://raw.githubusercontent.com/marekm4/color-extractor/master/examples/test.png)

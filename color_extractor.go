package color_extractor

import (
	"image"
	"image/color"
	"math"
	"sort"
)

type bucket struct {
	Count int
	Color color.Color
}

type ByCount []bucket

func (c ByCount) Len() int           { return len(c) }
func (c ByCount) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c ByCount) Less(i, j int) bool { return c[i].Count < c[j].Count }

type Config struct {
	DownSizeTo  float64
	SmallBucket float64
}

func ExtractColors(image image.Image) []color.Color {
	return ExtractColorsWithConfig(image, Config{
		DownSizeTo:  224.,
		SmallBucket: .01,
	})
}

func ExtractColorsWithConfig(image image.Image, config Config) []color.Color {
	width := image.Bounds().Max.X
	height := image.Bounds().Max.Y

	// calculate downsizing ratio
	stepX := int(math.Max(float64(width)/config.DownSizeTo, 1))
	stepY := int(math.Max(float64(height)/config.DownSizeTo, 1))

	// load image's pixels into buckets
	var buckets [2][2][2][]color.Color
	var colorsCount int
	for x := 0; x < width; x += stepX {
		for y := 0; y < height; y += stepY {
			color := image.At(x, y)
			r, g, b, a := color.RGBA()
			i := r >> (8 + 7)
			j := g >> (8 + 7)
			k := b >> (8 + 7)
			if a>>8 == 255 {
				buckets[i][j][k] = append(buckets[i][j][k], color)
				colorsCount++
			}
		}
	}

	// calculate bucket's averages
	var bucketsAverages []bucket
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				currentBucket := buckets[i][j][k]
				bucketLen := len(currentBucket)
				if bucketLen > 0 {
					var sums [3]int
					for _, color := range currentBucket {
						r, g, b, _ := color.RGBA()
						sums[0] += int(r >> 8)
						sums[1] += int(g >> 8)
						sums[2] += int(b >> 8)
					}
					bucketsAverages = append(bucketsAverages, bucket{
						Count: bucketLen,
						Color: color.RGBA{
							R: uint8(sums[0] / bucketLen),
							G: uint8(sums[1] / bucketLen),
							B: uint8(sums[2] / bucketLen),
							A: 255,
						},
					})
				}
			}
		}
	}

	// sort colors by cluster size
	sort.Sort(sort.Reverse(ByCount(bucketsAverages)))

	// extract color.Color from bucket, ignore small buckets
	colors := []color.Color{}
	for _, avg := range bucketsAverages {
		if float64(avg.Count)/float64(colorsCount) > config.SmallBucket {
			colors = append(colors, avg.Color)
		}
	}

	return colors
}

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

type colorSums struct {
	Red   float64
	Green float64
	Blue  float64
	Count float64
}

func ExtractColorsWithConfig(image image.Image, config Config) []color.Color {
	width := image.Bounds().Max.X
	height := image.Bounds().Max.Y

	// calculate downsizing ratio
	stepX := int(math.Max(float64(width)/config.DownSizeTo, 1))
	stepY := int(math.Max(float64(height)/config.DownSizeTo, 1))

	// load image's pixels into buckets
	var buckets [2][2][2]colorSums
	colorsCount := 0
	for x := 0; x < width; x += stepX {
		for y := 0; y < height; y += stepY {
			color := image.At(x, y)
			r, g, b, a := color.RGBA()
			r >>= 8
			g >>= 8
			b >>= 8
			i := r >> 7
			j := g >> 7
			k := b >> 7
			if a>>8 == 255 {
				buckets[i][j][k].Red += float64(r)
				buckets[i][j][k].Green += float64(g)
				buckets[i][j][k].Blue += float64(b)
				buckets[i][j][k].Count++
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
				if currentBucket.Count > 0 {
					bucketsAverages = append(bucketsAverages, bucket{
						Count: int(currentBucket.Count),
						Color: color.RGBA{
							R: uint8(currentBucket.Red / currentBucket.Count),
							G: uint8(currentBucket.Green / currentBucket.Count),
							B: uint8(currentBucket.Blue / currentBucket.Count),
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

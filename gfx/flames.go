package gfx

import (
	"image"
	"image/color"
	"math"
	"math/rand"
	"time"
)

type HeatMap struct {
	colorMap [37]color.RGBA
	width    int
	height   int
	data     []int
}

func NewHeatMap(width, height int) *HeatMap {
	colors := [37]color.RGBA{
		color.RGBA{0x00, 0x00, 0x00, 0x00},
		color.RGBA{0x07, 0x07, 0x07, 0x7F},
		color.RGBA{0x1f, 0x07, 0x07, 0x7F},
		color.RGBA{0x2f, 0x0f, 0x07, 0x7F},
		color.RGBA{0x47, 0x0f, 0x07, 0x7F},
		color.RGBA{0x57, 0x17, 0x07, 0x7F},
		color.RGBA{0x67, 0x1f, 0x07, 0x7F},
		color.RGBA{0x77, 0x1f, 0x07, 0xFF},
		color.RGBA{0x8f, 0x27, 0x07, 0xFF},
		color.RGBA{0x9f, 0x2f, 0x07, 0xFF},
		color.RGBA{0xaf, 0x3f, 0x07, 0xFF},
		color.RGBA{0xbf, 0x47, 0x07, 0xFF},
		color.RGBA{0xc7, 0x47, 0x07, 0xFF},
		color.RGBA{0xDF, 0x4F, 0x07, 0xFF},
		color.RGBA{0xDF, 0x57, 0x07, 0xFF},
		color.RGBA{0xDF, 0x57, 0x07, 0xFF},
		color.RGBA{0xD7, 0x5F, 0x07, 0xFF},
		color.RGBA{0xD7, 0x67, 0x0F, 0xFF},
		color.RGBA{0xcf, 0x6f, 0x0f, 0xFF},
		color.RGBA{0xcf, 0x77, 0x0f, 0xFF},
		color.RGBA{0xcf, 0x7f, 0x0f, 0xFF},
		color.RGBA{0xCF, 0x87, 0x17, 0xFF},
		color.RGBA{0xC7, 0x87, 0x17, 0xFF},
		color.RGBA{0xC7, 0x8F, 0x17, 0xFF},
		color.RGBA{0xC7, 0x97, 0x1F, 0xFF},
		color.RGBA{0xBF, 0x9F, 0x1F, 0xFF},
		color.RGBA{0xBF, 0x9F, 0x1F, 0xFF},
		color.RGBA{0xBF, 0xA7, 0x27, 0xFF},
		color.RGBA{0xBF, 0xA7, 0x27, 0xFF},
		color.RGBA{0xBF, 0xAF, 0x2F, 0xFF},
		color.RGBA{0xB7, 0xAF, 0x2F, 0xFF},
		color.RGBA{0xB7, 0xB7, 0x2F, 0xFF},
		color.RGBA{0xB7, 0xB7, 0x37, 0xFF},
		color.RGBA{0xCF, 0xCF, 0x6F, 0xFF},
		color.RGBA{0xDF, 0xDF, 0x9F, 0xFF},
		color.RGBA{0xEF, 0xEF, 0xC7, 0xFF},
		color.RGBA{0xFF, 0xFF, 0xFF, 0xFF},
	}

	buffer := make([]int, width*height, width*height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if y == height-1 {
				buffer[y*width+x] = 36
			} else {
				buffer[y*width+x] = 0
			}
		}
	}

	return &HeatMap{
		colorMap: colors,
		width:    width,
		height:   height,
		data:     buffer,
	}
}

func (hm *HeatMap) getMappedColor(x, y int) color.RGBA {
	return hm.colorMap[hm.data[y*hm.width+x]]
}

func (hm *HeatMap) SpreadHeat() {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for x := 0; x < hm.width; x++ {
		for y := 1; y < hm.height; y++ {

			src := y*hm.width + x
			pixel := hm.data[src]

			if pixel == 0 {
				hm.data[src-hm.width] = 0
			} else {
				randIdx := int(math.Round(r.Float64() * 3.0))
				dst := src - randIdx + 1
				hm.data[dst-hm.width] = pixel - (randIdx % 2)
			}
		}
	}
}

func (hm *HeatMap) GenerateTexture(wrapR, wrapS int32) (*Texture, error) {

	canvas := image.NewRGBA(image.Rect(0, 0, hm.width, hm.height))

	// build texture from heat map
	for y := 0; y < hm.height; y++ {
		for x := 0; x < hm.width; x++ {
			idx := canvas.PixOffset(x, y)
			color := hm.getMappedColor(x, y)
			canvas.Pix[idx+0] = color.R
			canvas.Pix[idx+1] = color.G
			canvas.Pix[idx+2] = color.B
			canvas.Pix[idx+3] = color.A
		}
	}

	return NewTexture(canvas, wrapR, wrapS)
}

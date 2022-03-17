package image

import (
	"github.com/easygithdev/gip/histogram"
	"github.com/easygithdev/gip/pixel"
)

const RED_CHANNEL = 0
const GREEN_CHANNEL = 1
const BLUE_CHANNEL = 2
const ALPHA_CHANNEL = 3

type GoImage struct {
	width    int
	height   int
	channels byte
	/* Palette-based image pixels */
	pixels [][]*pixel.GoPixel
	histo  *histogram.GoHistogram
}

func NewGoImage() *GoImage {
	img := GoImage{}
	img.histo = histogram.NewGoHistogram()
	return &img
}

func (gi *GoImage) SetDimension(w int, h int) {
	gi.width = w
	gi.height = h
	gi.pixels = make([][]*pixel.GoPixel, w)

	for i := range gi.pixels {
		gi.pixels[i] = make([]*pixel.GoPixel, h)
	}

}

func (gi *GoImage) GetDimension() (int, int) {
	return gi.width, gi.height
}

func (gi *GoImage) SetChannels(channels byte) {
	gi.channels = channels
}

func (gi *GoImage) GetChannels() byte {
	return gi.channels
}

func (gi *GoImage) CreatetPixel() *pixel.GoPixel {
	return pixel.NewGoPixel()
}

func (gi *GoImage) SetPixel(x int, y int, gp *pixel.GoPixel) {
	gi.pixels[x][y] = gp
}

func (gi *GoImage) GetPixel(x int, y int) *pixel.GoPixel {
	return gi.pixels[x][y]
}

func (gi *GoImage) GetHistogram() *histogram.GoHistogram {
	return gi.histo
}

func (gi *GoImage) Compute() {}

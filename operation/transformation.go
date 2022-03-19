package operation

import (
	"github.com/easygithdev/gip/image"
)

func transform(x, y int) (int, int) {
	var nx int = x + 10
	var ny int = y + 10
	return nx, ny
}

type GipTransformation struct {
}

func NewGipTransformation() *GipTransformation {
	transfo := GipTransformation{}
	return &transfo
}

func (conv *GipTransformation) Compute(img *image.GoImage) *image.GoImage {

	var dest *image.GoImage = new(image.GoImage)
	dest.SetChannels(img.GetChannels())
	dest.SetDimension(img.GetDimension())

	w, h := img.GetDimension()

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			pixel := img.GetPixel(x, y)
			nx, ny := transform(x, y)
			if nx > w-1 {
				nx = x
			}
			if ny > h-1 {
				ny = y
			}
			if nx < 0 {
				nx = 0
			}
			if ny < 0 {
				ny = 0
			}

			dest.SetPixel(nx, ny, pixel)

		}
	}

	return dest
}

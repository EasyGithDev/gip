package operation

import (
	"github.com/easygithdev/gip/image"
	"github.com/easygithdev/gip/pixel"
)

var THRESHOLD_BINARY = NewGipThreshold(
	func(gp *pixel.GoPixel) *pixel.GoPixel {

		if gp.GetRed() >= 128 {
			gp.SetRed(255)
		} else {
			gp.SetRed(0)
		}

		if gp.GetGreen() >= 128 {
			gp.SetGreen(255)
		} else {
			gp.SetGreen(0)
		}

		if gp.GetBlue() >= 128 {
			gp.SetBlue(255)
		} else {
			gp.SetBlue(0)
		}
		return gp
	},
)

var THRESHOLD_GRAY = NewGipThreshold(
	func(gp *pixel.GoPixel) *pixel.GoPixel {

		r := 0.2126 * float32(gp.GetRed())
		g := 0.7152 * float32(gp.GetGreen())
		b := 0.0722 * float32(gp.GetBlue())
		l := r + g + b
		gp.SetRed(byte(l))
		gp.SetGreen(byte(l))
		gp.SetBlue(byte(l))

		return gp

	},
)

type GipThreshold struct {
	threshold byte
	fn        func(gp *pixel.GoPixel) *pixel.GoPixel
}

func NewGipThreshold(fn func(gp *pixel.GoPixel) *pixel.GoPixel) *GipThreshold {
	gt := GipThreshold{fn: fn}
	return &gt
}

func (gt *GipThreshold) Compute(gi *image.GoImage) *image.GoImage {
	w, h := gi.GetDimension()

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			pixel := gt.fn(gi.GetPixel(x, y))
			gi.SetPixel(x, y, pixel)
		}
	}

	return gi
}
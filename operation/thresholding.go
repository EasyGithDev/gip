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

var THRESHOLD_AVG_GRAY = NewGipThreshold(
	func(gp *pixel.GoPixel) *pixel.GoPixel {

		r := float32(gp.GetRed())
		g := float32(gp.GetGreen())
		b := float32(gp.GetBlue())
		l := (r + g + b) / 3
		gp.SetRed(byte(l))
		gp.SetGreen(byte(l))
		gp.SetBlue(byte(l))

		return gp

	},
)

var THRESHOLD_FAST_GRAY = NewGipThreshold(
	func(gp *pixel.GoPixel) *pixel.GoPixel {

		gp.SetRed(gp.GetGreen())
		gp.SetBlue(gp.GetGreen())

		return gp

	},
)

var THRESHOLD_TRUE_GRAY = NewGipThreshold(
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

var THRESHOLD_NEGATIVE = NewGipThreshold(
	func(gp *pixel.GoPixel) *pixel.GoPixel {

		r := 255 - float32(gp.GetRed())
		g := 255 - float32(gp.GetGreen())
		b := 255 - float32(gp.GetBlue())
		gp.SetRed(byte(r))
		gp.SetGreen(byte(g))
		gp.SetBlue(byte(b))

		return gp

	},
)

var FN_BINARY_BY_VALUE = func(gp *pixel.GoPixel, val byte) *pixel.GoPixel {

	if gp.GetRed() >= val {
		gp.SetRed(255)
	} else {
		gp.SetRed(0)
	}

	if gp.GetGreen() >= val {
		gp.SetGreen(255)
	} else {
		gp.SetGreen(0)
	}

	if gp.GetBlue() >= val {
		gp.SetBlue(255)
	} else {
		gp.SetBlue(0)
	}
	return gp
}

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

func ComputeByValue(gi *image.GoImage, value byte) *image.GoImage {
	w, h := gi.GetDimension()

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			pixel := FN_BINARY_BY_VALUE(gi.GetPixel(x, y), value)
			gi.SetPixel(x, y, pixel)
		}
	}

	return gi
}

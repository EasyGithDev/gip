package gooperation

import (
	"github.com/easygithdev/gip/goimage"
)

var BLUR *Convolution = NewConvolution(
	[][]float32{
		{1.0, 1.0, 1.0},
		{1.0, 1.0, 1.0},
		{1.0, 1.0, 1.0},
	},
	1.0/9.0)

type Convolution struct {
	matrix [][]float32
	div    float32
}

func NewConvolution(matrix [][]float32, div float32) *Convolution {
	conv := Convolution{matrix: matrix, div: div}
	return &conv
}

func (conv *Convolution) Compute(img *goimage.GoImage) *goimage.GoImage {

	dest := img
	w, h := img.GetDimension()

	for c := byte(0); c < img.GetChannels(); c++ {
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				var sum float32 = 0

				for k := i - 1; k <= i+1; k++ {
					for l := j - 1; l <= j+1; l++ {
						if k > 0 && k < w && l > 0 && l < h {
							sum += float32(img.GetPixel(k, l).GetColor(c)) * conv.matrix[k-i+1][l-j+1]
						}
					}
				}
				dest.GetPixel(i, j).SetColor(c, byte(sum*conv.div))
			}
		}
	}

	return dest
}

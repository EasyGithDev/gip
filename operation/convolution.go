package operation

import (
	"github.com/easygithdev/gip/image"
)

var BLUR *Convolution = NewConvolution(
	[][]float32{
		{1.0, 1.0, 1.0},
		{1.0, 1.0, 1.0},
		{1.0, 1.0, 1.0},
	},
	1.0/9.0)

var IDENTITY *Convolution = NewConvolution(
	[][]float32{
		{0.0, 0.0, 0.0},
		{0.0, 0.0, 0.0},
		{0.0, 0.0, 0.0},
	},
	1.0)

var GAUSS_3x3 *Convolution = NewConvolution(
	[][]float32{
		{1.0, 2.0, 1.0},
		{2.0, 4.0, 2.0},
		{1.0, 2.0, 1.0},
	},
	1.0/16.0)

var SHARPNESS_IMPROVEMENT *Convolution = NewConvolution(
	[][]float32{
		{0.0, 1.0, 0.0},
		{1.0, -5.0, 1.0},
		{0.0, 1.0, 0.0},
	},

	1.0)

var EDGE_DETECTION_1 *Convolution = NewConvolution(
	[][]float32{
		{1.0, 0.0, -1.0},
		{0.0, 0.0, 0.0},
		{-1.0, 0.0, 1.0},
	},

	1.0)

var EDGE_DETECTION_2 *Convolution = NewConvolution(
	[][]float32{
		{1.0, 0.0, 1.0},
		{0.0, -4.0, 0.0},
		{1.0, 0.0, 1.0},
	},

	1.0)

var EDGE_DETECTION_3 *Convolution = NewConvolution(
	[][]float32{
		{-1.0, -1.0, -1.0},
		{-1.0, 8.0, -1.0},
		{-1.0, -1.0, -1.0},
	},

	1.0)

type Convolution struct {
	matrix    [][]float32
	threshold float32
}

func NewConvolution(matrix [][]float32, threshold float32) *Convolution {
	conv := Convolution{matrix: matrix, threshold: threshold}
	return &conv
}

func (conv *Convolution) Compute(img *image.GoImage) *image.GoImage {

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
				// convert float to byte
				color := sum * conv.threshold
				if color < 0.0 {
					color = 0.0
				} else if color > 255.0 {
					color = 255.0
				}
				bcolor := byte(color)

				dest.GetPixel(i, j).SetColor(c, bcolor)
			}
		}
	}

	return dest
}

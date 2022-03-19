package operation

import (
	"github.com/easygithdev/gip/image"
)

var BLUR *GipConvolution = NewGipConvolution(
	[][]float32{
		{1., 1., 1.},
		{1., 1., 1.},
		{1., 1., 1.},
	},
	3,
	1./9.,
	0.)

var IDENTITY *GipConvolution = NewGipConvolution(
	[][]float32{
		{0., 0., 0.},
		{0., 1., 0.},
		{0., 0., 0.},
	},
	3,
	1.,
	0.)

var GAUSS_3x3 *GipConvolution = NewGipConvolution(
	[][]float32{
		{1., 2., 1.},
		{2, 4., 2.},
		{1., 2., 1.},
	},
	3,
	1./16.,
	0.)

var SHARPNESS_IMPROVEMENT *GipConvolution = NewGipConvolution(
	[][]float32{
		{0., -1., 0.},
		{-1., 5, -1.},
		{0., -1., 0.},
	},
	3,
	1.,
	0.)

var EDGE_DETECTION_1 *GipConvolution = NewGipConvolution(
	[][]float32{
		{1., 0., -1.},
		{0., 0., 0.},
		{-1., 0., 1.},
	},
	3,
	1.,
	0.)

var EDGE_DETECTION_2 *GipConvolution = NewGipConvolution(
	[][]float32{
		{0., 1., 0.},
		{1., -4, 1.},
		{0., 1., 0.},
	},
	3,
	1.,
	0.)

var EDGE_DETECTION_3 *GipConvolution = NewGipConvolution(
	[][]float32{
		{-1., -1., -1.},
		{-1., 8., -1.},
		{-1., -1., -1.},
	},
	3,
	1.,
	0.)

var EMBOSS *GipConvolution = NewGipConvolution(
	[][]float32{
		{-2, -2, 0.},
		{-2, 6., 0.},
		{0., 0., 0.},
	},
	3,
	1.,
	0.)

var CONNEX *GipConvolution = NewGipConvolution(
	[][]float32{
		{0., -1., 0.},
		{-1., 4., -1.},
		{0., -1., 0.},
	},
	3,
	10./4.,
	128.)

type GipConvolution struct {
	kernel  [][]float32
	size    int
	divisor float32
	offset  float32
}

func NewGipConvolution(kernel [][]float32, size int, divisor float32, offset float32) *GipConvolution {
	conv := GipConvolution{kernel: kernel, divisor: divisor, offset: offset}
	return &conv
}

func (conv *GipConvolution) Compute(img *image.GoImage) *image.GoImage {

	dest := img
	w, h := img.GetDimension()

	var size int = conv.size / 2
	for channel := byte(0); channel < img.GetChannels(); channel++ {

		for y := size; y < h-size; y++ {
			for x := size; x < w-size; x++ {
				var sum float32 = 0.

				for line := size; line >= -size; line-- {
					for col := size; col >= -size; col-- {
						var pcolor byte = 0
						tx := x - col
						ty := y - line
						//println("tx ", tx, " ty ", ty)
						//	if y-line >= 0 && y-line <= h-1 && x-col >= 0 && x-col <= w-1 {

						pcolor = img.GetPixel(tx, ty).GetColor(channel)

						sum += float32(pcolor) * conv.kernel[line+size][col+size]
						//	println("ici sum", x, y, sum)

						// println(" line", line, " col", col, conv.kernel[col-x+1][line-y+1])
					}
				}
				//	os.Exit(0)

				//				println("ici sum", x, y, sum)
				//				os.Exit(0)
				// convert float to byte
				color := sum*conv.divisor + conv.offset

				//	println(" c", color)

				if color < 0 {
					color = 0
				} else if color > 255 {
					color = 255
				}
				bcolor := byte(color)
				//	println(" c", color, " bcol", bcolor)

				dest.GetPixel(x, y).SetColor(channel, bcolor)
			}
		}
	}

	return dest
}

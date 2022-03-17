package operation

import (
	"github.com/easygithdev/gip/image"
)

var BLUR *Convolution = NewConvolution(
	[][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	},
	9)

var IDENTITY *Convolution = NewConvolution(
	[][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	},
	1)

var GAUSS_3x3 *Convolution = NewConvolution(
	[][]int{
		{1, 2, 1},
		{2, 4, 2},
		{1, 2, 1},
	},
	16)

var SHARPNESS_IMPROVEMENT *Convolution = NewConvolution(
	[][]int{
		{0, -1, 0},
		{-1, 5, -1},
		{0, -1, 0},
	},

	1)

var EDGE_DETECTION_1 *Convolution = NewConvolution(
	[][]int{
		{1, 0, -1},
		{0, 0, 0},
		{-1, 0, 1},
	},

	1)

var EDGE_DETECTION_2 *Convolution = NewConvolution(
	[][]int{
		{0, 1, 0},
		{1, -4, 1},
		{0, 1, 0},
	},

	1)

var EDGE_DETECTION_3 *Convolution = NewConvolution(
	[][]int{
		{-1, -1, -1},
		{-1, 8, -1},
		{-1, -1, -1},
	},

	1)

var EMBOSS *Convolution = NewConvolution(
	[][]int{
		{-2, -2, 0},
		{-2, 6, 0},
		{0, 0, 0},
	},

	1)

type Convolution struct {
	kernel  [][]int
	divisor int
	offset  int
}

func NewConvolution(kernel [][]int, divisor int) *Convolution {
	conv := Convolution{kernel: kernel, divisor: divisor}
	return &conv
}

func (conv *Convolution) Compute(img *image.GoImage) *image.GoImage {

	dest := img
	w, h := img.GetDimension()

	var size int = 3 / 2
	for c := byte(0); c < img.GetChannels(); c++ {
		println(c)
		for y := size; y < h-size; y++ {
			for x := size; x < w-size; x++ {
				var sum int = 0

				for line := size; line >= -size; line-- {
					for col := size; col >= -size; col-- {
						var pcolor byte = 0
						tx := x - col
						ty := y - line
						//println("tx ", tx, " ty ", ty)
						//	if y-line >= 0 && y-line <= h-1 && x-col >= 0 && x-col <= w-1 {

						pcolor = img.GetPixel(tx, ty).GetColor(c)

						// println(" x", x+col, " y", y+line, " pix", pcolor, " conv", conv.kernel[line+size][col+size])

						//	}
						// println("la sum", x, y, sum)

						sum += int(pcolor) * conv.kernel[line+size][col+size]
						//	println("ici sum", x, y, sum)

						// println(" line", line, " col", col, conv.kernel[col-x+1][line-y+1])
					}
				}
				//	os.Exit(0)

				//				println("ici sum", x, y, sum)
				//				os.Exit(0)
				// convert float to byte
				color := sum/conv.divisor + conv.offset

				//	println(" c", color)

				if color < 0 {
					color = 0
				} else if color > 255 {
					color = 255
				}
				bcolor := byte(color)
				//	println(" c", color, " bcol", bcolor)

				dest.GetPixel(x, y).SetColor(c, bcolor)
			}
		}
	}

	return dest
}

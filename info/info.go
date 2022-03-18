package info

import (
	"fmt"
	"math"

	"github.com/easygithdev/gip/pixel"
)

const INFO_MAX_COLOR = 256
const INFO_MAX_CHANNEL = 4

type GipInfo struct {
	histo   [INFO_MAX_CHANNEL][INFO_MAX_COLOR]float64
	average [INFO_MAX_CHANNEL]float64
	sigma   [INFO_MAX_CHANNEL]float64
	min     [INFO_MAX_CHANNEL]float64
	max     [INFO_MAX_CHANNEL]float64
}

func NewGipInfo() *GipInfo {
	gi := GipInfo{}
	return &gi
}

func (gi *GipInfo) SetColor(pixel *pixel.GoPixel) {
	gi.histo[0][pixel.GetRed()]++
	gi.histo[1][pixel.GetGreen()]++
	gi.histo[2][pixel.GetBlue()]++
	gi.histo[3][pixel.GetAlpha()]++
}

func (gi *GipInfo) Compute() {

	sum := make([]float64, INFO_MAX_CHANNEL)
	for channel := 0; channel < INFO_MAX_CHANNEL; channel++ {

		gi.max[channel] = gi.histo[channel][0]
		gi.min[channel] = gi.histo[channel][0]

		for j := range gi.histo[channel] {
			sum[channel] += gi.histo[channel][j]

			if gi.histo[channel][j] > gi.max[channel] {
				gi.max[channel] = gi.histo[channel][j]
			}

			if gi.histo[channel][j] < gi.min[channel] {
				gi.min[channel] = gi.histo[channel][j]
			}
		}
	}

	for i := 0; i < INFO_MAX_CHANNEL; i++ {
		gi.average[i] = sum[i] / float64(INFO_MAX_COLOR)
	}

	for i := 0; i < INFO_MAX_CHANNEL; i++ {
		for j := range gi.histo[i] {
			gi.sigma[i] += (gi.histo[i][j] - gi.average[i]) * (gi.histo[i][j] - gi.average[i])
		}
		gi.sigma[i] /= float64(INFO_MAX_COLOR)
		gi.sigma[i] = math.Sqrt(gi.sigma[i])
	}
}

func (gi *GipInfo) ShowInfo(channel byte) {
	fmt.Printf("Min:%f\n", gi.min[channel])
	fmt.Printf("Max:%f\n", gi.max[channel])
	fmt.Printf("Average:%f\n", gi.average[channel])
	fmt.Printf("Sigma:%f\n", gi.sigma[channel])
}

func (gi *GipInfo) ShowHisto(channel byte) {

	k := gi.max[channel] * 0.019

	for i := 0; i < len(gi.histo[channel]); i++ {
		for j := 0; j < int(gi.histo[channel][i]/k); j++ {
			print("*")
		}
		println()
	}

}

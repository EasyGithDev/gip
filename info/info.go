package info

import (
	"fmt"
	"math"

	"github.com/easygithdev/gip/pixel"
)

const INFO_MAX_COLOR = 256
const INFO_MAX_CHANNEL = 4

type GipInfo struct {
	histo   [INFO_MAX_CHANNEL][INFO_MAX_COLOR]int
	average [INFO_MAX_CHANNEL]float64
	sigma   [INFO_MAX_CHANNEL]float64
	min     [INFO_MAX_CHANNEL]int
	max     [INFO_MAX_CHANNEL]int
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

	// Compute the average of each channel
	for channel := 0; channel < INFO_MAX_CHANNEL; channel++ {

		var nbPix int = 0
		var sum int = 0

		gi.max[channel] = gi.histo[channel][0]
		gi.min[channel] = gi.histo[channel][0]

		for j := range gi.histo[channel] {

			nbPix += gi.histo[channel][j]
			sum += j * gi.histo[channel][j]

			if gi.histo[channel][j] > gi.max[channel] {
				gi.max[channel] = gi.histo[channel][j]
			}

			if gi.histo[channel][j] < gi.min[channel] {
				gi.min[channel] = gi.histo[channel][j]
			}
		}

		gi.average[channel] = float64(sum) / float64(nbPix)
	}

	// Compute the sigma of each channel
	for channel := 0; channel < INFO_MAX_CHANNEL; channel++ {
		var variance float64 = 0.0
		var nbPix float64 = 0
		for j := range gi.histo[channel] {
			val := float64(gi.histo[channel][j])
			nbPix += val
			variance += (val - gi.average[channel]) * (val - gi.average[channel])
		}
		gi.sigma[channel] = math.Sqrt(variance / nbPix)
	}
}

func (gi *GipInfo) ShowInfo(channel byte) {
	fmt.Printf("Min:%d\n", gi.min[channel])
	fmt.Printf("Max:%d\n", gi.max[channel])
	fmt.Printf("Average:%f\n", gi.average[channel])
	fmt.Printf("Sigma:%f\n", gi.sigma[channel])
}

func (gi *GipInfo) ShowHisto(channel byte) {

	k := float32(gi.max[channel]) * 0.019

	for i := 0; i < len(gi.histo[channel]); i++ {
		for j := 0; j < int(float32(gi.histo[channel][i])*k); j++ {
			print("*")
		}
		println()
	}

}

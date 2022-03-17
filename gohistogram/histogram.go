package gohistogram

import "github.com/easygithdev/gip/gopixel"

type GoHistogram struct {
	t [4][256]int
}

func NewGoHistogram() *GoHistogram {
	gh := GoHistogram{}

	for i := 0; i < 4; i++ {
		for j := range gh.t[i] {
			gh.t[i][j] = 0
		}
	}
	return &gh
}

func (gh *GoHistogram) SetColor(pixel *gopixel.GoPixel) {
	gh.t[0][pixel.GetRed()]++
	gh.t[1][pixel.GetGreen()]++
	gh.t[2][pixel.GetBlue()]++
	gh.t[3][pixel.GetAlpha()]++
}

func (gh *GoHistogram) Show(channel byte) {

	max := gh.t[channel][0]
	for j := 0; j < len(gh.t[channel]); j++ {
		if gh.t[channel][j] > max {
			max = gh.t[channel][j]
		}
	}
	k := float32(max) * 0.019

	for i := 0; i < len(gh.t[channel]); i++ {
		for j := 0; j < gh.t[channel][i]/int(k); j++ {
			print("*")
		}
		println()
	}

}

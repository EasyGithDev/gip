package gorgba

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/easygithdev/gip/goimage"
)

func ReadRGB(filename string, width int, height int, gi *goimage.GoImage) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer f.Close()

	gi.SetDimension(width, height)
	gi.SetChannels(3)
	histo := gi.GetHistogram()
	buf := make([]byte, 3)
	i, j := 0, 0

	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		if n > 0 {

			pixel := gi.CreatetPixel()
			pixel.SetRGBA(buf[0], buf[1], buf[2], 0xff)
			gi.SetPixel(i, j, pixel)

			histo.SetColor(pixel)

			if i < width-1 {
				i++
			} else {
				i = 0
				j++

			}
		}
	}
}

func WriteRGB(filename string, gi *goimage.GoImage) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer f.Close()

	width, heigth := gi.GetDimension()

	for y := 0; y < width; y++ {
		for x := 0; x < heigth; x++ {
			pixel := gi.GetPixel(x, y)
			buf := []byte{pixel.GetRed(), pixel.GetGreen(), pixel.GetBlue()}
			_, err := f.Write(buf)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
}

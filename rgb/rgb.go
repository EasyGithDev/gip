package rgb

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/easygithdev/gip/image"
)

func ReadRGB(filename string, width int, height int, gi *image.GoImage) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer f.Close()

	gi.SetDimension(width, height)
	gi.SetChannels(3)

	infos := gi.GetInfos()
	buf := make([]byte, 3*width)
	x, y := 0, 0

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

			for i := 0; i < n; i = i + 3 {

				r := buf[i]
				g := buf[i+1]
				b := buf[i+2]

				// remvoving the black background
				if r == 0 && g == 0 && b == 0 {
					continue
				}

				pixel := gi.GetPixel(x, y)
				pixel.SetRGBA(r, g, b, 0xff)
				gi.SetPixel(x, y, pixel)
				infos.SetColor(pixel)
				x++
			}
			x = 0
			y++

		}
	}
}

func WriteRGB(filename string, gi *image.GoImage) {
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

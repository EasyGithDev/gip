package gray

import (
	"fmt"
	"log"
	"os"

	"github.com/easygithdev/gip/image"
)

func ReadGray(filename string, width int, height int, gi *image.GoImage) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer f.Close()

	gi.SetDimension(width, height)
	gi.SetChannels(1)

	infos := gi.GetInfos()
	buf := make([]byte, width*height)

	n, err := f.Read(buf)

	if err != nil {
		fmt.Println(err)

	}
	if n > 0 {

		for x := 0; x < width; x++ {
			for y := 0; y < height; y++ {
				r := buf[x*width+y]
				pixel := gi.GetPixel(x, y)
				pixel.SetRGBA(r, 0xff, 0xff, 0xff)
				gi.SetPixel(x, y, pixel)
				infos.SetColor(pixel)
			}
		}

	}

}

func WriteGray(filename string, gi *image.GoImage) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer f.Close()

	width, heigth := gi.GetDimension()

	for y := 0; y < width; y++ {
		for x := 0; x < heigth; x++ {
			pixel := gi.GetPixel(x, y)
			buf := []byte{pixel.GetRed()}
			_, err := f.Write(buf)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
}

package gray

import (
	"fmt"
	"log"
	"os"

	"github.com/easygithdev/gip/image"
)

func ReadGray(filename string, width int, height int, img *image.GoImage) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer f.Close()

	img.SetDimension(width, height)
	img.SetChannels(1)

	infos := img.GetInfos()
	buf := make([]byte, img.GetSize())
	x, y := 0, 0
	n, err := f.Read(buf)

	if err != nil {
		fmt.Println(err)
	}

	if n > 0 {
		for i := 0; i < n; i++ {
			gray := buf[i]
			pixel := img.GetPixel(x, y)
			pixel.SetRGBA(gray, 0xff, 0xff, 0xff)
			img.SetPixel(x, y, pixel)
			infos.SetColor(pixel)

			if x < width-1 {
				x++
			} else {
				x = 0
				y++
			}
		}
	}

}

func WriteGray(filename string, img *image.GoImage) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer f.Close()

	width, heigth := img.GetDimension()

	for y := 0; y < width; y++ {
		for x := 0; x < heigth; x++ {
			pixel := img.GetPixel(x, y)
			buf := []byte{pixel.GetRed()}
			_, err := f.Write(buf)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
}

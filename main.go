package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"github.com/easygithdev/gip/goimage"
	"github.com/easygithdev/gip/gooperation"
	"github.com/easygithdev/gip/gorgba"
)

const WIDTH = 512
const HEIGTH = 512

func createRaster(gi *goimage.GoImage) *canvas.Raster {
	return canvas.NewRasterWithPixels(
		func(x, y, w, h int) color.Color {

			pixel := gi.GetPixel(x, y)
			return color.RGBA{pixel.GetRed(),
				pixel.GetGreen(),
				pixel.GetBlue(),
				pixel.GetAlpha()}
		})
}

func createWindow(title string, gi *goimage.GoImage) {

	myApp := app.New()
	wSrc := myApp.NewWindow(title)

	wSrc.Resize(fyne.NewSize(WIDTH, HEIGTH))
	wSrc.SetContent(createRaster(gi))
	wSrc.ShowAndRun()
}

func main() {

	img := goimage.NewGoImage()

	// gorgba.ReadRGBA("result.rgb", img)
	gorgba.ReadRGB("data/lena_color.rgb", WIDTH, HEIGTH, img)

	// matrix := [][]float32{
	// 	{0.0, 0.0, 0.0},
	// 	{-2.0, 1.0, 0.0},
	// 	{-2.0, -2.0, 0.0},
	// }

	// conv := gooperation.NewConvolution(matrix, 1.0/9.0)

	// dest := conv.Compute(img)

	dest := gooperation.BLUR.Compute(img)
	createWindow("dest", dest)

	// gorgba.WriteRGB("./result.rgb", img)

	// myApp := app.New()
	// wSrc := myApp.NewWindow("Lena src")
	// wDest := myApp.NewWindow("Lena dest")

	// wSrc.Resize(fyne.NewSize(WIDTH, HEIGTH))
	// wSrc.SetContent(createRaster(img))
	// wSrc.ShowAndRun()

	// wDest.Resize(fyne.NewSize(WIDTH, HEIGTH))
	// wDest.SetContent(createRaster(img))
	// wDest.ShowAndRun()

}

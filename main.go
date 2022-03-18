package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"github.com/easygithdev/gip/image"
	"github.com/easygithdev/gip/rgb"
)

const WIDTH = 512
const HEIGTH = 512

func createRaster(gi *image.GoImage) *canvas.Raster {
	return canvas.NewRasterWithPixels(
		func(x, y, w, h int) color.Color {

			pixel := gi.GetPixel(x, y)
			return color.RGBA{pixel.GetRed(),
				pixel.GetGreen(),
				pixel.GetBlue(),
				pixel.GetAlpha()}
		})
}

func createWindow(title string, gi *image.GoImage) {

	myApp := app.New()
	wSrc := myApp.NewWindow(title)

	wSrc.Resize(fyne.NewSize(WIDTH, HEIGTH))
	wSrc.SetContent(createRaster(gi))
	wSrc.ShowAndRun()
}

func main() {

	img := image.NewGoImage()

	// gorgba.ReadRGBA("result.rgb", img)
	rgb.ReadRGB("data/lena_color.rgb", WIDTH, HEIGTH, img)
	img.GetInfos().Compute()
	img.GetInfos().ShowInfo(0)

	// matrix := [][]float32{
	// 	{0.0, 0.0, 0.0},
	// 	{-2.0, 1.0, 0.0},
	// 	{-2.0, -2.0, 0.0},
	// }

	// conv := gooperation.NewConvolution(matrix, 1.0/9.0)

	// dest := conv.Compute(img)

	// dest := operation.BLUR.Compute(img)
	//dest := operation.GAUSS_3x3.Compute(img)
	// dest := operation.EDGE_DETECTION_1.Compute(img)
	//dest := operation.EDGE_DETECTION_2.Compute(img)
	//dest := operation.EDGE_DETECTION_3.Compute(img)
	//dest := operation.SHARPNESS_IMPROVEMENT.Compute(img)
	//dest := operation.IDENTITY.Compute(img)
	//dest := operation.EMBOSS.Compute(img)

	//	operation.THRESHOLD_BINARY.Compute(img)
	// operation.THRESHOLD_GRAY.Compute(img)
	createWindow("dest", img)

	// rgb.WriteRGB("./nice-error.rgb", img)

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

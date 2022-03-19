package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"github.com/easygithdev/gip/gray"
	"github.com/easygithdev/gip/image"
)

const WIDTH = 512
const HEIGTH = 512
const RGB_FILE = "data/lena.rgb"
const GRAY_FILE = "data/lena.gray"

func showRgb(title string, gi *image.GoImage) {

	myApp := app.New()
	w := myApp.NewWindow(title)

	raster := canvas.NewRasterWithPixels(
		func(x, y, w, h int) color.Color {

			pixel := gi.GetPixel(x, y)
			return color.RGBA{pixel.GetRed(),
				pixel.GetGreen(),
				pixel.GetBlue(),
				pixel.GetAlpha()}
		})

	w.Resize(fyne.NewSize(WIDTH, HEIGTH))
	w.SetContent(raster)
	w.ShowAndRun()
}

func showGray(title string, gi *image.GoImage) {

	myApp := app.New()
	w := myApp.NewWindow(title)

	raster := canvas.NewRasterWithPixels(
		func(x, y, w, h int) color.Color {

			pixel := gi.GetPixel(x, y)
			return color.Gray{pixel.GetRed()}
		})

	w.Resize(fyne.NewSize(WIDTH, HEIGTH))
	w.SetContent(raster)
	w.ShowAndRun()
}

func showInfo(img *image.GoImage) {

	println("********************************************")
	println("***                RED                   ***")
	println("********************************************")
	img.GetInfos().ShowInfo(0)

	println("********************************************")
	println("***                GREEN                 ***")
	println("********************************************")
	img.GetInfos().ShowInfo(1)

	println("********************************************")
	println("***                BLUE                  ***")
	println("********************************************")
	img.GetInfos().ShowInfo(2)
}

func main() {

	/////////////////////////////////////////////////////////////
	// Create the image
	/////////////////////////////////////////////////////////////

	img := image.NewGoImage()

	/////////////////////////////////////////////////////////////
	// Read the file
	/////////////////////////////////////////////////////////////

	// rgb.ReadRGB(RGB_FILE, WIDTH, HEIGTH, img)
	gray.ReadGray(GRAY_FILE, WIDTH, HEIGTH, img)

	/////////////////////////////////////////////////////////////
	// Compute and show statistique
	/////////////////////////////////////////////////////////////

	img.GetInfos().Compute()
	showInfo(img)

	/////////////////////////////////////////////////////////////
	// Show the src image
	/////////////////////////////////////////////////////////////

	showGray("Src", img)

	/////////////////////////////////////////////////////////////
	// Compute the threshold
	/////////////////////////////////////////////////////////////

	// operation.THRESHOLD_BINARY.Compute(img)
	// operation.THRESHOLD_TRUE_GRAY.Compute(img)
	// operation.THRESHOLD_AVG_GRAY.Compute(img)
	// operation.THRESHOLD_FAST_GRAY.Compute(img)
	// operation.THRESHOLD_NEGATIVE.Compute(img)
	// operation.ThresholdByValue(img, 220)
	// operation.ThresholdByAverage(img)

	/////////////////////////////////////////////////////////////
	// Compute the convolution
	/////////////////////////////////////////////////////////////

	// Show the dest image
	// showImage("Dest", img)

	// matrix := [][]float32{
	// 	{0.0, 0.0, 0.0},
	// 	{-2.0, 1.0, 0.0},
	// 	{-2.0, -2.0, 0.0},
	// }

	// conv := gooperation.NewConvolution(matrix, 1.0/9.0)

	// dest := conv.Compute(img)

	// dest := operation.BLUR.Compute(img)
	// dest := operation.GAUSS_3x3.Compute(img)
	// dest := operation.EDGE_DETECTION_1.Compute(img)
	//dest := operation.EDGE_DETECTION_2.Compute(img)
	//dest := operation.EDGE_DETECTION_3.Compute(img)
	//dest := operation.SHARPNESS_IMPROVEMENT.Compute(img)
	//dest := operation.IDENTITY.Compute(img)
	// dest := operation.EMBOSS.Compute(img)
	// dest := operation.CONNEX.Compute(img)

	/////////////////////////////////////////////////////////////
	// Show the dest image
	/////////////////////////////////////////////////////////////

	// showImage("Dest", dest)
}

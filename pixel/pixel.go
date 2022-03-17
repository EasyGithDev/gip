package pixel

type GoPixel struct {
	red   byte
	green byte
	blue  byte
	alpha byte
}

func NewGoPixel() *GoPixel {
	pixel := GoPixel{}
	return &pixel
}

func (gp *GoPixel) SetRGBA(r byte, g byte, b byte, a byte) {
	gp.red = r
	gp.green = g
	gp.blue = b
	gp.alpha = a
}

func (gp *GoPixel) GetRed() byte {
	return gp.red
}

func (gp *GoPixel) GetGreen() byte {

	return gp.green
}

func (gp *GoPixel) GetBlue() byte {

	return gp.blue
}

func (gp *GoPixel) GetAlpha() byte {
	return gp.alpha
}

func (gp *GoPixel) SetRed(red byte) {
	gp.red = red
}

func (gp *GoPixel) SetGreen(green byte) {
	gp.green = green
}

func (gp *GoPixel) SetBlue(blue byte) {
	gp.blue = blue
}

func (gp *GoPixel) SetAlpha(alpha byte) {
	gp.alpha = alpha
}

func (gp *GoPixel) GetColor(channel byte) byte {
	switch channel {
	case 0:
		return gp.red
	case 1:
		return gp.green
	case 2:
		return gp.blue
	case 3:
		return gp.alpha
	}
	return 0
}

func (gp *GoPixel) SetColor(channel byte, val byte) {
	switch channel {
	case 0:
		gp.red = val
	case 1:
		gp.green = val
	case 2:
		gp.blue = val
	case 3:
		gp.alpha = val
	}

}

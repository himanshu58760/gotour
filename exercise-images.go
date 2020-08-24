package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	width, height int
	pColor        uint8
}

func (self *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, self.width, self.height)
}

// set the image color model
func (self *Image) ColorModel() color.Model {
	return color.RGBAModel
}

// set the image color per pixel
func (self *Image) At(x, y int) color.Color {
	return color.RGBA{
		self.pColor + uint8(x), self.pColor + uint8(y), 255, 255}
}

func main() {
	m := Image{100,100,56}
	pic.ShowImage(&m)
}

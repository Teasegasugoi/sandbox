package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/image/image.go;l=38
// Image is a finite rectangular grid of color.Color values taken from a color
// model.
// type Image interface {
// 	ColorModel() color.Model
// 	Bounds() Rectangle
// 	At(x, y int) color.Color
// }

type Image struct {
	w int
	h int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

func (i Image) At(x, y int) color.Color {
	v := uint8(x * y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{256, 256}
	pic.ShowImage(m)
}

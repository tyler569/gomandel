// Mandelbrot in Go
// Copyright (c) 2014, Tyler Philbrick
// See COPYING for license details

package mandel

import (
	//"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

type corner struct {
	r, i float64
	x, y int
}

type Mandel struct {
	ul, lr         corner
	ratior, ratioi float64
}

func (m Mandel) ColorModel() color.Model {
	return color.NRGBAModel
}

func (m Mandel) Bounds() image.Rectangle {
	return image.Rect(m.ul.x, m.ul.y, m.lr.x, m.lr.y)
}

func (m Mandel) At(x, y int) color.Color {
	c := m.CoordAt(x, y)
	var z complex128
	for i := uint8(0); i < 64; i++ {
		z = z*z + c
		if cmplx.Abs(z) > 10 {
			return color.NRGBA{i * 4, i * 4, i * 4, 0xff}
		}
	}
	return color.NRGBA{0xff, 0xff, 0xff, 0xff}
}

func (m Mandel) CoordAt(x, y int) complex128 {
	r := float64(x)/m.ratior + m.ul.r
	i := float64(y)/m.ratioi + m.ul.i
	return complex(r, i)
}

func calcRatios(m *Mandel) {
	mranger := m.lr.r - m.ul.r
	mrangei := m.lr.i - m.ul.i
	irangex := float64(m.lr.x - m.ul.x)
	irangey := float64(m.lr.y - m.ul.y)
	m.ratior = irangex / mranger
	m.ratioi = irangey / mrangei
}

func MandelImage(w *io.Writer) {
	c1 := corner{-2, 1, 0, 0}
	c2 := corner{1, -1, 1500, 1000}
	m := Mandel{c1, c2, 0, 0}
	calcRatios(&m)

	err := png.Encode(w, m)
	if err != nil {
		panic(err)
	}
}

package main

import (
	"image/png"
	"math"
	"os"

	"github.com/unixpickle/model3d/model2d"
	"github.com/unixpickle/model3d/model3d"
)

type BitmapSlab struct {
	Thickness     float64
	Origin        model3d.Coord3D
	Width, Height float64

	bm                  *model2d.Bitmap
	imgWidth, imgHeight float64
}

func NewBitmapSlab(url string, c model2d.ColorBitFunc, origin model3d.Coord3D, thickness, width, height float64) *BitmapSlab {
	f, err := os.Open(url)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	img, err := png.Decode(f)
	if err != nil {
		panic(err)
	}
	return &BitmapSlab{
		Thickness: thickness,
		Origin:    origin,
		Width:     width,
		Height:    height,
		bm:        model2d.NewBitmapImage(img, c),
		imgWidth:  float64(img.Bounds().Dx()),
		imgHeight: float64(img.Bounds().Dy()),
	}
}

func (b *BitmapSlab) Min() model3d.Coord3D {
	return b.Origin
}

func (b *BitmapSlab) Max() model3d.Coord3D {
	return model3d.Coord3D{
		X: b.Origin.X + b.Width,
		Y: b.Origin.Y + b.Height,
		Z: b.Origin.Z + b.Thickness,
	}
}

func (b *BitmapSlab) Contains(c model3d.Coord3D) bool {
	if c.Z < b.Origin.Z || c.Z > b.Origin.Z+b.Thickness {
		return false
	}

	return b.bm.Get(int(math.Round(c.X*(b.imgWidth/b.Width))), int(math.Round(c.Y*(b.imgHeight/b.Height))))
}

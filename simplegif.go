package main

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/gif"
	"math"
	"os"
)

type Circle struct {
	X, Y, R float64
}

func (c *Circle) Brightness(x, y float64) uint8 {
	var dx, dy float64 = c.X - x, c.Y - y
	d := math.Sqrt(dx*dx+dy*dy) / c.R
	if d > 1 {
		return 0
	} else {
		return 255
	}
}
func main() {
	const (
		w, h int     = 256, 256
		R    float64 = 50
		// 大约一秒30帧
		frameN int = 30
		delay  int = 100 / frameN
	)
	// 216色 + 透明
	var palette = append(palette.WebSafe, color.Opaque)
	var images []*image.Paletted
	var delays []int
	var disposals []byte
	var hw, hh float64 = float64(w / 2), float64(h / 2)
	circles := []*Circle{&Circle{}, &Circle{}, &Circle{}}
	for step := 0; step < frameN; step++ {
		img := image.NewPaletted(image.Rect(0, 0, w, h), palette)
		images = append(images, img)
		delays = append(delays, delay)
		disposals = append(disposals, gif.DisposalPrevious)
		θ := 2.0 * math.Pi / float64(frameN) * float64(step)
		for i, circle := range circles {
			θ0 := 2 * math.Pi / 3 * float64(i)
			circle.X = hw - 30*math.Sin(θ0) - 30*math.Sin(θ0+θ)
			circle.Y = hh - 30*math.Cos(θ0) - 30*math.Cos(θ0+θ)
			circle.R = R
		}
		for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
				cr := circles[0].Brightness(float64(x), float64(y))
				cg := circles[1].Brightness(float64(x), float64(y))
				cb := circles[2].Brightness(float64(x), float64(y))
				if cr|cg|cb > 0x00 {
					img.Set(x, y, color.RGBA{cr, cg, cb, 255})
				} else {
					img.Set(x, y, color.Transparent)
				}
			}
		}
	}
	// 把GIF写入文件rgb.gif
	f, err := os.Create("rgb.gif")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_ = gif.EncodeAll(f, &gif.GIF{
		Image:    images,
		Delay:    delays,
		Disposal: disposals,
	})
}

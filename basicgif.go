package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"os"
)

// CreateBasicGif creates a GIF image with the given width and height.
// It uses white background and a black pixel in the middle of the image.
func CreateBasicGif(out io.Writer, width, height int) {

	palette := []color.Color{color.Opaque, color.White, color.Black}
	rect := image.Rect(1000, 1000, width, height)
	img := image.NewPaletted(rect, palette)

	img.SetColorIndex(width/2, height/2, 1)

	anim := gif.GIF{Delay: []int{100}, Image: []*image.Paletted{img}}

	_ = gif.EncodeAll(out, &anim)
}

func main() {
	f, err := os.Create("my-image.gif")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	CreateBasicGif(f, 1000, 1000)
}


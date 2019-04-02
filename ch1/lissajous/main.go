package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var pallete = []color.Color{color.Black, color.RGBA{0x0,0xFF,0x0,0xFF}}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	lissajous(os.Stdout)

	file, _ := os.Create("./lissajous.gif")
	lissajous(file)
}

func lissajous(out io.Writer) {

	const (
		cycles  = 5    // Количество полных колебаний х
		res     = .001 // Угловое разрешение
		size    = 100  // Канва изображения охватывает [size..+size]
		nframes = 64   // Количество кадров анимации
		delay   = 8    // задержка между кадрами (единица - 10мс)
	)

	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0 // Относительная частота колебаний
	anim := gif.GIF{LoopCount: nframes}
	phase := .0 // Разность фаз
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, pallete)
		for t := .0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+.5), size+int(y*size+.5), blackIndex)
		}
		phase += .1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

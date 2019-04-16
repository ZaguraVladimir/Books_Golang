package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var pallete = []color.Color{color.Black, color.RGBA{0x0, 0xFF, 0x0, 0xFF}}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(writer http.ResponseWriter, request *http.Request) {

	if err := request.ParseForm(); err != nil {
		log.Println(err)
	}

	cycles := 5
	if v, ok := request.Form["cycles"]; ok {
		cycles, _ = strconv.Atoi(v[0])
	}

	lissajous(writer, cycles)
}

func lissajous(out io.Writer, cycles int) {
	//cycles  = Количество полных колебаний х
	const (
		res     = .001 // Угловое разрешение
		size    = 300  // Канва изображения охватывает [size..+size]
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
		for t := .0; t < float64(cycles)*2*math.Pi; t += res {
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

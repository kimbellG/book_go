package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"time"
)

var palette = []color.Color{color.White, color.RGBA{0, 117, 37, 1}, color.RGBA{225, 255, 0, 1}}

const (
	whiteIndex = 0
	greenIndex = 1
	yelowIndex = 2
)

type parametrs struct {
	nframse, delay    int
	res, cycles, size float64
}

//Lissajous создание гифки на основе константы
func Lissajous(out io.Writer, p parametrs) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframse = 64
		delay   = 8
	)

	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: p.nframse}
	phase := 0.0
	col := rand.Int() % 3

	for i := 0; i < p.nframse; i++ {
		rect := image.Rect(0, 0, int(2*p.size+1), int(2*p.size+1))
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < p.cycles*2*math.Pi; t += p.res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			if col == 0 || col == 1 {
				img.SetColorIndex(int(p.size+x*p.size+0.5), int(p.size+y*p.size+0.5), greenIndex)
			} else {
				img.SetColorIndex(int(p.size+x*p.size+0.5), int(p.size+y*p.size+0.5), yelowIndex)
			}
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, p.delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"
)

func main() {

	file, err := os.Create(".log")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create file: ", err)
		os.Exit(2)
	}

	log.SetOutput(file)

	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
		numdev                 = 4
	)

	var y, x int

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		var ys [numdev]float64

		for i := 0; i < numdev; i++ {
			ys[i] = float64(y)/height*(ymax-ymin) + ymin
			y += i
		}

		for px := 0; px < width; px++ {
			var xs [numdev]float64
			var zs [numdev]complex128

			for i := 0; i < numdev; i++ {
				xs[i] = float64(x)/width*(xmax-xmin) + xmin
				x += i

				zs[i] = complex(xs[i], ys[i])

			}

			img.Set(px, py, mandelbrot(zs))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(zs [4]complex128) color.Color {
	const iteration = 200
	const contrast = 15

	var ns [4]uint8
	var v complex128

	for i := 0; i < 4; i++ {
		v = 0
		for ns[i] = 0; ns[i] < iteration; ns[i]++ {
			v = v*v + zs[i]
			if cmplx.Abs(v) > 2 {
				break

			}
		}
	}

	var average uint8

	for _, n := range ns {
		if n == iteration {
			continue
		}

		average += n
	}

	average /= 4

	return color.YCbCr{255 - contrast*average, 255 - contrast*average, 255 - contrast*average}
}

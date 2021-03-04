package svggrath

import (
	"fmt"
	"io"
	"math"
)

type parametrs struct {
	width, Height int
	cells         int
	xyrange       float64
	xyscale       float64
	zscale        float64
	angle         float64
}

var Parametrs parametrs = parametrs{}

var sin30, cos30 float64

func init() {
	Parametrs.width = 600
	Parametrs.Height = 320
	Parametrs.cells = 100
	Parametrs.xyrange = 30.0
	Parametrs.angle = math.Pi / 6
	fmt.Printf("from init()\n")
	sin30, cos30 = math.Sin(Parametrs.angle), math.Cos(Parametrs.angle)
}

func Draw(out io.Writer) {
	Parametrs.xyscale = float64(Parametrs.width) / 2 / Parametrs.xyrange
	Parametrs.zscale = float64(Parametrs.Height) * 0.4

	fmt.Printf("Start drawing, start parametrs: \n"+
		"width = %d, Height = %d, cells = %d\n"+
		"xyrange = %f, xyscale = %f, zscale = %f, angle = %f\n",
		Parametrs.width, Parametrs.Height, Parametrs.cells, Parametrs.xyrange, Parametrs.xyscale,
		Parametrs.zscale, Parametrs.angle)

	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0,7' "+
		"width='%d' height='%d'>", Parametrs.width, Parametrs.Height)

	for i := 0; i < Parametrs.cells; i++ {
		for j := 0; j < Parametrs.cells; j++ {
			ax, ay, _ := corner(i+1, j)
			bx, by, color := corner(i, j)
			cx, cy, _ := corner(i, j+1)
			dx, dy, _ := corner(i+1, j+1)
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g '\n"+
				"style='fill: %v; fill-opacity:0.7; stroke: black; ' />\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)

		}
	}
	fmt.Fprintf(out, "</svg>")
}

func corner(i int, j int) (float64, float64, string) {
	x := Parametrs.xyrange * (float64(i)/float64(Parametrs.cells) - 0.5)
	y := Parametrs.xyrange * (float64(j)/float64(Parametrs.cells) - 0.5)

	z := f(x, y)
	var color string

	if z > 0.4 {
		color = "#ff0000"
	} else {
		color = "#0000ff"
	}

	sx := float64(Parametrs.width)/2 + (x-y)*cos30*Parametrs.xyscale
	sy := float64(Parametrs.Height)/2 + (x+y)*sin30*Parametrs.xyscale - z*Parametrs.zscale
	return sx, sy, color
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

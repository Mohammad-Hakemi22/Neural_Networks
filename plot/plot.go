package plot

import (
	"image/color"
	"log"
	"os"

	"github.com/Mohammad-Hakemi22/NN/datasets"
	"github.com/Mohammad-Hakemi22/NN/utility"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

var CMap map[float64]color.Color = map[float64]color.Color{
	0: color.RGBA{0, 0, 255, 255},
	1: color.RGBA{0, 255, 0, 255},
	2: color.RGBA{255, 0, 0, 255},
}

func Scatter() {
	x, y := datasets.SpiralData(100, 3)
	_ = y
	p := plot.New()
	for key, val := range CMap {
		xys := plotter.XYs{}
		for i := 0; i < len(x); i++ {
			if y[i] == int(key) {
				xys = append(xys, plotter.XY{X: x[i].X, Y: x[i].Y})
			}
		}
		scatter, err := plotter.NewScatter(xys)
		utility.ErrorHandler("can't initial scatter plot", err)
		scatter.Color = val
		p.Add(scatter)
	}
	wt, err := p.WriterTo(450, 400, "png")
	utility.ErrorHandler("can't write to writer", err)
	path := "plot.png"
	f, err := os.Create(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		err := f.Close()
		utility.ErrorHandler("something went wrong in closing file", err)
	}()
	wt.WriteTo(f)
}

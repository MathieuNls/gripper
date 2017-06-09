package gripper

import (
	"sort"
	"time"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/plotutil"
	"github.com/gonum/plot/vg"
)

type performance struct {
	lines  [][]*point
	labels []string
}

type point struct {
	x float64
	y float64
}

func PerfPlotter() *performance {

	perf := &performance{
		lines:  [][]*point{},
		labels: []string{},
	}

	return perf
}

func (perf *performance) Analyze(
	callback func(size int),
	max int, increment int,
	retries int, label string) *performance {

	pts := make([]*point, max/increment)

	for i := 0; i < max; i += increment {

		times := make([]float64, retries)

		for r := 0; r < retries; r++ {
			start := time.Now()
			callback(i)
			times[r] = float64(time.Since(start).Nanoseconds())
		}

		sort.Float64s(times)
		pts[i] = &point{
			x: float64(i),
			y: float64(times[len(times)/2]),
		}
	}

	perf.lines = append(perf.lines, pts)
	perf.labels = append(perf.labels, label)

	return perf
}

func (perf *performance) Plot(xLabel string, yLabel string,
	legend string, path string) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = legend
	p.X.Label.Text = xLabel
	p.Y.Label.Text = yLabel

	plotLines := []interface{}{}

	for i := 0; i < len(perf.lines); i++ {

		pts := make(plotter.XYs, len(perf.lines[i]))

		for j := 0; j < len(perf.lines[i]); j++ {
			pts[j].X = perf.lines[i][j].x
			pts[j].Y = perf.lines[i][j].y
		}

		plotLines = append(plotLines, perf.labels[i], pts)

	}

	err = plotutil.AddLinePoints(p, plotLines...)
	if err != nil {
		panic(err)
	}

	if err := p.Save(4*vg.Inch, 4*vg.Inch, path); err != nil {
		panic(err)
	}
}

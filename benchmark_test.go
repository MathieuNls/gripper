package gripper

import (
	"os"
	"testing"
)

func TestPerfPlotter(t *testing.T) {

	var p *performance
	p = PerfPlotter()
	if p == nil {
		t.Error("Expected plotter got ", p)
	}
}

func TestAnalyze(t *testing.T) {

	var p *performance
	p = PerfPlotter()

	p.Analyze(
		func(size int) {

		}, 10, 1, 10, "nothing",
	)
	if p == nil || len(p.lines) != 1 || len(p.lines[0]) != 10 ||
		len(p.labels) != 1 || p.labels[0] != "nothing" {
		t.Error("Expected plotter got ", p)
	}

	p.Plot("x", "y", "x/y", "testing.png")
	if _, err := os.Stat("testing.png"); os.IsNotExist(err) {
		t.Error("Expected plot to be written to disk")
	}

}

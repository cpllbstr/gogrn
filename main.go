package main

import (
	"os"

	"github.com/cpllbstr/gogrn/grn"
	"github.com/cpllbstr/gogrn/slv"
)

func main() {
	params := slv.StartParams{
		Velocity: 1,
		Length:   10,
	}
	out, err := os.Create("./dat/output.dat")
	if err != nil {
		panic(err)
	}
	defer out.Close()
	b := grn.ThreeBodyModel{
		K: [3]float64{1, 1, 1},
		M: [3]float64{30, 0.5, 17},
	}
	StMach := slv.StateMachFromModel(b, params.Velocity, params.Length, 0.0001, 0, 50)
	StMach.Mute = true
	slv.SimulateAv(StMach, b, out)
	er1 := slv.GoGnuPlot("./plt/plotav.sh", out, "~/univer/rpz/notes/img/optmass.pdf")
	if er1 != nil {
		panic(er1)
	}
	slv.
}

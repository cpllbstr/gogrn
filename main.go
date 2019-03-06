package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/cpllbstr/gogrn/ode/envs"

	"github.com/cpllbstr/gogrn/grn"
	"github.com/cpllbstr/gogrn/ode"

	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/plot/plotter"
)

func CreateDot(A *mat.Dense) ode.Func2Var {
	return func(x interface{}, y interface{}) interface{} {
		vec := y.(*mat.VecDense)
		res := mat.NewVecDense(vec.Len(), nil)
		res.MulVec(A, vec)
		return res
	}
}

func main() {
	contactC := grn.ThreeBodyModel{
		K: [3]float64{1.5, 1.5, 1.5},
		M: [3]float64{1, 1, 1},
	}
	freeC := grn.ThreeBodyModel{
		K: [3]float64{0, 1.5, 1.5},
		M: [3]float64{1, 1, 1},
	}

	vec := mat.NewVecDense(6, []float64{0, 0, 0, -1, -1, -1})

	gonumrk4 := ode.Rk4FromEnv(envs.GonumVecDenseEnv)
	var StateFuncs = map[grn.StateEnum]ode.Func2Var{
		grn.Started:     CreateDot(contactC.GenMatr()),
		grn.HitWall:     CreateDot(contactC.GenMatr()),
		grn.BouncedBack: CreateDot(freeC.GenMatr()),
		grn.NonPhis:     nil,
	}

	StMach := grn.NewStateMachine(gonumrk4, *vec, 1.5, 0.01, 0, 25, StateFuncs)
	//res := vec.RawVector()
	//fmt.Println(res.Data)
	/*
		plt, err := plot.New()
		if err != nil {
			panic(err)
		}

		plt.Title.Text = "ThreeBodyModel"
		plt.X.Label.Text = "X"
		plt.Y.Label.Text = "T"

		plt.Add()
		err = plotutil.AddLines(plt,
			"First", randomPoints(15),
			"Second", randomPoints(15),
			"Third", randomPoints(15))
		if err != nil {
			panic(err)
		}
		// Save the plot to a PNG file.
		if err := plt.Save(20*vg.Centimeter, 20*vg.Centimeter, "points.svg"); err != nil {
			panic(err)
		}

		pts1 := make(plotter.XYs, 0)
		pts2 := make(plotter.XYs, 0)
		pts3 := make(plotter.XYs, 0)
	*/

	gnuplt, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := gnuplt.Close(); err != nil {
			panic(err)
		}
	}()

	for ok := StMach.UpdateState(); ok == nil; ok = StMach.UpdateState() {
		StMach.NextStep()
		StepCond := StMach.GetCondition()
		gnuplt.WriteString(fmt.Sprintf("%v, %v, %v, %v \n", StMach.CurTime, StepCond.X[0], StepCond.X[1], StepCond.X[2]))
	}

}

/*
func CondiPlot(c grn.Condition, time float64) plotter.XYs {
	for i := 0; i < len(c.X); i++ {

	}
}
*/

func randomPoints(n int) plotter.XYs {
	pts := make(plotter.XYs, n)
	for i := range pts {
		if i == 0 {
			pts[i].X = rand.Float64()
		} else {
			pts[i].X = pts[i-1].X + rand.Float64()
		}
		pts[i].Y = pts[i].X + 10*rand.Float64()
	}
	return pts
}

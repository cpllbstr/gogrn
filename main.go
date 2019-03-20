package main

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/cpllbstr/gogrn/ode/envs"

	"github.com/cpllbstr/gogrn/grn"
	"github.com/cpllbstr/gogrn/ode"

	"gonum.org/v1/gonum/mat"
)

func CreateDot(A *mat.Dense) ode.Func2Var {
	return func(x interface{}, y interface{}) interface{} {
		vec := y.(*mat.VecDense)
		res := mat.NewVecDense(vec.Len(), nil)
		res.MulVec(A, vec)
		return res
	}
}

func plotCondition(gnuplt *os.File, StepCond grn.Condition, curtime, length float64) {
	gnuplt.WriteString(fmt.Sprintf("%v %v %v %v ", curtime, StepCond.X[0]+length, StepCond.X[1]+2.*length, StepCond.X[2]+3.*length))
	gnuplt.WriteString(fmt.Sprintf("%v %v %v \n", StepCond.V[0], StepCond.V[1], StepCond.V[2]))
}

func plotEnergy(gnuplt *os.File, var1, var2, E float64) {
	gnuplt.WriteString(fmt.Sprintf("%v %v %v\n", var1, var2, E))
}

func Simulate(StMach grn.StateMachine, file ...*os.File) (grn.Condition, grn.StateEnum) {
	for ok := StMach.UpdateState(); ok == nil; ok = StMach.UpdateState() {
		StMach.NextStep()
		if len(file) != 0 {
			plotCondition(file[0], StMach.GetCondition(), StMach.CurTime, StMach.Length)
		}
	}
	return StMach.GetCondition(), StMach.CurState
}

func CalcEnergy(c grn.Condition, b grn.ThreeBodyModel, v float64) (float64, float64, float64) {
	Emidd := ((math.Pow(b.M[0]*c.V[0]+b.M[1]*c.V[1]+b.M[2]*c.V[2], 2.)) / (b.M[0] + b.M[1] + b.M[2])) / 2. //энергия центра масс
	Estrt := (b.M[0] + b.M[1] + b.M[2]) * (math.Pow(v, 2.)) / 2.                                           //энергия до соударения
	dE := (Estrt - Emidd) / Estrt                                                                          //коэфф запасания энергии
	return Emidd, Estrt, dE
}

func StateMachFromModel(b grn.ThreeBodyModel, v float64) grn.StateMachine {
	vec := mat.NewVecDense(6, []float64{0, 0, 0, -v, -v, -v})
	gonumrk4 := ode.Rk4FromEnv(envs.GonumVecDenseEnv)
	var StateFuncs = map[grn.StateEnum]ode.Func2Var{
		grn.Started:     CreateDot(b.GenMatr()),
		grn.HitWall:     CreateDot(b.GenMatr()),
		grn.BouncedBack: CreateDot(b.GenMatrFree()),
		grn.NonPhis:     nil,
	}
	return grn.NewStateMachine(gonumrk4, *vec, 1.5, 0.01, 0, 10, StateFuncs)
}

func main() {
	strartVel := 1. //начальная скорость

	plotter, err := os.Create("./dat/output.dat")
	if err != nil {
		panic(err)
	}
	defer plotter.Close()

	energplot, err := os.Create("./dat/energy.dat")
	if err != nil {
		panic(err)
	}
	defer energplot.Close()

	b := grn.ThreeBodyModel{
		K: [3]float64{1.5, 1.5, 1.5},
		M: [3]float64{1, 1, 1},
	}
	StMach := StateMachFromModel(b, strartVel)
	Simulate(StMach, plotter)
	str := time.Now()
	for m := 1.; m <= 100; m += 0.1 {

		//fmt.Println(m)
		b := grn.ThreeBodyModel{
			K: [3]float64{1.5, 1.5, 1.5},
			M: [3]float64{m, 1, 1},
		}
		StMach := StateMachFromModel(b, strartVel)
		StMach.Mute = true
		StMach.Length = 10.
		Cond, st := Simulate(StMach)
		if st != grn.NonPhis {
			_, _, coeff := CalcEnergy(Cond, b, strartVel)
			plotEnergy(energplot, b.M[0], b.M[1], coeff)
		} else {
			fmt.Println("NonPhis", st)
			break
		}
	}
	fmt.Printf("Time elapsed: %v\n", time.Since(str))
}

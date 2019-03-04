package main

import (
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
	for {
		StMach.NextStep()
		StMach.UpdateState()
		//cc := grn.ConditionFromVec(*StMach.CurCond)
		//	fmt.Println("X:", cc.X)
	}

}

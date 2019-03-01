package main

import (
	"fmt"

	"github.com/cpllbstr/gogrn/grn"
	"github.com/cpllbstr/gogrn/ode/envs"

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
	/*freeC := grn.ThreeBodyModel{
		K: [3]float64{0, 1.5, 1.5},
		M: [3]float64{1, 1, 1},
	}*/
	//free := freeC.GenMatr()
	contact := contactC.GenMatr()
	//fmt.Printf("% v\n", mat.Formatted(contact))
	vec := mat.NewVecDense(6, []float64{0, 0, 0, -1, -1, -1})
	//fmt.Printf("% v\n%v\n", mat.Formatted(free), mat.Formatted(vec))
	//A := mat.NewDense(2, 2, []float64{5, 4, 4, 5})
	//Y0 := mat.NewVecDense(2, []float64{2, 0})
	dot := CreateDot(contact)
	rk4 := ode.Rk4FromEnv(envs.GonumVecDenseEnv)
	res := rk4(vec, 0., 0.02, dot)
	/*for i := 0; i < 100; i++ {
		res = rk4(res, float64(i)*0.01, 0.01, dot)
	}*/

	fmt.Println(mat.Formatted(res.(*mat.VecDense)))

}

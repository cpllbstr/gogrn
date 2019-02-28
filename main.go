package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

type ff func(param, param) interface{}

type function struct {
	fun ff
}
type param struct {
	val interface{}
}

func Call(p1 param, p2 param, f function) interface{} {
	return f.fun(p1, p2)
}

/*
func call(f fu, param1, param2 interface{}) interface{} {
	res := f(param1, param2)
	return res
}*/

func plus(x param, y param) interface{} {
	return x.val.(float64) + y.val.(float64)
}
func dot(x param, y param) interface{} {
	m := x.val.(mat.Matrix)
	v := y.val.(mat.Vector)
	var res mat.VecDense
	res.MulVec(m, v)
	return res
}

func main() {
	/*
		contactC := grn.Body{
			K: [3]float64{1.5, 1.5, 1.5},
			M: [3]float64{1, 1, 1},
			L: 1.5,
		}

		freeC := grn.Body{
			K: [3]float64{0, 1.5, 1.5},
			M: [3]float64{1, 1, 1},
			L: 1.5,
		}
			free := freeC.GenMatr()
			contact := contactC.GenMatr()
			fmt.Printf("% v\n", mat.Formatted(contact))
			vec := mat.NewVecDense(6, []float64{0, 0, 0, 1, 1, 1})
			fmt.Printf("% v\n%v\n", mat.Formatted(free), mat.Formatted(vec))
	*/
	matr := mat.NewDense(2, 2, []float64{1, 2, 3, 4})
	vec1 := mat.NewVecDense(2, []float64{1, 2})
	fmt.Printf("% v\n%v\n\n", mat.Formatted(matr), mat.Formatted(vec1))

	x := param{val: matr}
	y := param{val: vec1}
	f := function{fun: dot}
	res := Call(x, y, f)

	out := res.(mat.VecDense)
	fmt.Println(mat.Formatted(&out))

	/*
		x := param{val: 10.}
		y := param{val: 20.}
		f := function{fun: plus}
		x = param{val: contact}
		y = param{val: vec}
		f = function{fun: dot}
		res := Call(x, y, f)

		out := res.(mat.VecDense)
		fmt.Println(mat.Formatted(&out))
	*/
}

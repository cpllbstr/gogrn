package envs

import (
	"gonum.org/v1/gonum/mat"
)

//Float64Env - predefined MathEnv for float64 type
var Float64Env float64Env

type float64Env struct {
}

func (e float64Env) MultFloat(f float64, x interface{}) interface{} {
	return f * x.(float64)
}

func (e float64Env) SubFloat(x interface{}, f float64) interface{} {
	return x.(float64) / f
}

func (e float64Env) Summ(args ...interface{}) interface{} {
	summ := 0.
	for i := 0; i < len(args); i++ {
		summ += args[i].(float64)
	}
	return summ
}

//Complex128Env - predefined MathEnv for complex128 type
var Complex128Env complex128Env

type complex128Env struct {
}

func (e complex128Env) MultFloat(f float64, x interface{}) interface{} {
	return complex(f, 0) * x.(complex128)
}

func (e complex128Env) SubFloat(x interface{}, f float64) interface{} {
	return x.(complex128) / complex(f, 0)
}

func (e complex128Env) Summ(args ...interface{}) interface{} {
	summ := complex(0., 0.)
	for i := 0; i < len(args); i++ {
		summ += args[i].(complex128)
	}
	return summ
}

//GonumVecDenseEnv - predefined MathEnv for gonum DenseVec  type
var GonumVecDenseEnv gonumDenseEnv

type gonumDenseEnv struct {
}

func (e gonumDenseEnv) MultFloat(f float64, x interface{}) interface{} {
	vec := x.(*mat.VecDense)
	res := mat.NewVecDense(vec.Len(), nil)
	res.ScaleVec(f, vec)
	return res
}

func (e gonumDenseEnv) SubFloat(x interface{}, f float64) interface{} {
	vec := x.(*mat.VecDense)
	res := mat.NewVecDense(vec.Len(), nil)
	res.ScaleVec(1/f, vec)
	return res
}

func (e gonumDenseEnv) Summ(args ...interface{}) interface{} {
	summ := mat.NewVecDense(args[0].(*mat.VecDense).Len(), nil)
	for i := 0; i < len(args); i++ {
		x := args[i].(*mat.VecDense)
		summ.AddVec(summ, x)
	}
	return summ
}

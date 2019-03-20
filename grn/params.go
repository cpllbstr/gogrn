package grn

import (
	"gonum.org/v1/gonum/mat"
)

//ThreeBodyModel - Body for Math model
type ThreeBodyModel struct {
	K [3]float64
	M [3]float64
}

//Condition - describes condition of 3 body model for time t
type Condition struct {
	V []float64
	X []float64
}

func ConditionFromVec(vec mat.VecDense) Condition {
	rawvec := vec.RawVector().Data
	return Condition{
		V: rawvec[3:6],
		X: rawvec[0:3],
	}

	/*return Condition{
		V: [3]float64{vec.At(0, 0), vec.At(1, 0), vec.At(2, 0)},
		X: [3]float64{vec.At(3, 0), vec.At(4, 0), vec.At(5, 0)},
	}*/
}

//GenMatr - generates matrix for setted Body
func (c ThreeBodyModel) GenMatr() *mat.Dense {
	data := []float64{
		0., 0., 0., 1., 0., 0.,
		0., 0., 0., 0., 1., 0.,
		0., 0., 0., 0., 0., 1.,
		-((c.K[0] + c.K[1]) / c.M[0]), c.K[1] / c.M[0], 0, 0, 0, 0,
		c.K[1] / c.M[1], -((c.K[1] + c.K[2]) / c.M[1]), c.K[2] / c.M[1], 0, 0, 0,
		0, c.K[2] / c.M[2], -c.K[2] / c.M[2], 0, 0, 0,
	}
	return mat.NewDense(6, 6, data)
}

func (c ThreeBodyModel) GenMatrFree() *mat.Dense {
	c.K[0] = 0.
	data := []float64{
		0., 0., 0., 1., 0., 0.,
		0., 0., 0., 0., 1., 0.,
		0., 0., 0., 0., 0., 1.,
		-((c.K[0] + c.K[1]) / c.M[0]), c.K[1] / c.M[0], 0, 0, 0, 0,
		c.K[1] / c.M[1], -((c.K[1] + c.K[2]) / c.M[1]), c.K[2] / c.M[1], 0, 0, 0,
		0, c.K[2] / c.M[2], -c.K[2] / c.M[2], 0, 0, 0,
	}
	return mat.NewDense(6, 6, data)
}

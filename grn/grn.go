package grn

import (
	"gonum.org/v1/gonum/mat"
)

//Conditions - conditions for Math model
type Conditions struct {
	K1 float64
	K2 float64
	K3 float64
	M1 float64
	M2 float64
	M3 float64
	L  float64
}

//GenMatr - generates matrix for setted Conditions
func (c Conditions) GenMatr() *mat.Dense {
	data := []float64{
		0., 0., 0., 1., 0., 0.,
		0., 0., 0., 0., 1., 0.,
		0., 0., 0., 0., 0., 1.,
		-((c.K1 + c.K2) / c.M1), c.K1 / c.M1, 0, 0, 0, 0,
		c.K2 / c.M2, -((c.K2 + c.K3) / c.M2), c.K3 / c.M2, 0, 0, 0,
		0, c.K3 / c.M3, c.K3 / c.M3, 0, 0, 0,
	}
	return mat.NewDense(6, 6, data)
}

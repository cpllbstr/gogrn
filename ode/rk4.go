package ode

import (
	"gonum.org/v1/gonum/mat"
)

type RK4 struct {
	y0 interface{}
    interface{}
}

//Rk4 - Runge-Kutta4
/*
 y(t_0) = y_0 - Border conditions
 dy/dt = f(t, y) - equation
 t - start time
 h - step
*//*
func Rk4(yi, t, h float64, f func(interface{}, interface{}) interface{}) {
	k_1 := h * f(t_n, y_n)
	k_2 := h * f(t_n+h/2., y_n+k_1/2.)
	k_3 := h * f(t_n+h/2., y_n+k_2/2.)
	k_4 := h * f(t_n+h, y_n+k_3)

	return y_n + 1/6.*(k_1+2*k_2+2*k_3+k_4)
}*/

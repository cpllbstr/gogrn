package ode

//Func2Var - type of function with 2 variables
/*
	Example:

	func plus(x interface{}, y interface{}) interface{} {
		return x.(float64) + y.(float64)
	}
*/
type Func2Var func(interface{}, interface{}) interface{}

//Rk4Func - desccribes how rk4 should look
/*
	Rk4 for any type of y
	!You should define MathEnviroment for your type of data

	rk4(y0 interface, time, step float64, func Func2Var)

	y(t_0) = y_0 - Border conditions
 	dy/dt = f(t, y) - equation
 	t - start time
	h - time step
*/
type Rk4Func func(interface{}, float64, float64, Func2Var) interface{}

//MathEnv - describes how to operate with your type of data
/*
	X - your type of data
	MultFloat(float64, interface{}) interface{}
	returns X as float*X
	SubFloat(interface{}, float64) interface{}
	returns X as X/float
	Summ(...interface{}) interface{}
	returns X as Summ of args
*/
type MathEnv interface {
	MultFloat(float64, interface{}) interface{}
	Summ(...interface{}) interface{}
	SubFloat(interface{}, float64) interface{}
}

//Rk4FromEnv - Runge-Kutta4 for any type of input data
/*
	Rk4 for any type of y
	You should define MathEnviroment for your type of data
	y(t_0) = y_0 - Border conditions
 	dy/dt = f(t, y) - equation
 	t - start time
	h - time step
*/
func Rk4FromEnv(env MathEnv) Rk4Func {
	return func(y interface{}, t, h float64, f Func2Var) interface{} {
		k1 := env.MultFloat(h, f(t, y))                                     //k_1 = h*f(t,y)
		k2 := env.MultFloat(h, f(t+h/2., env.Summ(y, env.SubFloat(k1, 2)))) //k_2 = h*f(t+h/2,y+k_1/2)
		k3 := env.MultFloat(h, f(t+h/2., env.Summ(y, env.SubFloat(k2, 2)))) //k_3 = h*f(t+h/2,y+k_2/2)
		k4 := env.MultFloat(h, f(t+h, env.Summ(y, k3)))                     //k_2 = h*f(t+h,y+k_3)
		return env.Summ(y, env.SubFloat(env.Summ(k1, env.MultFloat(2, k2), env.MultFloat(2, k3), k4), 6))
	}
}

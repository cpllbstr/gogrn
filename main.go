package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"runtime"
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

func StateMachFromModel(b grn.ThreeBodyModel, v, length, step, starttime, fintime float64) grn.StateMachine {
	vec := mat.NewVecDense(6, []float64{0, 0, 0, -v, -v, -v})
	gonumrk4 := ode.Rk4FromEnv(envs.GonumVecDenseEnv)
	var StateFuncs = map[grn.StateEnum]ode.Func2Var{
		grn.Started:     CreateDot(b.GenMatr()),
		grn.HitWall:     CreateDot(b.GenMatr()),
		grn.BouncedBack: CreateDot(b.GenMatrFree()),
		grn.NonPhis:     nil,
	}
	return grn.NewStateMachine(gonumrk4, *vec, length, step, starttime, fintime, StateFuncs)
}

func FindStiff(k float64) bool {
	SpringLength := 1.
	vel := 1.
	for m := 1.; m <= 100; m += 1 {
		b := grn.ThreeBodyModel{
			K: [3]float64{k, k, k},
			M: [3]float64{m, 1, 1},
		}
		StMach := StateMachFromModel(b, vel, SpringLength, 0.01, 0, 30)
		StMach.Mute = true
		_, st := Simulate(StMach)
		if st == grn.NonPhis {
			return false
		}
	}
	return true
}

func Variate2Masses(i, j int, vari0, varin, varj0, varjn, gridstep float64, params startParams, b grn.ThreeBodyModel, c chan string) {
	str := time.Now()
	var output string
	for m1 := vari0; m1 <= varin; m1 += gridstep {
		for m2 := varj0; m2 <= varjn; m2 += gridstep {
			b.M[i] = m1
			b.M[j] = m2
			StMach := StateMachFromModel(b, params.Velocity, params.Length, 0.01, 0, 30)
			StMach.Mute = true
			Cond, st := Simulate(StMach)
			if st != grn.NonPhis {
				_, _, coeff := CalcEnergy(Cond, b, params.Velocity)
				//plotEnergy(fil, m1, m2, coeff)
				output = fmt.Sprintln(output, b.M[0], b.M[1], b.M[2], coeff)
			} else {
				log.Println("NonPhis on Mass:", m1, m2)
				break
			}
		}
	}
	log.Printf("Time elapsed: %v\n", time.Since(str))
	c <- output
}

func Variate2Stiffs(i, j int, vari0, varin, varj0, varjn, gridstep float64, params startParams, b grn.ThreeBodyModel, c chan string) {
	str := time.Now()
	var output string
	for m1 := vari0; m1 <= varin; m1 += gridstep {
		for m2 := varj0; m2 <= varjn; m2 += gridstep {
			b := grn.ThreeBodyModel{
				K: [3]float64{1, 1, 1},
				M: [3]float64{1, 1, 1},
			}
			b.K[i] = m1
			b.K[j] = m2
			StMach := StateMachFromModel(b, params.Velocity, params.Length, 0.01, 0, 30)
			StMach.Mute = true
			Cond, st := Simulate(StMach)
			if st != grn.NonPhis {
				_, _, coeff := CalcEnergy(Cond, b, params.Velocity)
				output = fmt.Sprintln(output, b.M[0], b.M[1], b.M[2], coeff)
				//plotEnergy(fil, m1, m2, coeff)
			} else {
				log.Println("NonPhis on Mass:", m1, m2)
				break
			}
		}
	}
	log.Printf("Time elapsed: %v\n", time.Since(str))
	c <- output
}

func barrier(nrouts int, ch chan string, fil *os.File) {
	for r := 0; r < nrouts; r++ {
		x := <-ch
		log.Println("goroutine ", r, "writed", len(x))
		fil.WriteString(x)
	}
}

func Variate2Params(typ string, i, j, nrouts int, total, gridstep float64, fil *os.File, params startParams) {
	type ff func(int, int, float64, float64, float64, float64, float64, startParams, grn.ThreeBodyModel, chan string)
	h, m, s := time.Now().Clock()
	fmt.Printf("Evaluating energy since: %02v:%02v:%02v\nNumber of goroutines: %v\nGridstep: %v\n", h, m, s, nrouts, gridstep)
	var f ff
	switch typ {
	case "m":
		f = Variate2Masses
	case "k":
		f = Variate2Stiffs
	}
	b := grn.ThreeBodyModel{
		K: [3]float64{1, 1, 1},
		M: [3]float64{1, 1, 1},
	}
	strt := time.Now()
	step := total / float64(nrouts)
	ch := make(chan string, nrouts)
	for n := 0; n < nrouts; n++ {
		str := 1 + float64(n)*step
		fin := str + step
		go f(i, j, str, fin, 1, total, gridstep, params, b, ch)
	}
	barrier(nrouts, ch, fil)
	fmt.Println("All goroutines finished in:", time.Since(strt))
}

func Variate3Masses(total, gridstep float64, out *os.File, params startParams) {
	nrouts := runtime.NumCPU()
	ch := make(chan string, nrouts)
	step := gridstep * float64(nrouts)
	strt := time.Now()
	for stepm := 0.1; stepm <= total; stepm += step {
		for i := 0; i <= nrouts; i++ {
			if total-(stepm+float64(i)*gridstep) > 0 {
				b := grn.ThreeBodyModel{
					K: [3]float64{1, 1, 1},
					M: [3]float64{stepm + float64(i)*gridstep, 1, 1},
				}
				go Variate2Masses(1, 2, 0.1, total, 0.1, total, gridstep, params, b, ch)
			} else {
				go func(c chan string) {
					c <- ""
				}(ch)
			}
		}
		barrier(nrouts, ch, out)
	}
	fmt.Println("Three massses evaluated in: ", time.Since(strt))
}

func Variate3Stiffs(total, gridstep float64, out *os.File, params startParams) {
	nrouts := runtime.NumCPU()
	ch := make(chan string, nrouts)
	step := gridstep * float64(nrouts)
	strt := time.Now()
	for stepm := 0.1; stepm <= total; stepm += step {
		for i := 0; i <= nrouts; i++ {
			if total-(stepm+float64(i)*gridstep) > 0 {
				b := grn.ThreeBodyModel{
					K: [3]float64{stepm + float64(i)*gridstep, 1, 1},
					M: [3]float64{1, 1, 1},
				}
				go Variate2Stiffs(1, 2, 0.1, total, 0.1, total, gridstep, params, b, ch)
			} else {
				go func(c chan string) {
					c <- ""
				}(ch)
			}
		}
		barrier(nrouts, ch, out)
	}
	fmt.Println("Three stiffs evaluated in: ", time.Since(strt))
}

func testFunc(i int) {
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	log.Println("FIN", i)
}

type startParams struct {
	Velocity float64
	Length   float64
}

func main() {

	params := startParams{
		Velocity: 1,
		Length:   10,
	}
	/*enm1m2, err := os.Create("./dat/enm1m2.dat")
	if err != nil {
		panic(err)
	}
	defer enm1m2.Close()
	/*enm2m3, err := os.Create("./dat/enm2m3.dat")
	if err != nil {
		panic(err)
	}
	enm1m3, err := os.Create("./dat/enm1m3.dat")
	if err != nil {
		panic(err)
	}
	defer func() {
		enm1m2.Close()
		enm1m3.Close()
		enm2m3.Close()
	}()*/
	//Variate2Params("m", 0, 1, 8, 50, 1, enm1m2, params)

	//Variate2Params("m", 1., 2., 8, 50, 0.2, enm2m3, params)
	//Variate2Params("m", 0., 2., 8, 50, 0.2, enm1m3, params)
	three, err := os.Create("./dat/tst.dat")
	if err != nil {
		panic(err)
	}
	fmt.Println("Tst")
	Variate3Stiffs(10, 1, three, params)
}

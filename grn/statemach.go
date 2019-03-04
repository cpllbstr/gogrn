package grn

import (
	"errors"
	"fmt"

	"github.com/cpllbstr/gogrn/ode"
	"gonum.org/v1/gonum/mat"
)

//StateMachine - state machine struct
type StateMachine struct {
	CurState  StateEnum
	StateFunc map[StateEnum]ode.Func2Var
	Stepper   ode.Rk4Func
	CurCond   *mat.VecDense
	Length    float64
	Step      float64
	CurTime   float64
	FinTime   float64
}

//NewStateMachine - returns new state machine with setted params
func NewStateMachine(stepper ode.Rk4Func, currCond mat.VecDense, length, step, starttime, fintime float64, statepars map[StateEnum]ode.Func2Var) StateMachine {
	return StateMachine{
		CurState:  Started,
		StateFunc: statepars,
		Stepper:   stepper,
		CurCond:   &currCond,
		Length:    length,
		Step:      step,
		CurTime:   starttime,
		FinTime:   fintime,
	}

}

func (st *StateMachine) NextStep() {
	res := st.Stepper(st.CurCond, st.CurTime, st.Step, st.StateFunc[st.CurState])
	rr := res.(*mat.VecDense)
	st.CurCond = rr
	st.CurTime += st.Step
}

//UpdateState - checks the conditions of state machine
func (st *StateMachine) UpdateState() error {
	cond := ConditionFromVec(*st.CurCond)
	if cond.X[0] <= -st.Length || cond.X[1] <= -2*st.Length || cond.X[2] <= -3*st.Length || cond.X[1] <= cond.X[0]-st.Length || cond.X[2] <= cond.X[1]-st.Length {
		st.CurState = NonPhis
	}
	switch st.CurState {
	case Started:
		fmt.Printf("HitWall: %v sec\n", st.CurTime)
		st.CurState = HitWall
	case HitWall:
		if cond.X[0] > 0 {
			fmt.Printf("BouncedBack: %v sec\n", st.CurTime)
			st.CurState = BouncedBack
		}
	case BouncedBack:
		if cond.X[0] <= 0. {
			fmt.Printf("HitWall: %v sec\n", st.CurTime)
			st.CurState = HitWall
		}
	case NonPhis:
		return errors.New("NonPhis conditions reached")
	}
	return nil
}

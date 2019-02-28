package grn

//State - finite automat states
type State int

//States - enumerable
const (
	Started State = iota
	HitTheWall
	BouncedBack
)

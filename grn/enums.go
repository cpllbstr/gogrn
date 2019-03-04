package grn

//State - finite automat states
type StateEnum int

//States - enumerable
const (
	Started StateEnum = iota
	HitWall
	BouncedBack
	NonPhis
)

package controller

type pos struct {
	X, Y int
}

type Describer interface {
	Walkable() bool
	Type() string
}

type Controller struct{
	*Player
	*Grid
}

type Input int

const (
	DownPressed Input = iota + 1
	UpPressed
)


func (c Controller) EvaluateInput(inputs ... Input) interface{} {

}



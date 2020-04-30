package controller

var _ Describer = &Nothing{}

type Nothing struct{}

func (w Nothing) Type() string {
	return "nothing"
}

func (w Nothing) Walkable() bool {
	return true
}

var theNothing = &Nothing{}
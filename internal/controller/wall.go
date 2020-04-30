package controller

var _ Describer = &Wall{}

type Wall struct{}

func (w Wall) Type() string {
	return "wall"
}

func (w Wall) Walkable() bool {
	return false
}

var theWall = &Wall{}
package controller

var _ Describer = &Player{}

type Player struct {
	pos
}

func (p Player) Walkable() bool {
	return false
}

func (p Player) Type() string {
	return "player"
}


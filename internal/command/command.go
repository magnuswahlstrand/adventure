package command

import (
	"fmt"
	"github.com/kyeett/single-player-game/internal/comp"
)

type CommandType string

type Command struct {
	Execute func() error
	Undo    func()
	Name    string
	Step    int64
}

type Movable interface {
	GetEntity() comp.Entity
	GetPosition() *comp.Position
}


func Move(unit Movable, target *comp.Position) *Command {
	pos := unit.GetPosition()
	x0, y0 := pos.X, pos.Y

	execute := func() error {
		pos.X = target.X
		pos.Y = target.Y
		return nil
	}
	undo := func() {
		pos.X = x0
		pos.Y = y0
	}
	return &Command{execute, undo, fmt.Sprintf("MoveTo(%d,%d)", target.X, target.Y), -1}
}

func ChangeHitpoints(hp *comp.Hitpoints, change int64) *Command {
	execute := func() error {
		hp.Amount = hp.Amount + change
		return nil
	}
	undo := func() {
		hp.Amount = hp.Amount - change
	}
	return &Command{execute, undo, "ChangeHitpoints", -1}
}
package command

import (
	"fmt"
	"github.com/kyeett/single-player-game/internal/comp"
)

type CommandType string

const (
	TypeMoveBy = "Move"
)

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

func MoveBy(unit Movable, dx, dy int) *Command {
	pos := unit.GetPosition()
	x0, y0 := pos.X, pos.Y

	execute := func() error {
		pos.X = x0 + dx
		pos.Y = y0 + dy
		return nil
	}
	undo := func() {
		fmt.Printf("Move back to ")
		pos.X = x0
		pos.Y = y0
	}
	return &Command{execute, undo, fmt.Sprintf("MoveBy(%d,%d)", dx, dy), -1}
}

type MoveTo struct {
	ActorID comp.ID
	Target  *comp.Position
}

func MoveBy2(unit Movable, dx, dy int) MoveTo {
	pos := unit.GetPosition()
	x0, y0 := pos.X, pos.Y

	return MoveTo{
		ActorID: unit.GetEntity().ID,
		Target:  comp.P(x0+dx, y0+dy),
	}
}

func MoveToCommand(unit Movable, target *comp.Position) *Command {
	pos := unit.GetPosition()
	x0, y0 := pos.X, pos.Y

	execute := func() error {
		pos.X = target.X
		pos.Y = target.Y
		return nil
	}
	undo := func() {
		fmt.Printf("Move back to ")
		pos.X = x0
		pos.Y = y0
	}
	return &Command{execute, undo, fmt.Sprintf("MoveTo(%d,%d)", target.X, target.Y), -1}
}

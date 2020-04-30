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
}

type Movable interface {
	GetPosition() *comp.Position
}

//func Move(unit Movable, x, y int) *Command {
//	pos := unit.GetPosition()
//	x0, y0 := pos.X, pos.Y
//
//	execute := func() {
//		pos.X = x
//		pos.Y = y
//	}
//	undo := func() {
//		pos.X = x0
//		pos.Y = y0
//	}
//	return &Command{execute, undo}
//}

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
	return &Command{execute, undo, fmt.Sprintf("MoveBy(%d,%d)", dx, dy)}
}

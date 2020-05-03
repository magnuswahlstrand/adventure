package command

import (
	"fmt"
	"github.com/kyeett/adventure/internal/comp"
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

func AddToInventory(inventory *comp.Inventory, item comp.Inventorable) *Command {
	execute := func() error {
		inventory.Add(item)
		return nil
	}
	undo := func() {
		inventory.Remove(item)
	}
	return &Command{execute, undo, "AddToInventory", -1}
}

func RemoveFromInventory(inventory *comp.Inventory, item comp.Inventorable) *Command {
	execute := func() error {
		inventory.Remove(item)
		return nil
	}
	undo := func() {
		inventory.Add(item)
	}
	return &Command{execute, undo, "RemoveFromInventory", -1}
}

package entitymanager

import (
	"fmt"
	"github.com/kyeett/single-player-game/internal/command"
	"github.com/kyeett/single-player-game/internal/comp"
)

type EntityLifeCycler interface {
	Add(v interface{})
	Remove(id comp.ID)
	FindEntityByID(id comp.ID) interface{}
}


func RemoveCommand(lifeCycler EntityLifeCycler, id comp.ID) *command.Command {
	e := lifeCycler.FindEntityByID(id)
	execute := func() error {
		lifeCycler.Remove(id)
		return nil
	}
	undo := func() {
		lifeCycler.Add(e)
	}
	return &command.Command{execute, undo, fmt.Sprintf("Remove entity")}
}

func AddCommand(lifeCycler EntityLifeCycler, id comp.ID, target interface{}) *command.Command {
	//e := lifeCycler.FindEntityByID(id)
	execute := func() error {
		lifeCycler.Add(target)
		return nil
	}
	undo := func() {
		lifeCycler.Remove(id)
	}
	return &command.Command{execute, undo, fmt.Sprintf("Add entity")}
}

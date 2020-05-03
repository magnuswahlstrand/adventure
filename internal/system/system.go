package system

import (
	"github.com/kyeett/adventure/internal/command"
	"github.com/kyeett/adventure/internal/comp"
	"github.com/kyeett/adventure/internal/event"
)

// A RenderSystem is updated every iteration, and draws to a screen
type System interface {
	Update(evt event.Event) []*command.Command

	Add(v interface{})
	Remove(id comp.ID)
}

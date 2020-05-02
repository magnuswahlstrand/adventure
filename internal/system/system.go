package system

import (
	"github.com/kyeett/single-player-game/internal/command"
	"github.com/kyeett/single-player-game/internal/comp"
	"github.com/kyeett/single-player-game/internal/event"
)

// A RenderSystem is updated every iteration, and draws to a screen
type System interface {
	Update(evt event.Event) []*command.Command

	Add(v interface{})
	Remove(id comp.ID)
}

package death

import (
	"github.com/kyeett/adventure/internal/command"
	"github.com/kyeett/adventure/internal/entitymanager"
	"github.com/kyeett/adventure/internal/event"
)

func (s *Death) Update(_ event.Event) []*command.Command{
	var commands []*command.Command

	// Find entities that have 0 or negative hit points
	for id, e := range s.entities {
		if e.Hitpoints.Amount <= 0 {
			commands = append(commands, entitymanager.RemoveEntity(s.lifeCycler, id))
		}
	}
	return commands
}

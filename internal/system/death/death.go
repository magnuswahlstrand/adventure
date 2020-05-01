package death

import (
	"github.com/kyeett/single-player-game/internal/command"
	"github.com/kyeett/single-player-game/internal/entitymanager"
)

func (s *Death) Update() []*command.Command{
	var commands []*command.Command

	// Find entities that have 0 or negative hit points
	for id, e := range s.entities {
		if e.Hitpoints.Amount <= 0 {
			commands = append(commands, entitymanager.RemoveCommand(s.lifeCycler, id))
		}
	}
	return commands
}

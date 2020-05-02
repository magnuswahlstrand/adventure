package base

import (
	"github.com/kyeett/single-player-game/internal/command"
	"github.com/kyeett/single-player-game/internal/comp"
	"github.com/kyeett/single-player-game/internal/entitymanager"
	"github.com/kyeett/single-player-game/internal/unit"
)

func (s *Base) move(aID comp.ID, position comp.Position) []*command.Command {
	var commands []*command.Command
	player := s.entities[aID]

	commands = append(commands, command.Move(player, &position))
	return commands
}

func (s *Base) takeItem(aID, itemID comp.ID) []*command.Command {
	var commands []*command.Command

	player := s.entities[aID]
	item := s.entities[itemID]

	// Remove item
	commands = append(commands, entitymanager.RemoveEntity(s.lifeCycler, item.ID))

	// Take position of item
	commands = append(commands, command.Move(player, item.Position))
	return commands
}

func (s *Base) openChest(_, chestID comp.ID) []*command.Command {
	var commands []*command.Command

	chest := s.entities[chestID]

	// Create item from what is inside the chest
	item := unit.NewItem(chest.Position.X, chest.Position.Y)

	// Remove chest
	commands = append(commands, entitymanager.RemoveEntity(s.lifeCycler, chest.ID))

	// Add item
	commands = append(commands, entitymanager.AddEntity(s.lifeCycler, item.ID, item))
	return commands
}
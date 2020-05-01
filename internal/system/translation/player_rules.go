package translation

import (
	"fmt"
	"github.com/kyeett/single-player-game/internal/command"
	"github.com/kyeett/single-player-game/internal/comp"
	"github.com/kyeett/single-player-game/internal/entitymanager"
	"github.com/kyeett/single-player-game/internal/unit"
)

func (s *Translation) playerInteract(actor TranslatableEntity, target TranslatableEntity) []*command.Command {
	s.logger.Debug(fmt.Sprintf("player interact with %v", target))

	switch target.Type {
	case comp.TypeEnemy:
		return s.playerWithEnemy(actor, target)
	case comp.TypeChest:
		return s.playerWithChest(actor, target)
	case comp.TypeItem:
		return s.playerWithItem(actor, target)
	}
	return nil
}

func (s *Translation) playerWithItem(player TranslatableEntity, item TranslatableEntity) []*command.Command {
	var commands []*command.Command

	// Remove item
	commands = append(commands, entitymanager.RemoveCommand(s.lifeCycler, item.ID))
	return commands
}

func (s *Translation) playerWithChest(player TranslatableEntity, chest TranslatableEntity) []*command.Command {
	var commands []*command.Command

	// Create item from what is inside the chest
	item := unit.NewItem(chest.Position.X, chest.Position.Y)

	// Remove chest
	commands = append(commands, entitymanager.RemoveCommand(s.lifeCycler, chest.ID))

	// Add item
	commands = append(commands, entitymanager.AddCommand(s.lifeCycler, item.ID, item))
	return commands
}

func (s *Translation) playerWithEnemy(player TranslatableEntity, enemy TranslatableEntity) []*command.Command {
	var commands []*command.Command

	// "Kill" player
	//commands = append(commands, command.MoveToCommand(player, comp.P(0, 0)))
	return commands
}

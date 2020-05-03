package base

import (
	"github.com/kyeett/single-player-game/internal/command"
	"github.com/kyeett/single-player-game/internal/comp"
	"github.com/kyeett/single-player-game/internal/entitymanager"
	"github.com/kyeett/single-player-game/internal/unit"
)

func (s *Base) takeItem(aID, itemID comp.ID) []*command.Command {
	var commands []*command.Command

	player := s.entities[aID]
	//item := s.entities[itemID]

	// TODO: Find nicer way of looking up entities
	item, ok := s.lifeCycler.FindEntityByID(itemID).(*unit.Item)
	if !ok {
		return nil
	}

	// Remove item
	commands = append(commands, entitymanager.RemoveEntity(s.lifeCycler, item.ID))

	// Take item
	commands = append(commands, command.AddToInventory(s.player.Inventory, item))

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

func (s *Base) openDoor(aID, chestID comp.ID) []*command.Command {
	var commands []*command.Command

	door := s.entities[chestID]

	key := s.player.Inventory.GetType(comp.TypeItem)
	if key == nil {
		return nil
	}

	// Use key
	commands = append(commands, command.RemoveFromInventory(s.player.Inventory, key))

	// Remove door
	commands = append(commands, entitymanager.RemoveEntity(s.lifeCycler, door.ID))
	return commands
}
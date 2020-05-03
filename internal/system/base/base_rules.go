package base

import (
	"github.com/kyeett/adventure/internal/command"
	"github.com/kyeett/adventure/internal/comp"
	"github.com/kyeett/adventure/internal/entitymanager"
	"github.com/kyeett/adventure/internal/unit"
)

func (s *Base) takeItem(aID, itemID comp.ID) []*command.Command {
	var commands []*command.Command

	player := s.players[aID]

	// TODO: Find nicer way of looking up entities
	item, ok := s.lifeCycler.FindEntityByID(itemID).(*unit.Item)
	if !ok {
		return nil
	}

	// Remove item
	commands = append(commands, entitymanager.RemoveEntity(s.lifeCycler, item.ID))

	// Take item
	commands = append(commands, command.AddToInventory(player.Inventory, item))

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
	player := s.players[aID]

	key := player.Inventory.GetType(comp.TypeItem)
	if key == nil {
		return nil
	}

	// Use key
	commands = append(commands, command.RemoveFromInventory(player.Inventory, key))

	// Remove door
	commands = append(commands, entitymanager.RemoveEntity(s.lifeCycler, door.ID))
	return commands
}
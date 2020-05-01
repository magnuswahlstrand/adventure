package system

import (
	"fmt"
	"github.com/kyeett/single-player-game/internal/command"
	"github.com/kyeett/single-player-game/internal/comp"
	"github.com/kyeett/single-player-game/internal/unit"
)

func (s *Translation) playerInteract(e *unit.Player, target Translatable) []*command.Command {
	s.logger.Info(fmt.Sprintf("player interact with %v", target))
	switch v := target.(type) {
	case *unit.Enemy:
		return s.playerWithEnemy(e, v)
	case *unit.Chest:
		return s.playerWithChest(e, v)
	case *unit.Item:
		return s.playerWithItem(e, v)
	}
	return nil
}

func (s *Translation) playerWithItem(p *unit.Player, item *unit.Item) []*command.Command {
	return []*command.Command{
		s.RemoveCommand(item),
	}
}

func (s *Translation) playerWithChest(p *unit.Player, chest *unit.Chest) []*command.Command {
	item := unit.NewItem(chest.Position.X, chest.Position.Y)
	return []*command.Command{
		s.RemoveCommand(chest),
		s.AddCommand(item),
	}
}

func (s *Translation) playerWithEnemy(p *unit.Player, enemy *unit.Enemy) []*command.Command {
	return []*command.Command{
		command.MoveToCommand(p, comp.P(0, 0)),
	}
}

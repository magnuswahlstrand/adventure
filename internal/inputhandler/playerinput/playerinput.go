package playerinput

import (
	"errors"
	"fmt"
	"github.com/kyeett/adventure/internal/comp"
	"github.com/kyeett/adventure/internal/event"
)

var (
	TheWall = TranslatableEntity{Entity: comp.Entity{Type: comp.TypeWall}}
	TheNil  = TranslatableEntity{Entity: comp.Entity{Type: comp.TypeNil}}
)

func (s *Translation) findAtPosition(p *comp.Position) TranslatableEntity {
	if p.X >= 6 || p.X < 0 {
		return TheWall
	}

	if p.Y >= 6 || p.Y < 0 {
		return TheWall
	}

	if p.Y == 2 && p.X != 2 {
		return TheWall
	}

	for _, e := range s.entities {
		if p.Equals(e.GetPosition()) {
			return e
		}
	}

	return TheNil
}

func (s *Translation) findByID(ID comp.ID) (TranslatableEntity, error) {
	e, found := s.entities[ID]
	if !found {
		return TranslatableEntity{}, errors.New("not found")
	}

	return e, nil
}

func (s *Translation) playerInteract(target TranslatableEntity) event.Event {
	s.logger.Debug(fmt.Sprintf("player interact with %v", target))
	switch target.Type {
	case comp.TypeEnemy:
		return event.Attack{
			Actor:  s.player.ID,
			Target: target.ID,
		}
	case comp.TypeChest:
		return event.OpenChest{
			Actor:  s.player.ID,
			Target: target.ID,
		}
	case comp.TypeItem:
		return event.TakeItem{
			Actor:  s.player.ID,
			Target: target.ID,
		}
	case comp.TypeDoor:
		return event.OpenDoor{
			Actor:  s.player.ID,
			Target: target.ID,
		}
	case comp.TypeGoal:
		return event.ReachGoal{}
	}
	return nil
}

package translation

import (
	"errors"
	"fmt"
	"github.com/kyeett/single-player-game/internal/command"
	"github.com/kyeett/single-player-game/internal/comp"
	"go.uber.org/zap"
	"log"
)

type EntityLifeCycler interface {
	Add(v interface{})
	Remove(id comp.ID)
}

var (
	TheWall = TranslatableEntity{Entity: comp.Entity{Type: comp.TypeWall}}
	TheNil  = TranslatableEntity{Entity: comp.Entity{Type: comp.TypeNil}}
)

func (s *Translation) Translate(c interface{}) []*command.Command {
	if c == nil {
		return nil
	}

	switch v := c.(type) {
	case command.MoveTo:
		actor, err := s.findByID(v.ActorID)
		if err != nil {
			s.logger.Debug("didn't find actor", zap.String("ID", string(v.ActorID)))
			return nil
		}

		target := s.findAtPosition(v.Target)
		switch target.Type {
		case comp.TypeEnemy,
			comp.TypeChest,
			comp.TypeItem:
			return s.interact(actor, target)
		case comp.TypeNil:
			return []*command.Command{command.MoveToCommand(actor, v.Target)}
		default:
			fmt.Printf("%T\n", target)
			return []*command.Command{}
		}
	default:
		log.Fatal("translate:", v)
	}
	return nil
}

func (s *Translation) findAtPosition(p *comp.Position) TranslatableEntity {
	if p.X >= 6 || p.X < 0 {
		return TheWall
	}

	if p.Y >= 6 || p.Y < 0 {
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

func (s *Translation) interact(actor TranslatableEntity, target TranslatableEntity) []*command.Command {
	switch actor.Type {
	case comp.TypePlayer:
		return s.playerInteract(actor, target)
	}

	return nil
}

func (s *Translation) RemoveCommand(id comp.ID, target interface{}) *command.Command {
	execute := func() error {
		s.logger.Info(fmt.Sprintf("remove %v", target))
		s.lifeCycler.Remove(id)
		return nil
	}
	undo := func() {
		s.logger.Info(fmt.Sprintf("readd %v", target))
		s.lifeCycler.Add(target)
	}
	return &command.Command{execute, undo, fmt.Sprintf("Remove entity")}
}

func (s *Translation) AddCommand(id comp.ID, target interface{}) *command.Command {
	execute := func() error {
		s.logger.Info(fmt.Sprintf( "Add %v", target))
		s.lifeCycler.Add(target)
		return nil
	}
	undo := func() {
		s.logger.Info(fmt.Sprintf( "Remove %v", target))
		s.lifeCycler.Remove(id)
	}
	return &command.Command{execute, undo, fmt.Sprintf("Add entity")}
}

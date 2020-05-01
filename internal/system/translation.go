package system

import (
	"fmt"
	"github.com/kyeett/single-player-game/internal/command"
	"github.com/kyeett/single-player-game/internal/comp"
	"github.com/kyeett/single-player-game/internal/unit"
	"log"
)

type EntityLifeCycler interface {
	Add(v interface{})
	Remove(id comp.ID)
}

const (
	translationSystem = "translationsystem"
)

func (s *Translation) target(p *comp.Position) TranslatableEntity {
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

type Wall struct{
	TranslatableEntity
}
type Nil struct{
	TranslatableEntity
}
var TheWall *Wall
var TheNil *Nil

func (s *Translation) Translate(c interface{}) []*command.Command {
	if c == nil {
		return nil
	}

	switch v := c.(type) {
	case command.MoveTo:
		target := s.target(v.Target)
		switch v2 := target.(type) {
		case *unit.Enemy,
			*unit.Chest,
			*unit.Item:
			return s.interact(v.Actor, v2)
		case *Nil:
			return []*command.Command{command.MoveToCommand(v.Actor, v.Target)}
		default:
			fmt.Printf("%T\n", v2)
			return []*command.Command{}
		}
	default:
		log.Fatal("translate:", v)
	}
	return nil
}

func (s *Translation) interact(actor Translatable, target Translatable) []*command.Command {
	switch v := actor.(type) {
	case *unit.Player:
		return s.playerInteract(v, target)
	}

	return nil
}

func (s *Translation) RemoveCommand(target Translatable) *command.Command {
	execute := func() error {
		s.logger.Info(fmt.Sprintf("remove %v", target))
		s.lifeCycler.Remove(target.GetEntity().GetID())
		return nil
	}
	undo := func() {
		s.logger.Info(fmt.Sprintf("readd %v", target))
		s.lifeCycler.Add(target)
	}
	return &command.Command{execute, undo, fmt.Sprintf("Remove entity")}
}

func (s *Translation) AddCommand(e Translatable) *command.Command {
	execute := func() error {
		s.logger.Info(fmt.Sprintf("Add %v", e))
		s.lifeCycler.Add(e)
		return nil
	}
	undo := func() {
		s.logger.Info(fmt.Sprintf("Remove %v", e))
		s.lifeCycler.Remove(e.GetEntity().ID)
	}
	return &command.Command{execute, undo, fmt.Sprintf("Add entity")}
}

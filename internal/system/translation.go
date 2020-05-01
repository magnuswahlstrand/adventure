package system

import (
	"fmt"
	"github.com/kyeett/single-player-game/internal/command"
	"github.com/kyeett/single-player-game/internal/comp"
	"github.com/kyeett/single-player-game/internal/unit"
	"go.uber.org/zap"
	"log"
)

var _ System = &Translation{}

type Translation struct {
	entities   []SomeInterface
	logger     *zap.Logger
	lifeCycler EntityLifeCycler
}

type EntityLifeCycler interface {
	Add(v interface{})
	Remove(v interface{})
}

func NewTranslation(logger *zap.Logger, lifeCycler EntityLifeCycler) *Translation {
	return &Translation{
		entities:   []SomeInterface{},
		logger:     logger,
		lifeCycler: lifeCycler,
	}
}

type SomeInterface interface {
	GetEntity() comp.Entity
	GetPosition() *comp.Position
}

const (
	translationSystem = "translationsystem"
)

func (t *Translation) Add(v interface{}) {
	e, ok := v.(SomeInterface)
	if !ok {
		return
	}
	t.logger.Info("add entity to " + translationSystem)
	t.entities = append(t.entities, e)
}

func (t *Translation) Remove(v interface{}) {
	e, ok := v.(comp.Removable)
	if !ok {
		return
	}
	id := e.GetID()
	for i, e := range t.entities {
		if e.GetEntity().ID == id {
			t.entities = append(t.entities[:i], t.entities[i+1:]...)
			return
		}
	}
}

func (t *Translation) target(p *comp.Position) comp.Entity {
	if p.X >= 6 || p.X < 0 {
		return comp.BaseWall
	}

	if p.Y >= 6 || p.Y < 0 {
		return comp.BaseWall
	}

	for _, e := range t.entities {
		if p.Equals(e.GetPosition()) {
			return e.GetEntity()
		}
	}

	return comp.BaseNil
}

func (t *Translation) lookup(en comp.Entity) SomeInterface {
	for _, e := range t.entities {
		if e.GetEntity() == en {
			return e
		}
	}

	return nil
}

func (t *Translation) Translate(c interface{}) []*command.Command {
	if c == nil {
		return nil
	}

	switch v := c.(type) {
	case command.MoveTo:

		target := t.target(v.Target)
		switch target.GetType() {
		case comp.TypeEnemy,
			comp.TypeChest,
			comp.TypeItem:
			return t.interact(v.Actor, target)
		case comp.TypeNil:
			return []*command.Command{command.MoveToCommand(v.Actor, v.Target)}
		default:
			return []*command.Command{}
		}
	default:
		log.Fatal("translate:", v)
	}
	return nil
}

func (t *Translation) interact(e command.Movable, target comp.Entity) []*command.Command {
	v := t.lookup(target)
	t.logger.Info(fmt.Sprintf("interact with %v", v))
	switch e.GetEntity().Type {
	case comp.TypePlayer:
		return t.playerInteract(e, v)
	}

	return nil
}

func (t *Translation) playerInteract(e command.Movable, target SomeInterface) []*command.Command {
	t.logger.Info(fmt.Sprintf("player interact with %v", target))
	switch v := target.(type) {
	case *unit.Enemy:
		t.logger.Info("enemy")
		return []*command.Command{
			command.MoveToCommand(e, comp.P(0, 0)),
		}
	case *unit.Chest:
		t.logger.Info("chest")

		item := unit.NewItem(v.Position.X, v.Position.Y)
		return []*command.Command{
			t.RemoveCommand(target),
			t.AddCommand(item),
		}
	case *unit.Item:
		t.logger.Info("item")
		return []*command.Command{
			t.RemoveCommand(target),
		}
	}

	return nil
}

func (t *Translation) RemoveCommand(target SomeInterface) *command.Command {
	execute := func() error {
		t.logger.Info(fmt.Sprintf("remove %v", target))
		t.lifeCycler.Remove(target)
		return nil
	}
	undo := func() {
		t.logger.Info(fmt.Sprintf("readd %v", target))
		t.lifeCycler.Add(target)
	}
	return &command.Command{execute, undo, fmt.Sprintf("Remove entity")}
}

func (t *Translation) AddCommand(e SomeInterface) *command.Command {
	execute := func() error {
		t.logger.Info(fmt.Sprintf("Add %v", e))
		t.lifeCycler.Add(e)
		return nil
	}
	undo := func() {
		t.logger.Info(fmt.Sprintf("Remove %v", e))
		t.lifeCycler.Remove(e)
	}
	return &command.Command{execute, undo, fmt.Sprintf("Add entity")}
}

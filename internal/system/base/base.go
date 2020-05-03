package base

import (
	"github.com/kyeett/adventure/internal/command"
	"github.com/kyeett/adventure/internal/comp"
	"github.com/kyeett/adventure/internal/entitymanager"
	"github.com/kyeett/adventure/internal/event"
	"github.com/kyeett/adventure/internal/logger"
	"github.com/kyeett/adventure/internal/system"
	"github.com/kyeett/adventure/internal/unit"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ system.System = &Base{}

type Base struct {
	entities   map[comp.ID]BaseEntity
	players    map[comp.ID]*unit.Player
	logger     *zap.SugaredLogger
	lifeCycler entitymanager.EntityLifeCycler
}

func NewSystem(logLevel zapcore.Level, lifeCycler entitymanager.EntityLifeCycler) *Base {
	return &Base{
		entities:   map[comp.ID]BaseEntity{},
		players:    map[comp.ID]*unit.Player{},
		logger:     logger.NewNamed("base", logLevel, logger.Green),
		lifeCycler: lifeCycler,
	}
}

type Basable interface {
	GetEntity() comp.Entity
	GetPosition() *comp.Position
}

type BaseEntity struct {
	comp.Entity
	*comp.Position
	source interface{}
}

func (s *BaseEntity) GetSource() interface{} {
	return s.source
}

func (s *Base) Update(evt event.Event) []*command.Command {
	switch v := evt.(type) {
	case event.TakeItem:
		return s.takeItem(v.Actor, v.Target)
	case event.OpenChest:
		return s.openChest(v.Actor, v.Target)
	case event.OpenDoor:
		return s.openDoor(v.Actor, v.Target)
	}

	return nil
}

func (s *Base) Add(v interface{}) {
	i, ok := v.(Basable)
	if !ok {
		return
	}
	e := BaseEntity{
		Entity:   i.GetEntity(),
		Position: i.GetPosition(),
		source:   v,
	}
	s.logger.Info("entity added")
	s.entities[e.ID] = e

	// Add potential player to player list
	p, ok := v.(*unit.Player)
	if !ok {
		return
	}
	s.players[p.ID] = p
}

func (s *Base) Remove(id comp.ID) {
	delete(s.entities, id)
}

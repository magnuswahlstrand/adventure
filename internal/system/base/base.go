package base

import (
	"github.com/kyeett/single-player-game/internal/command"
	"github.com/kyeett/single-player-game/internal/comp"
	"github.com/kyeett/single-player-game/internal/entitymanager"
	"github.com/kyeett/single-player-game/internal/event"
	"github.com/kyeett/single-player-game/internal/logger"
	"github.com/kyeett/single-player-game/internal/system"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ system.System = &Base{}

type Base struct {
	entities   map[comp.ID]BaseEntity
	logger     *zap.SugaredLogger
	lifeCycler entitymanager.EntityLifeCycler
}

func NewSystem(logLevel zapcore.Level, lifeCycler entitymanager.EntityLifeCycler) *Base {
	return &Base{
		entities:   map[comp.ID]BaseEntity{},
		logger:   logger.NewNamed("base", logLevel, logger.Green),
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
	switch v := evt.(type)  {
	case event.Move:
		return s.move(v.Actor, v.Position)
	case event.Attack:
		return s.attack(v.Actor, v.Target)
	case event.TakeItem:
		return s.takeItem(v.Actor, v.Target)
	case event.OpenChest:
		return s.openChest(v.Actor, v.Target)
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
		source: v,
	}
	s.logger.Info("entity added")
	s.entities[e.ID] = e
}

func (s *Base) Remove(id comp.ID) {
	delete(s.entities, id)
}
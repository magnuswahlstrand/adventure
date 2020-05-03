package attack

import (
	"github.com/kyeett/adventure/internal/command"
	"github.com/kyeett/adventure/internal/comp"
	"github.com/kyeett/adventure/internal/entitymanager"
	"github.com/kyeett/adventure/internal/event"
	"github.com/kyeett/adventure/internal/logger"
	"github.com/kyeett/adventure/internal/system"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ system.System = &Attack{}

type Attack struct {
	entities   map[comp.ID]BaseEntity
	logger     *zap.SugaredLogger
	lifeCycler entitymanager.EntityLifeCycler
}

func NewSystem(logLevel zapcore.Level, lifeCycler entitymanager.EntityLifeCycler) *Attack {
	return &Attack{
		entities:   map[comp.ID]BaseEntity{},
		logger:   logger.NewNamed("attack", logLevel, logger.Red),
		lifeCycler: lifeCycler,
	}
}

type Attackable interface {
	GetEntity() comp.Entity
	GetHitpoints() *comp.Hitpoints
}

type BaseEntity struct {
	comp.Entity
	*comp.Hitpoints
}

func (s *Attack) Update(evt event.Event) []*command.Command {
	switch v := evt.(type)  {
	case event.Attack:
		return s.attack(v.Actor, v.Target)
	}

	return nil
}

func (s *Attack) Add(v interface{}) {
	i, ok := v.(Attackable)
	if !ok {
		return
	}
	e := BaseEntity{
		Entity:   i.GetEntity(),
		Hitpoints: i.GetHitpoints(),
	}
	s.logger.Info("entity added")
	s.entities[e.ID] = e
}

func (s *Attack) Remove(id comp.ID) {
	delete(s.entities, id)
}

func (s *Attack) attack(_, tID comp.ID) []*command.Command {
	var commands []*command.Command
	target := s.entities[tID]
	damage := -int64(1)
	commands = append(commands, command.ChangeHitpoints(target.Hitpoints, damage))
	return commands
}

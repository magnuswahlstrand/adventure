package movement

import (
	"github.com/kyeett/adventure/internal/command"
	"github.com/kyeett/adventure/internal/comp"
	"github.com/kyeett/adventure/internal/event"
	"github.com/kyeett/adventure/internal/logger"
	"github.com/kyeett/adventure/internal/system"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ system.System = &Base{}

type Base struct {
	entities   map[comp.ID]BaseEntity
	logger     *zap.SugaredLogger
}

func NewSystem(logLevel zapcore.Level) *Base {
	return &Base{
		entities:   map[comp.ID]BaseEntity{},
		logger:   logger.NewNamed("move", logLevel, logger.BrightGreen),
	}
}

type Basable interface {
	GetEntity() comp.Entity
	GetPosition() *comp.Position
}

type BaseEntity struct {
	comp.Entity
	*comp.Position
}

func (s *Base) Update(evt event.Event) []*command.Command {
	switch v := evt.(type)  {
	case event.Move:
		return s.move(v.Actor, v.Position)
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
	}
	s.logger.Info("entity added")
	s.entities[e.ID] = e
}

func (s *Base) Remove(id comp.ID) {
	delete(s.entities, id)
}

func (s *Base) move(aID comp.ID, position comp.Position) []*command.Command {
	var commands []*command.Command
	player := s.entities[aID]

	commands = append(commands, command.Move(player, &position))
	return commands
}
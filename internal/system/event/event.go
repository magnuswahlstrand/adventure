package event

import (
	"github.com/kyeett/adventure/internal/command"
	"github.com/kyeett/adventure/internal/comp"
	"github.com/kyeett/adventure/internal/event"
	"github.com/kyeett/adventure/internal/logger"
	"github.com/kyeett/adventure/internal/system"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ system.System = &Event{}

type Event struct {
	logger     *zap.SugaredLogger
	gamestate *string
}

func NewSystem(logLevel zapcore.Level, gamestate *string) *Event {
	return &Event{
		logger:   logger.NewNamed("event", logLevel, logger.Blue),
		gamestate: gamestate,
	}
}

func (s *Event) Update(evt event.Event) []*command.Command {
	if evt != nil {
		s.logger.Infof("READ EVENT %v, %T", evt, evt)
	}

	switch evt.(type)  {
	case event.ReachGoal:
		state := "finished"
		*s.gamestate = state
		s.logger.Infof("set game state to finished %v", s.gamestate)
	}

	return nil
}

func (s *Event) Add(_ interface{}) {}

func (s *Event) Remove(_ comp.ID) {}
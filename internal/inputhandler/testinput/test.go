package testinput

import (
	"github.com/kyeett/adventure/internal/event"
	"github.com/kyeett/adventure/internal/inputhandler"
	"github.com/kyeett/adventure/internal/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ inputhandler.InputHandler = &TestInput{}

type TestInput struct {
	events []event.Event
	logger *zap.SugaredLogger
}

func New(logLevel zapcore.Level) *TestInput {
	return &TestInput{
		events: []event.Event{},
		logger: logger.NewNamed("test", logLevel, logger.Yellow),
	}
}

func (s *TestInput) AddEvent(evt event.Event) {
	s.events = append(s.events, evt)
}

func (s *TestInput) GetEvent() event.Event {
	n := len(s.events)
	if n == 0 {
		return nil
	}

	evt := s.events[n-1]
	s.events = s.events[:n-1]
	return evt
}

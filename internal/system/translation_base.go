package system

import (
	"github.com/kyeett/single-player-game/internal/comp"
	"go.uber.org/zap"
)

var _ System = &Translation{}

type Translation struct {
	entities map[comp.ID]TranslatableEntity
	logger     *zap.Logger
	lifeCycler EntityLifeCycler
}

func NewTranslation(logger *zap.Logger, lifeCycler EntityLifeCycler) *Translation {
	return &Translation{
		entities:   map[comp.ID]TranslatableEntity{},
		logger:     logger,
		lifeCycler: lifeCycler,
	}
}

type Translatable interface {
	GetEntity() comp.Entity
	GetPosition() *comp.Position
}

type TranslatableEntity struct {
	comp.Entity
	*comp.Position
}

func (s *Translation) Add(v interface{}) {
	i, ok := v.(Translatable)
	if !ok {
		return
	}
	e := TranslatableEntity{
		Entity:   i.GetEntity(),
		Position: i.GetPosition(),
	}
	s.logger.Info("add entity to " + translationSystem)
	s.entities[e.ID] = e
}

func (s *Translation) Remove(id comp.ID) {
	delete(s.entities, id)
}

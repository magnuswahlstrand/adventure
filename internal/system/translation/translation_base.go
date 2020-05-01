package translation

import (
	"github.com/kyeett/single-player-game/internal/comp"
	"github.com/kyeett/single-player-game/internal/logger"
	"github.com/kyeett/single-player-game/internal/system"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ system.System = &Translation{}

type Translation struct {
	entities   map[comp.ID]TranslatableEntity
	logger     *zap.SugaredLogger
	lifeCycler EntityLifeCycler
}

func NewTranslation(logLevel zapcore.Level, lifeCycler EntityLifeCycler) *Translation {
	return &Translation{
		entities:   map[comp.ID]TranslatableEntity{},
		logger:   logger.NewNamed("sprite", logLevel),
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
	source interface{}
}

func (s *TranslatableEntity) GetSource() interface{} {
	return s.source
}

func (s *Translation) Update(_ float64) {}

func (s *Translation) Add(v interface{}) {
	i, ok := v.(Translatable)
	if !ok {
		return
	}
	e := TranslatableEntity{
		Entity:   i.GetEntity(),
		Position: i.GetPosition(),
		source: v,
	}
	s.logger.Info("entity added")
	s.entities[e.ID] = e
}

func (s *Translation) Remove(id comp.ID) {
	delete(s.entities, id)
}

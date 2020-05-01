package death

import (
	"github.com/kyeett/single-player-game/internal/comp"
	"github.com/kyeett/single-player-game/internal/logger"
	"github.com/kyeett/single-player-game/internal/system"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ system.System = &Death{}

type Death struct {
	entities   map[comp.ID]DeathC
	logger     *zap.SugaredLogger
	lifeCycler EntityLifeCycler
}

func NewSystem(logLevel zapcore.Level, lifeCycler EntityLifeCycler) *Death {
	return &Death{
		entities:   map[comp.ID]DeathC{},
		logger:     logger.NewNamed("death", logLevel, logger.BrightBlack),
		lifeCycler: lifeCycler,
	}
}

type deathable interface {
	GetEntity() comp.Entity
	GetHitpoints() *comp.Hitpoints
}

type DeathC struct {
	comp.Entity
	*comp.Hitpoints
}

func (s *Death) Add(v interface{}) {
	i, ok := v.(deathable)
	if !ok {
		return
	}
	e := DeathC{
		Entity:   i.GetEntity(),
		Hitpoints: i.GetHitpoints(),
	}
	s.logger.Info("entity added")
	s.entities[e.ID] = e
}

func (s *Death) Remove(id comp.ID) {
	delete(s.entities, id)
}

type EntityLifeCycler interface {
	Add(v interface{})
	Remove(id comp.ID)
}
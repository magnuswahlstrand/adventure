package rendersystem

import (
	"github.com/kyeett/single-player-game/internal/comp"
	"github.com/kyeett/single-player-game/internal/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ System = &Render{}

// Render is responsible for drawing entities to the screen
type Render struct {
	entities map[comp.ID]DrawableEntity
	logger   *zap.SugaredLogger
}

func NewRender(logLevel zapcore.Level) *Render {
	return &Render{
		entities: map[comp.ID]DrawableEntity{},
		logger:   logger.NewNamed("sprite", logLevel, logger.Red),
	}
}

type Drawable interface {
	GetEntity() comp.Entity
	GetPosition() *comp.Position
	GetSprite() *comp.Sprite
}

type DrawableEntity struct {
	comp.Entity
	*comp.Position
	*comp.Sprite
}

const (
	spriteSystem = "spritesystem"
)

func (s *Render) Add(v interface{}) {
	i, ok := v.(Drawable)
	if !ok {
		return
	}
	e := DrawableEntity{
		Entity:   i.GetEntity(),
		Position: i.GetPosition(),
		Sprite:   i.GetSprite(),
	}
	s.logger.Info("add entity to " + spriteSystem)
	s.entities[e.ID] = e
}

func (s *Render) Remove(id comp.ID) {
	delete(s.entities, id)
}

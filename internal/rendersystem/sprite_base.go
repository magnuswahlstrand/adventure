package rendersystem

import (
	"github.com/kyeett/single-player-game/internal/comp"
	"go.uber.org/zap"
)

var _ System = &Render{}

// Render is responsible for drawing entities to the screen
type Render struct {
	//em  *entity.Manager
	entities map[comp.ID]DrawableEntity
	logger   *zap.Logger
}

func NewRender(logger *zap.Logger) *Render {
	return &Render{
		entities: map[comp.ID]DrawableEntity{},
		logger:   logger,
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

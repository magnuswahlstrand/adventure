package rendersystem

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/single-player-game/internal/comp"
	"github.com/peterhellberg/gfx"
	"go.uber.org/zap"
)

var _ System = &Render{}

// Render is responsible for drawing entities to the screen
type Render struct {
	//em  *entity.Manager
	entities []Drawable
	logger   *zap.Logger
}

func (r *Render) Remove(v interface{}) {
	e, ok := v.(comp.Removable)
	if !ok {
		return
	}
	id := e.GetID()
	for i, e := range r.entities {
		if e.GetEntity().ID == id {
			r.entities = append(r.entities[:i], r.entities[i+1:]...)
			return
		}
	}
}

func NewRender(logger *zap.Logger) *Render{
	return   &Render{
		entities: []Drawable{},
		logger: logger,
	}
}

func scale(scale float64, position *comp.Position) gfx.Vec {
	return gfx.IV(position.X, position.Y).Scaled(scale)
}

func (r *Render) Draw(screen *ebiten.Image) {
	for _, e := range r.entities {
		opt := &ebiten.DrawImageOptions{}
		t := scale(8+1, e.GetPosition()).AddXY(1,1)
		opt.GeoM.Translate(t.X, t.Y)
		screen.DrawImage(e.GetSprite().Image, opt)
	}
}

type Drawable interface {
	GetEntity() comp.Entity
	GetPosition() *comp.Position
	GetSprite() *comp.Sprite
}

const (
	spriteSystem = "spritesystem"
)

func (r *Render) Add(v interface{}) {
	e, ok := v.(Drawable)
	if !ok {
		return
	}
	r.logger.Info("add entity to " + spriteSystem)
	r.entities = append(r.entities, e)
}




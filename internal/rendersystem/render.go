package rendersystem

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/single-player-game/internal/comp"
	"github.com/peterhellberg/gfx"
	"golang.org/x/image/colornames"
)

var _ System = &Render{}

// Render is responsible for drawing entities to the screen
type Render struct {
	//em  *entity.Manager
	entities []Drawable
}

func NewRender() *Render{
	return   &Render{
		entities: []Drawable{},
	}
}

func scale(scale float64, position *comp.Position) gfx.Vec {
	return gfx.IV(position.X, position.Y).Scaled(scale)
}

func (r *Render) Draw(screen *ebiten.Image) {
	for _, e := range r.entities {
		gfx.DrawCircleFilled(screen, scale(10, e.GetPosition()).Sub(gfx.V(10, 10)), 4, colornames.Blueviolet)
	}
}

type Drawable interface {
	GetPosition() *comp.Position
}

func (r *Render) Add(v interface{}) {
	e, ok := v.(Drawable)
	if !ok {
		return
	}
	r.entities = append(r.entities, e)
}




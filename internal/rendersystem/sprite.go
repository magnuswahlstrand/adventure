package rendersystem

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/single-player-game/internal/comp"
	"github.com/peterhellberg/gfx"
)


func scale(scale float64, position *comp.Position) gfx.Vec {
	return gfx.IV(position.X, position.Y).Scaled(scale)
}

func (s *Render) Draw(screen *ebiten.Image) {
	for _, e := range s.entities {
		opt := &ebiten.DrawImageOptions{}
		t := scale(8+1, e.Position).AddXY(1,1)
		opt.GeoM.Translate(t.X, t.Y)
		screen.DrawImage(e.Sprite.Image, opt)
	}
}


package rendersystem

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/single-player-game/internal/comp"
	"github.com/peterhellberg/gfx"
)


func scale(scale float64, position *comp.Position) gfx.Vec {
	return gfx.IV(position.X, position.Y).Scaled(scale)
}

const (
	borderSize = 2
)

func (s *Render) Draw(screen *ebiten.Image) {
	for _, e := range s.entities {
		opt := &ebiten.DrawImageOptions{}
		opt.GeoM.Scale(2,2)
		t := scale(16+borderSize, e.Position).AddXY(borderSize,borderSize)
		opt.GeoM.Translate(t.X, t.Y)
		screen.DrawImage(e.Sprite.Image, opt)
	}
}


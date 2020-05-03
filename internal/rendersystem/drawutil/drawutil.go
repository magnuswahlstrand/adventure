package drawutil

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/adventure/internal/comp"
	"github.com/peterhellberg/gfx"
)

const (
	gridSize   = 16
	borderSize = 1
)

func scale(scale float64, position *comp.Position) gfx.Vec {
	return gfx.IV(position.X, position.Y).Scaled(scale)
}


func DrawSprite(screen *ebiten.Image, position *comp.Position, sprite *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Scale(2, 2)
	t := scale(gridSize+borderSize, position).AddXY(borderSize, borderSize)
	opt.GeoM.Translate(t.X, t.Y)
	screen.DrawImage(sprite, opt)
}
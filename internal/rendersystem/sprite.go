package rendersystem

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/single-player-game/internal/rendersystem/drawutil"
)

func (s *Render) Draw(screen *ebiten.Image) {
	for _, e := range s.entities {
		drawutil.DrawSprite(screen, e.Position, e.Sprite.Image)
	}
}


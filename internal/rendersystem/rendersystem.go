package rendersystem

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/single-player-game/internal/comp"
)

// A RenderSystem is updated every iteration, and draws to a screen
type System interface {
	Draw(*ebiten.Image)

	Add(v interface{})
	Remove(id comp.ID)
}

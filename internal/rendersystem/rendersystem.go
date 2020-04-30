package rendersystem

import "github.com/hajimehoshi/ebiten"

// A RenderSystem is updated every iteration, and draws to a screen
type System interface {
	Draw(*ebiten.Image)
	Add(v interface{})
}

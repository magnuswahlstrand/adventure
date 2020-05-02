package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"golang.org/x/image/colornames"
)

const (
	gridSize = 16
	borderSize = 2
)

func (g *Game) drawGrid(screen *ebiten.Image) {
	for iy := float64(0); iy < 6; iy++ {
		for ix := float64(0); ix < 6; ix++ {
			ebitenutil.DrawRect(screen, (gridSize+borderSize)*ix+borderSize, (gridSize+borderSize)*iy+borderSize, gridSize, gridSize, colornames.Darkgoldenrod)
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "press 'z' to undo", 10, 2*70)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("step: %d", g.GameState.step), 2*60, 0)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("turn: %d", g.GameState.turn), 2*60, 15)

	// Draw stack
	for i, s := range g.events {
		ebitenutil.DebugPrintAt(screen, string(s.Type()), 2*60, (30+15*i))
	}
	g.drawGrid(screen)

	for _, rs := range g.rendersystems {
		rs.Draw(screen)
	}
}

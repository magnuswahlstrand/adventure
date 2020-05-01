package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"golang.org/x/image/colornames"
)

const (
	gridSize = 8
)

func (g *Game) drawGrid(screen *ebiten.Image) {
	for iy := float64(0); iy < 6; iy++ {
		for ix := float64(0); ix < 6; ix++ {
			ebitenutil.DrawRect(screen, (gridSize+1)*ix+1, (gridSize+1)*iy+1, gridSize, gridSize, colornames.Darkgoldenrod)
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "press 'z' to undo", 10		,70)
	g.drawGrid(screen)

	for _, rs := range g.rendersystems {
		rs.Draw(screen)
	}
}

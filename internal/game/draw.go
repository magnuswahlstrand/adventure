package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"golang.org/x/image/colornames"
)

const (
	gridSize = 10
)

func (g *Game) drawGrid(screen *ebiten.Image) {
	for iy := float64(0); iy < 10; iy++ {
		for ix := float64(0); ix < 10; ix++ {
			ebitenutil.DrawRect(screen, 1+gridSize*ix, 1+gridSize*iy, gridSize-2, gridSize-2, colornames.Yellow)
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawGrid(screen)

	for _, rs := range g.rendersystems {
		rs.Draw(screen)
	}
}

package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/kyeett/adventure/internal/comp"
	"github.com/kyeett/adventure/internal/rendersystem/drawutil"
	"golang.org/x/image/colornames"
)

const (
	gridSize   = 16
	borderSize = 1
)

func (g *Game) drawGrid(screen *ebiten.Image) {
	for iy := float64(0); iy < 6; iy++ {
		for ix := float64(0); ix < 6; ix++ {
			ebitenutil.DrawRect(screen, (gridSize+borderSize)*ix+borderSize, (gridSize+borderSize)*iy+borderSize, gridSize, gridSize, colornames.Darkgoldenrod)
		}
	}
}

func (g *Game) drawWall(screen *ebiten.Image) {
	for i := 0; i < 6; i++ {
		if i == 2 {
			continue
		}
		drawutil.DrawSprite(screen, comp.PP(i, 2), comp.SpriteWall.Image)
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "press 'z' to undo", 10, 2*70)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("actions left: %d", actionsPerTurn - g.ActionStack.actions), 2*60, 0)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("turn: %d", g.ActionStack.turn), 2*60, 15)

	// Draw stack
	for i, s := range g.events {
		ebitenutil.DebugPrintAt(screen, string(s.Type()), 2*60, (30 + 15*i))
	}
	g.drawGrid(screen)
	g.drawWall(screen)

	for _, rs := range g.rendersystems {
		rs.Draw(screen)
	}
}
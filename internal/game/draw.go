package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "press 'z' to undo", 10, 200)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("actions left: %d", actionsPerTurn - g.ActionStack.actions), 195, 0)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("turn: %d", g.ActionStack.turn), 195, 15)

	// Draw stack
	for i, s := range g.events {
		ebitenutil.DebugPrintAt(screen, string(s.Type()), 195, (30 + 15*i))
	}

	for _, rs := range g.rendersystems {
		rs.Draw(screen)
	}
}
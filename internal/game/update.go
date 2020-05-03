package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/kyeett/adventure/internal/event"
)

func (g *Game) Update(_ *ebiten.Image) error {
	switch *g.state {
	case "finished":
		g.updateGameFinished()
	default:
		g.updateGameOngoing()
	}
	return nil
}

func (g *Game) updateGameFinished() {
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.logger.Info("R pressed")
		g.restart()
	}
}

func (g *Game) restart() {
	*g = *(NewDefaultGame(1))
}

func (g *Game) restart2Players() {
	*g = *(NewDefaultGame(2))
}

func (g *Game) updateGameOngoing() {
	if inpututil.IsKeyJustPressed(ebiten.Key1) {
		g.restart()
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.Key2) {
		g.restart2Players()
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.Key8) {
		*g = *(NewWebsocketGame(true))
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.Key9) {
		*g = *(NewWebsocketGame(false))
		return
	}

	g.ActionStack.updatedThisIteration = false

	// Handle input
	if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
		g.undo()
	}

	var evt event.Event
	switch g.activePlayer.Remote {
	case true:
		evt = g.network.GetEvents()
	default:
		evt = g.InputHandlers[g.activePlayer.ID].GetEvent()
	}

	if evt != nil {
		g.logger.Debug("event:" + evt.Type())
		g.ActionStack.events = append(g.ActionStack.events, evt)
	}

	for _, s := range g.systems {
		commands := s.Update(evt)
		g.execute(commands)
	}

	g.incrementStep()
	g.incrementRound()
}

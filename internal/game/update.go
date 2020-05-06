package game

import (
	"errors"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/kyeett/adventure/internal/event"
)

var (
	ErrorEndOfGame = errors.New("End of game")
)

func (g *Game) Update() error {

	switch *g.state {
	case "finished":
		return ErrorEndOfGame
	default:
		g.updateGameOngoing()
	}
	return nil
}

func (g *Game) updateGameOngoing() {


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

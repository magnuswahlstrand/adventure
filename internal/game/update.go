package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
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

func (g *Game) updateGameFinished()  {
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.logger.Info("R pressed")
		g.restart()
	}
}

func (g *Game) restart()  {
	g.logger.Info("restart game")
	*g = *(NewGame())
	g.logger.Info(g.state)
}

func (g *Game) updateGameOngoing()  {
	g.GameState.updatedThisIteration = false

	// Handle input
	if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
		g.undo()
	}

	event := g.translation.GetEvent()
	if event != nil {
		g.logger.Debug("event:" + event.Type())
		g.GameState.events = append(g.GameState.events, event)
	}

	for _, s := range g.systems {
		commands := s.Update(event)
		g.execute(commands)
	}

	g.incrementStep()
	g.incrementRound()
}



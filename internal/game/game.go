package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/adventure/internal/command"
	"github.com/kyeett/adventure/internal/comp"
	"github.com/kyeett/adventure/internal/event"
	"github.com/kyeett/adventure/internal/inputhandler"
	"github.com/kyeett/adventure/internal/network"
	"github.com/kyeett/adventure/internal/rendersystem"
	"github.com/kyeett/adventure/internal/system"
	"github.com/kyeett/adventure/internal/unit"
	"go.uber.org/zap"
	"log"
)

type Game struct {
	*ActionStack
	InputHandlers map[comp.ID]inputhandler.InputHandler
	systems       []system.System
	rendersystems []rendersystem.System
	lookup        map[comp.ID]interface{}
	logger        *zap.SugaredLogger
	state         *string
	activePlayer  *unit.Player
	playerList    []*unit.Player
	network       network.Controller
}



func (g *Game) FindEntityByID(id comp.ID) interface{} {
	e, found := g.lookup[id]
	if !found {
		log.Fatal("not found")
	}

	return e
}

type ActionStack struct {
	stack                []*command.Command
	updatedThisIteration bool
	actions              int64
	turn                 int64
	events               []event.Event
}


func (g *Game) Add(v interface{}) {
	type Entity interface {
		GetEntity() comp.Entity
	}

	e, valid := v.(Entity)
	if !valid {
		log.Fatal("not an entity")
	}
	g.lookup[e.GetEntity().ID] = e

	for _, s := range g.systems {
		s.Add(v)
	}
	for _, rs := range g.rendersystems {
		rs.Add(v)
	}

}

func (g *Game) Remove(id comp.ID) {
	for _, s := range g.systems {
		s.Remove(id)
	}
	for _, rs := range g.rendersystems {
		rs.Remove(id)
	}
}

func NewGameState() *ActionStack {
	return &ActionStack{
		stack: []*command.Command{},
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth / 2, outsideHeight / 2
}

func (g *Game) execute(commands []*command.Command) {
	if commands == nil {
		return
	}

	// Mark update as true
	g.ActionStack.updatedThisIteration = true

	for _, c := range commands {
		if err := c.Execute(); err != nil {
			log.Fatal(err)
		}
		c.Step = g.ActionStack.actions
		g.stack = append(g.stack, c)
	}
}

// TODO: Hard to understand. Clean up
func (g *Game) undo() {
	size := len(g.stack) - 1
	if size < 0 {
		return
	}
	var updated bool
	firstStep := g.stack[size].Step
	n := size
	for {
		if n < 0 {
			break
		}

		if g.stack[n].Step != firstStep {
			break
		}

		updated = true

		g.stack[n].Undo()
		g.stack[n] = nil
		g.stack = g.stack[:n]
		n--
	}

	if updated {
		g.ActionStack.actions--
		g.events = g.events[:len(g.events)-1]
	}
}

func (g *Game) incrementStep() {
	if g.ActionStack.updatedThisIteration {
		g.ActionStack.actions++
	}
}

const actionsPerTurn = 3

func (g *Game) incrementRound() {
	endOfTurn := g.ActionStack.actions >= actionsPerTurn
	if endOfTurn {
		g.ActionStack.actions = 0
		g.ActionStack.turn++
		g.ActionStack.stack = nil

		if !g.activePlayer.Remote {
			g.network.Broadcast(g.ActionStack.events)
		}


		g.ActionStack.events = nil

		for i, p := range g.playerList {
			// Go to next player

			if p == g.activePlayer {
				n := (i + 1) % len(g.playerList)
				g.activePlayer = g.playerList[n]
				break
			}
		}
	}
}

var _ ebiten.Game = &Game{}

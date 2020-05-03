package game

import (
	"github.com/kyeett/adventure/internal/comp"
	"github.com/kyeett/adventure/internal/inputhandler"
	"github.com/kyeett/adventure/internal/inputhandler/playerinput"
	"github.com/kyeett/adventure/internal/logger"
	"github.com/kyeett/adventure/internal/network"
	"github.com/kyeett/adventure/internal/rendersystem"
	"github.com/kyeett/adventure/internal/rendersystem/playerui"
	"github.com/kyeett/adventure/internal/system"
	"github.com/kyeett/adventure/internal/system/attack"
	"github.com/kyeett/adventure/internal/system/base"
	"github.com/kyeett/adventure/internal/system/death"
	eventsys "github.com/kyeett/adventure/internal/system/event"
	"github.com/kyeett/adventure/internal/system/movement"
	"github.com/kyeett/adventure/internal/unit"
	"go.uber.org/zap"
)

func stringPointer(s string) *string {
	return &s
}

func NewGame(players ...*unit.Player) *Game {
	comp.ResetTypeCounter()
	firstPlayer := players[0]
	g := &Game{
		ActionStack: NewGameState(),
		rendersystems: []rendersystem.System{
			rendersystem.NewRender(zap.InfoLevel),
			playerui.NewRender(zap.InfoLevel, firstPlayer),
		},
		lookup:       map[comp.ID]interface{}{},
		logger:       logger.NewNamed("game", zap.InfoLevel, logger.BrightWhite),
		state:        stringPointer("started"),
		activePlayer: firstPlayer,
		network:      &network.NoOp{},
		playerList:   players,
	}

	systems := []system.System{
		base.NewSystem(zap.InfoLevel, g),
		movement.NewSystem(zap.InfoLevel),
		attack.NewSystem(zap.InfoLevel, g),
		death.NewSystem(zap.InfoLevel, g),
		eventsys.NewSystem(zap.InfoLevel, g.state),
	}

	inputs := map[comp.ID]inputhandler.InputHandler{}
	for _, p := range players {
		h := playerinput.NewTranslation(zap.DebugLevel, p)
		inputs[p.ID] = h
		systems = append(systems, h)
	}

	g.InputHandlers = inputs
	g.systems = systems
	return g
}

func NewDefaultGame(numPlayers int) *Game {
	p := unit.NewPlayer(2, 3)

	units := []interface{}{
		unit.NewEnemySnake(4, 5),
		unit.NewEnemyRat(5, 4),
		unit.NewChest(5, 5),
		unit.NewDoor(2, 2),
		unit.NewGoal(2, 0),
	}

	players := []*unit.Player{p}
	units = append(units, p)
	if numPlayers > 1 {
		p2 := unit.NewPlayer(2, 4)
		players = append(players, p2)
		units = append(units, p2)
	}

	g := NewGame(players...)
	for _, u := range units {
		g.Add(u)
	}
	return g
}

func NewWebsocketGame(hosting bool) *Game {
	p1 := unit.NewPlayer(2, 3)
	p2 := unit.NewPlayer(2, 4)

	switch hosting {
	case true:
		p1.Remote = false
		p2.Remote = true
	case false:
		p2.Remote = false
		p1.Remote = true
	}

	units := []interface{}{
		unit.NewEnemySnake(4, 5),
		unit.NewEnemyRat(5, 4),
		unit.NewChest(5, 5),
		unit.NewDoor(2, 2),
		unit.NewGoal(2, 0),
		p1,
		p2,
	}

	players := []*unit.Player{p1, p2}
	g := NewGame(players...)
	g.network = network.NewWebsocketConnection("1234")
	for _, u := range units {
		g.Add(u)
	}
	return g
}

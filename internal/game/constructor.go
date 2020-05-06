package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/adventure/assets"
	"github.com/kyeett/adventure/internal/comp"
	"github.com/kyeett/adventure/internal/inputhandler"
	"github.com/kyeett/adventure/internal/inputhandler/playerinput"
	"github.com/kyeett/adventure/internal/logger"
	"github.com/kyeett/adventure/internal/network"
	"github.com/kyeett/adventure/internal/rendersystem"
	"github.com/kyeett/adventure/internal/rendersystem/background"
	"github.com/kyeett/adventure/internal/rendersystem/playerui"
	"github.com/kyeett/adventure/internal/system"
	"github.com/kyeett/adventure/internal/system/attack"
	"github.com/kyeett/adventure/internal/system/base"
	"github.com/kyeett/adventure/internal/system/death"
	eventsys "github.com/kyeett/adventure/internal/system/event"
	"github.com/kyeett/adventure/internal/system/movement"
	"github.com/kyeett/adventure/internal/unit"
	"github.com/kyeett/tiledutil"
	"github.com/lafriks/go-tiled"
	"go.uber.org/zap"
	"log"
)

func stringPointer(s string) *string {
	return &s
}

func convertPositions(ps []tiledutil.Position) []comp.Position {
	var positions []comp.Position
	for _, p := range ps {
		positions = append(positions, comp.Position{
			X: int(p.X),
			Y: int(p.Y),
		})
	}
	return positions
}

var FilterOnlyWalls = func(l *tiled.Layer, tile *tiled.LayerTile, typ string) bool {
	return typ == "wall"
}

func NewGame(players ...*unit.Player) *Game {
	comp.ResetTypeCounter()

	logger := logger.NewNamed("game", zap.InfoLevel, logger.BrightWhite)

	m := tiledutil.MustFromBytes("assets/maps/world.tmx", func(path string) []byte {
		return assets.MustAsset(path)
	})
	logger.Info("map load")
	wallPositions := convertPositions(m.Positions(FilterOnlyWalls))
	logger.Info("wall pos")
	backgroundImage, err := ebiten.NewImageFromImage(m.MustImage(FilterOnlyWalls), ebiten.FilterDefault)
	if err != nil {
		log.Fatal("new ebiten image", err)
	}

	logger.Info("ebiten image")

	firstPlayer := players[0]
	g := &Game{
		ActionStack: NewGameState(),
		rendersystems: []rendersystem.System{
			background.NewRender(zap.DebugLevel, backgroundImage),
			rendersystem.NewRender(zap.InfoLevel),
			playerui.NewRender(zap.InfoLevel, firstPlayer),
		},
		lookup:       map[comp.ID]interface{}{},
		logger:       logger,
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
		h := playerinput.NewHandler(zap.DebugLevel, p, wallPositions, m.Width, m.Height)
		inputs[p.ID] = h
		systems = append(systems, h)
	}

	g.InputHandlers = inputs
	g.systems = systems

	for _, t := range m.Tiles() {
		var u interface{}
		switch t.Type {
		case "rat":
			u = unit.NewEnemyRat(t.X, t.Y)
		case "snake":
			u = unit.NewEnemySnake(t.X, t.Y)
		case "chest":
			u = unit.NewChest(t.X, t.Y)
		case "door":
			u = unit.NewDoor(t.X, t.Y)
		case "goal":
			u = unit.NewGoal(t.X, t.Y)
		case "player_1":
			players[0].X = t.X
			players[0].Y = t.Y
		case "player_2":
			if len(players) > 1 {
				players[1].X = t.X
				players[1].Y = t.Y
			}
		}

		if u != nil {
			g.Add(u)
		}
	}

	return g
}

func NewDefaultGame(numPlayers int) *Game {
	p := unit.NewPlayer(2, 3)

	units := []interface{}{}

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

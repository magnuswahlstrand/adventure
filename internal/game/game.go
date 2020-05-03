package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/single-player-game/internal/command"
	"github.com/kyeett/single-player-game/internal/comp"
	"github.com/kyeett/single-player-game/internal/event"
	"github.com/kyeett/single-player-game/internal/logger"
	"github.com/kyeett/single-player-game/internal/rendersystem"
	"github.com/kyeett/single-player-game/internal/rendersystem/playerui"
	"github.com/kyeett/single-player-game/internal/system"
	"github.com/kyeett/single-player-game/internal/system/attack"
	"github.com/kyeett/single-player-game/internal/system/base"
	"github.com/kyeett/single-player-game/internal/system/death"
	eventsys "github.com/kyeett/single-player-game/internal/system/event"
	"github.com/kyeett/single-player-game/internal/system/movement"
	"github.com/kyeett/single-player-game/internal/system/translation"
	"github.com/kyeett/single-player-game/internal/unit"
	"go.uber.org/zap"
	"log"
)

type Game struct {
	*GameState
	translation   *translation.Translation
	systems       []system.System
	rendersystems []rendersystem.System
	lookup        map[comp.ID]interface{}
	logger        *zap.SugaredLogger
	state         *string
}

func (g *Game) FindEntityByID(id comp.ID) interface{} {
	e, found := g.lookup[id]
	if !found {
		log.Fatal("not found")
	}

	return e
}

type GameState struct {
	stack                []*command.Command
	updatedThisIteration bool
	step                 int64
	turn                 int64
	events               []event.Event
}

func NewGame() *Game {

	p := unit.NewPlayer(2, 3)


	e := unit.NewEnemySnake(4, 5)
	e2 := unit.NewEnemyRat(5, 4)

	c := unit.NewChest(5, 5)

	d := unit.NewDoor(2, 2)
	gl := unit.NewGoal(2, 0)

	state := "started"
	g := &Game{
		GameState: NewGameState(),
		rendersystems: []rendersystem.System{
			rendersystem.NewRender(zap.InfoLevel),
			playerui.NewRender(zap.InfoLevel, p),
		},
		lookup: map[comp.ID]interface{}{},
		logger: logger.NewNamed("game", zap.InfoLevel, logger.BrightWhite),
		state: &state,
	}

	trans := translation.NewTranslation(zap.DebugLevel, g, p)
	systems := []system.System{
		trans,
		base.NewSystem(zap.InfoLevel, g, p),
		movement.NewSystem(zap.InfoLevel),
		attack.NewSystem(zap.InfoLevel, g),
		death.NewSystem(zap.InfoLevel, g),
		eventsys.NewSystem(zap.InfoLevel, g.state),
	}
	g.translation = trans
	g.systems = systems

	g.Add(p)
	g.Add(e)
	g.Add(e2)
	g.Add(c)
	g.Add(d)
	g.Add(gl)

	return g
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

func NewGameState() *GameState {
	return &GameState{
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
	g.GameState.updatedThisIteration = true

	for _, c := range commands {
		if err := c.Execute(); err != nil {
			log.Fatal(err)
		}
		c.Step = g.GameState.step
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
		g.GameState.step--
		g.events = g.events[:len(g.events)-1]
	}
}

func (g *Game) incrementStep() {
	if g.GameState.updatedThisIteration {
		g.GameState.step++
	}
}

const stepsPerTurn = 5

func (g *Game) incrementRound() {
	if g.GameState.step > stepsPerTurn {
		g.GameState.step = 0
		g.GameState.turn++
		g.GameState.stack = nil
		g.GameState.events = nil
	}
}

var _ ebiten.Game = &Game{}

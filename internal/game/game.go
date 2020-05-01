package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/kyeett/single-player-game/internal/command"
	"github.com/kyeett/single-player-game/internal/comp"
	"github.com/kyeett/single-player-game/internal/rendersystem"
	"github.com/kyeett/single-player-game/internal/system"
	"github.com/kyeett/single-player-game/internal/system/death"
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
}

func NewGame() *Game {

	p := unit.NewPlayer(0, 0)
	e := unit.NewEnemySnake(2, 2)
	e2 := unit.NewEnemyRat(5, 0)
	c := unit.NewChest(1, 1)

	g := &Game{
		GameState: NewGameState(),
		rendersystems: []rendersystem.System{
			rendersystem.NewRender(zap.InfoLevel),
		},
		lookup: map[comp.ID]interface{}{},
	}

	trans := translation.NewTranslation(zap.InfoLevel, g, p)
	systems := []system.System{
		trans,
		death.NewSystem(zap.InfoLevel, g),
	}
	g.translation = trans
	g.systems = systems

	g.Add(p)
	g.Add(e)
	g.Add(e2)
	g.Add(c)

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

func (g *Game) Update(_ *ebiten.Image) error {
	g.GameState.updatedThisIteration = false

	// Handle input
	if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
		g.undo()
	}

	for _, s := range g.systems {
		commands := s.Update()
		g.execute(commands)
	}

	g.incrementStep()
	g.incrementRound()
	return nil
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
	}
}

var _ ebiten.Game = &Game{}

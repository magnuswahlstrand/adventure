package game

import (
	"fmt"
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
	lookup map[comp.ID]interface{}
}

func (g *Game) FindEntityByID(id comp.ID) interface{} {
	e, found := g.lookup[id]
	if !found {
		log.Fatal("not found")
	}

	return e
}

type GameState struct {
	stack  []*command.Command
	player *unit.Player
}

func NewGame() *Game {

	p := unit.NewPlayer(0, 0)
	e := unit.NewEnemySnake(2, 2)
	e2 := unit.NewEnemyRat(5, 0)
	c := unit.NewChest(1, 1)

	g := &Game{
		GameState: NewGameState(p),
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
	type Entity interface{
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

func NewGameState(p *unit.Player) *GameState {
	return &GameState{
		stack:  []*command.Command{},
		player: p,
	}
}



func (g *Game) Update(_ *ebiten.Image) error {



	// Handle input
	if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
		g.undo()
		return nil
	}

	//command := g.handleInput()

	// Translate input to commands
	//commands := g.translation.Translate(command)

	// Execute commands

	for _, s := range g.systems {
		commands := s.Update()
		g.execute(commands)
	}

	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth / 5, outsideHeight / 5
}

func (g *Game) execute(commands []*command.Command) {
	if commands == nil {
		return
	}

	for _, c := range commands {
		if err := c.Execute(); err != nil {
			log.Fatal(err)
		}
		g.stack = append(g.stack, c)
	}
}

func (g *Game) undo() {
	n := len(g.stack) - 1
	if n < 0 {
		return
	}
	c := g.stack[n]

	fmt.Println("undo", c.Name)
	c.Undo()

	// Remove from stack
	g.stack[n] = nil
	g.stack = g.stack[:n]
}

var _ ebiten.Game = &Game{}

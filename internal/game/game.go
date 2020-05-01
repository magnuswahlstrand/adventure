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
	}

	trans := translation.NewTranslation(zap.InfoLevel, g)
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

func (g *Game) handleInput() interface{} {
	if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
		g.undo()
		return nil
	}

	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyRight):
		return command.MoveBy2(g.player, 1, 0)
	case inpututil.IsKeyJustPressed(ebiten.KeyLeft):
		return command.MoveBy2(g.player, -1, 0)
	case inpututil.IsKeyJustPressed(ebiten.KeyUp):
		return command.MoveBy2(g.player, 0, -1)
	case inpututil.IsKeyJustPressed(ebiten.KeyDown):
		return command.MoveBy2(g.player, 0, 1)
	}

	return nil
}

func (g *Game) Update(_ *ebiten.Image) error {

	// Handle input
	command := g.handleInput()

	// Translate input to commands
	commands := g.translation.Translate(command)

	// Execute commands
	g.execute(commands)

	for _, s := range g.systems {
		s.Update(0)
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

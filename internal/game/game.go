package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/kyeett/single-player-game/internal/command"
	"github.com/kyeett/single-player-game/internal/comp"
	"github.com/kyeett/single-player-game/internal/rendersystem"
)

type Player struct {
	*comp.Position
}

type Game struct {
	*GameState
	//systems []system.System
	rendersystems []rendersystem.System
}

type GameState struct {
	stack    []*command.Command
	player   *Player
}

func NewGame() *Game{

	p := &Player{
		&comp.Position{1,2},
	}
	g := &Game{
		GameState: NewGameState(p),
		rendersystems: []rendersystem.System{
			rendersystem.NewRender(),
		},
	}

	g.Add(p)
	return g
}

func (g *Game) Add(v interface{}) {
	for _, rs := range g.rendersystems {
		rs.Add(v)
	}
}

func NewGameState(p *Player) *GameState {
	return &GameState{
		stack: []*command.Command{},
		player: p,
	}
}



func (g *Game) handleInput() *command.Command {
	if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
		g.undo()
		return nil
	}

	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyRight):
		return command.MoveBy(g.player, 1, 0)
	case inpututil.IsKeyJustPressed(ebiten.KeyLeft):
		return command.MoveBy(g.player, -1, 0)
	case inpututil.IsKeyJustPressed(ebiten.KeyUp):
		return command.MoveBy(g.player, 0, -1)
	case inpututil.IsKeyJustPressed(ebiten.KeyDown):
		return command.MoveBy(g.player, 0, 1)
	}

	return nil
}

func (g *Game) Update(_ *ebiten.Image) error {



	// Handle input
	command := g.handleInput()
	//command = g.translate(command)
	g.execute(command)

	//for _, s := range w.systems {
	//	s.Update(timeStep)
	//}


	return nil
}

func (g *Game) translate(c *command.Command) *command.Command {
	//switch c.Type {
	//case command.TypeMoveBy:
	//	target := g.GameState.At()
	//	if
	//default:
	//	return nil
	//}
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth / 4, outsideHeight / 4
}

func (g *Game) execute(c *command.Command) {
	if c == nil {
		return
	}

	fmt.Println("exec",c.Name)

	if err := c.Execute(); err != nil {
		return
	}
	g.stack = append(g.stack, c)
}

func (g *Game) undo()  {
	n := len(g.stack) - 1
	if n < 0 {
		return
	}
	c := g.stack[n]

	fmt.Println("undo",c.Name)
	c.Undo()

	// Remove from stack
	g.stack[n] = nil
	g.stack = g.stack[:n]
}




var _ ebiten.Game = &Game{}

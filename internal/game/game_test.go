package game

import (
	"encoding/json"
	"github.com/kyeett/single-player-game/internal/comp"
	"github.com/kyeett/single-player-game/internal/event"
	"github.com/kyeett/single-player-game/internal/inputhandler"
	"github.com/kyeett/single-player-game/internal/inputhandler/testinput"
	"github.com/kyeett/single-player-game/internal/unit"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"testing"
)

func testSetup() (*unit.Player, *Game, *testinput.TestInput) {
	p := unit.NewPlayer(0, 0)
	g := NewGame(p)
	g.Add(p)
	input := testinput.New(zap.InfoLevel)
	g.InputHandlers = map[comp.ID]inputhandler.InputHandler{
		p.ID: input,
	}
	return p, g, input
}

func testSetup2Players() (*unit.Player, *unit.Player, *Game, *testinput.TestInput) {
	p := unit.NewPlayer(0, 0)
	p2 := unit.NewPlayer(0, 0)
	g := NewGame(p, p2)
	g.Add(p)
	g.Add(p2)

	input := testinput.New(zap.InfoLevel)
	g.InputHandlers = map[comp.ID]inputhandler.InputHandler{
		p.ID:  input,
		p2.ID: input,
	}
	return p, p2, g, input
}

func TestMove(t *testing.T) {
	player, g, input := testSetup()

	// Act
	pos := *player.Position
	g.Update(nil)

	// Assert
	assert.EqualValues(t, pos, *player.Position)

	// Arrange
	newPos := *comp.PP(4, 4)
	input.AddEvent(event.Move{
		Actor:    player.ID,
		Position: newPos,
	})

	// Act
	g.Update(nil)

	// Assert
	assert.EqualValues(t, newPos, *player.Position)
}

func TestMarshalEvent(t *testing.T) {
	player, g, input := testSetup()
	target := comp.Position{4, 4}
	evtBefore := event.Move{
		Actor:    player.ID,
		Position: target,
	}

	b, err := json.Marshal(evtBefore)
	require.NoError(t, err)

	// "Send over network"

	var evt event.Move
	require.NoError(t, json.Unmarshal(b, &evt))

	// Arrange
	input.AddEvent(evt)

	// Act
	g.Update(nil)

	// Assert
	assert.EqualValues(t, target, *player.Position)
}

func TestTwoPlayers(t *testing.T) {
	p1, p2, g, input := testSetup2Players()

	//target := comp.Position{0, 0}
	evt := event.Move{Actor: p1.ID, Position: comp.P(0,0)}

	// Arrange
	for i := 0; i < actionsPerTurn; i++ {
		input.AddEvent(evt)
		g.Update(nil)
	}

	assert.Equal(t, p2, g.activePlayer)

	// Move 2nd player
	target := comp.P(2, 2)
	evt = event.Move{Actor: p2.ID, Position: target}

	// Act
	input.AddEvent(evt)
	g.Update(nil)

	assert.EqualValues(t, target, *p2.Position)
}

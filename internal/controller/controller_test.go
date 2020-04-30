package controller

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var test3x3Grid, _ = NewGrid(
	"---",
	"---",
	"---")

func TestWalkable(t *testing.T) {
	g := test3x3Grid

	p := &Player{pos{1, 1}}
	g.Set(1,1, p)

	assert.True(t, g.At(0,0).Walkable())
	assert.False(t, g.At(-1,-1).Walkable())
	assert.False(t, g.At(4,4).Walkable())
	assert.False(t, g.At(1,1).Walkable())
}

type MovePlayer struct{
	pos
}

func TestController(t *testing.T) {
	g := test3x3Grid

	p := &Player{pos{1, 1}}
	g.Set(1,1, p)

	c := Controller{
		Player: p,
		Grid:   g,
	}

	actions := c.EvaluateInput(DownPressed)
	assert.Equal(t, []interface{}{MovePlayer{pos{1,2}}}, actions)
}
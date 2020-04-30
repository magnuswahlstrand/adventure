package unit

import "github.com/kyeett/single-player-game/internal/comp"

type Player struct {
	comp.Entity
	*comp.Position
	*comp.Sprite
}

type Enemy struct {
	comp.Entity
	*comp.Position
	*comp.Sprite
}

func NewEnemy() *Enemy {
	return &Enemy{
		comp.Entity{
			comp.NewID(),
			comp.TypeEnemy,
		},
		comp.P(2, 2),
		comp.SpriteEnemySnake,
	}
}

func NewPlayer() *Player {
	return &Player{
		comp.Entity{
			comp.NewID(),
			comp.TypePlayer,
		},
		comp.P(0, 0),
		comp.SpritePlayer,
	}
}

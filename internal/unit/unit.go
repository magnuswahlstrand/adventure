package unit

import "github.com/kyeett/single-player-game/internal/comp"

type Player struct {
	comp.Entity
	*comp.Position
	*comp.Sprite
	*comp.Hitpoints
}

type Enemy struct {
	comp.Entity
	*comp.Position
	*comp.Sprite
	*comp.Hitpoints
}

type Chest struct {
	comp.Entity
	*comp.Position
	*comp.Sprite
	*comp.Content
}

type Item struct {
	comp.Entity
	*comp.Position
	*comp.Sprite
}

func NewEnemySnake(x, y int) *Enemy {
	return &Enemy{
		comp.Entity{
			comp.NewID(),
			comp.TypeEnemy,
		},
		comp.P(x, y),
		comp.SpriteEnemySnake,
		comp.HP(3),
	}
}

func NewEnemyRat(x, y int) *Enemy {
	return &Enemy{
		comp.Entity{
			comp.NewID(),
			comp.TypeEnemy,
		},
		comp.P(x, y),
		comp.SpriteEnemyRat,
		comp.HP(2),
	}
}

func NewPlayer(x, y int) *Player {
	return &Player{
		comp.Entity{
			comp.NewID(),
			comp.TypePlayer,
		},
		comp.P(x, y),
		comp.SpritePlayer,
		comp.HP(5),
	}
}

func NewChest(x, y int) *Chest {
	return &Chest{
		comp.Entity{
			comp.NewID(),
			comp.TypeChest,
		},
		comp.P(x, y),
		comp.SpriteChest,
		&comp.Content{"key"},
	}
}

func NewItem(x, y int) *Item {
	return &Item{
		comp.Entity{
			comp.NewID(),
			comp.TypeItem,
		},
		comp.P(x, y),
		comp.SpriteKey,
	}
}
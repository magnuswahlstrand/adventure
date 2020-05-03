package unit

import (
	"github.com/kyeett/single-player-game/internal/comp"
)

type Player struct {
	comp.Entity
	*comp.Position
	*comp.Sprite
	*comp.Hitpoints
	*comp.Inventory
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

type Door struct {
	comp.Entity
	*comp.Position
	*comp.Sprite
}

type Goal struct {
	comp.Entity
	*comp.Position
	*comp.Sprite
}

func NewEnemySnake(x, y int) *Enemy {
	return &Enemy{
		comp.Entity{
			comp.NewIDFromType(comp.TypeEnemy),
			comp.TypeEnemy,
		},
		comp.PP(x, y),
		comp.SpriteEnemySnake,
		comp.HP(3),
	}
}

func NewEnemyRat(x, y int) *Enemy {
	return &Enemy{
		comp.Entity{
			comp.NewIDFromType(comp.TypeEnemy),
			comp.TypeEnemy,
		},
		comp.PP(x, y),
		comp.SpriteEnemyRat,
		comp.HP(2),
	}
}

func NewPlayer(x, y int) *Player {
	entity := comp.Entity{
		comp.NewIDFromType(comp.TypePlayer),
		comp.TypePlayer,
	}
	var sprite *comp.Sprite
	switch entity.ID {
	case "player_0002":
		sprite = comp.SpritePlayer2
	case "player_0001":
		sprite = comp.SpritePlayer
	default:
		sprite = comp.SpritePlayer
	}

	return &Player{
		entity,
		comp.PP(x, y),
		sprite,
		comp.HP(5),
		&comp.Inventory{},
	}
}

func NewChest(x, y int) *Chest {
	return &Chest{
		comp.Entity{
			comp.NewIDFromType(comp.TypeChest),
			comp.TypeChest,
		},
		comp.PP(x, y),
		comp.SpriteChest,
		&comp.Content{"key"},
	}
}

func NewItem(x, y int) *Item {
	return &Item{
		comp.Entity{
			comp.NewIDFromType(comp.TypeItem),
			comp.TypeItem,
		},
		comp.PP(x, y),
		comp.SpriteKey,
	}
}

func NewDoor(x, y int) *Door {
	return &Door{
		comp.Entity{
			comp.NewIDFromType(comp.TypeDoor),
			comp.TypeDoor,
		},
		comp.PP(x, y),
		comp.SpriteDoor,
	}
}

func NewGoal(x, y int) *Goal {
	return &Goal{
		comp.Entity{
			comp.NewIDFromType(comp.TypeGoal),
			comp.TypeGoal,
		},
		comp.PP(x, y),
		comp.SpriteGoal,
	}
}

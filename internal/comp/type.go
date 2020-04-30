package comp

import (
	"github.com/google/uuid"
)

type Entity struct {
	ID ID
	Type
}

type Type string
type ID string

const (
	TypeNil    Type = ""
	TypeWall        = "wall"
	TypeEnemy       = "enemy"
	TypePlayer      = "player"
)

var (
	BaseNil  = Entity{Type: TypeNil}
	BaseWall = Entity{Type: TypeWall}
)

func (b Entity) GetType() Type {
	return b.Type
}
func (b Entity) GetID() ID {
	return b.ID
}

func (b Entity) GetEntity() Entity {
	return b
}

func NewID() ID {
	return ID(uuid.New().String())
}

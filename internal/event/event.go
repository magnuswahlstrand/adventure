package event

import "github.com/kyeett/single-player-game/internal/comp"

type Type string

const (
	TypeMove      = "Move"
	TypeAttack    = "Attack"
	TypeOpenChest = "Open Chest"
	TypeTakeItem  = "Take Item"
	TypeOpenDoor  = "Open Door"
	TypeReachGoal  = "Reach Goal"
	TypeUnknown   = "Unknown"
)

type Event interface {
	Type() Type
}

type Move struct {
	Actor  comp.ID
	Position comp.Position
}

type Attack struct {
	Actor, Target comp.ID
}

type OpenChest struct {
	Actor, Target comp.ID
}

type TakeItem struct {
	Actor, Target comp.ID
}

type OpenDoor struct {
	Actor, Target comp.ID
}

type ReachGoal struct {}

func (Move) Type() Type      { return TypeMove }
func (Attack) Type() Type    { return TypeAttack }
func (OpenChest) Type() Type { return TypeOpenChest }
func (TakeItem) Type() Type  { return TypeTakeItem }
func (OpenDoor) Type() Type  { return TypeOpenDoor }
func (ReachGoal) Type() Type  { return TypeReachGoal }
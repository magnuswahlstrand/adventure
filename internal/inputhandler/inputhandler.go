package inputhandler

import (
	"github.com/kyeett/single-player-game/internal/event"
)

type InputHandler interface{
	GetEvent() event.Event
}

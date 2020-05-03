package inputhandler

import (
	"github.com/kyeett/adventure/internal/event"
)

type InputHandler interface{
	GetEvent() event.Event
}

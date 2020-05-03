package network

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/kyeett/adventure/internal/comp"
	"github.com/kyeett/adventure/internal/event"
	"log"
	"net/url"
)

var _ Controller = &NoOp{}

type NoOp struct{}

func (c *NoOp) Broadcast(_ []event.Event) {}
func (c *NoOp) GetEvents() event.Event    { return nil }

type Controller interface {
	Broadcast([]event.Event)
	GetEvents() event.Event
}

type WebsocketConnection struct {
	conn *websocket.Conn
}

func NewWebsocketConnection(roomID string) *WebsocketConnection {
	u := url.URL{Scheme: "wss", Host: "room-server-game.herokuapp.com"}
	u.Path = "/room/" + roomID
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	wsConn := &WebsocketConnection{
		conn: c,
	}

	go wsConn.Start()

	return wsConn
}

type envelope struct{
	typ string
	event.Event
}

func (c *WebsocketConnection) Broadcast(events []event.Event) {
	for _, evt := range events {
		err := c.conn.WriteJSON(envelope{
			typ:   string(evt.Type()),
			Event: evt,
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (c *WebsocketConnection) GetEvents() event.Event    {
	return event.Move{
		Actor:    "player_0002",
		Position: comp.Position{1,1},
	}
}

func (c *WebsocketConnection) Start() {
	for {
		_, message, err := c.conn.ReadMessage()

		var env envelope
		if err != json.Unmarshal(message, &env) {
			log.Fatal(err)
		}

		fmt.Println("received ", env.typ)
	}
}

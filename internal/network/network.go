package network

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/kyeett/adventure/internal/event"
	"log"
	"net/url"
	"sync"
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
	conn             *websocket.Conn
	events           map[int64]event.Event
	lock             *sync.Mutex
	sentN, receivedN int64
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
		conn:   c,
		events: map[int64]event.Event{},
		lock:   &sync.Mutex{},
	}

	go wsConn.Start()

	return wsConn
}

type envelope struct {
	Type  string
	Event string
	N     int64
}

func (c *WebsocketConnection) Broadcast(events []event.Event) {
	fmt.Println("broadcast!!")
	for _, evt := range events {
		b, err := json.Marshal(evt)
		if err != nil {
			log.Fatal(err)
		}

		c.lock.Lock()
		n := c.sentN
		c.sentN++
		c.lock.Unlock()
		err = c.conn.WriteJSON(envelope{
			Type:  string(evt.Type()),
			Event: base64.StdEncoding.EncodeToString(b),
			N: n,
		})
		if err != nil {
			log.Fatal(err)
		}


	}
}

func (c *WebsocketConnection) GetEvents() event.Event {
	c.lock.Lock()
	defer c.lock.Unlock()

	evt, found := c.events[c.receivedN]
	if !found {
		return nil
	}
	c.receivedN++
	return evt
}

func (c *WebsocketConnection) Start() {
	for {
		var env envelope
		err := c.conn.ReadJSON(&env)
		if err != nil {
			log.Fatal("read:", err)
		}

		var receivedEvent event.Event
		switch env.Type {
		case "Move":

			b, err := base64.StdEncoding.DecodeString(env.Event)
			if err != nil {
				log.Fatal(err)
			}

			var evt event.Move
			if err := json.Unmarshal(b, &evt); err != nil {
				log.Fatal("unmarshal:", err)
			}

			receivedEvent = evt

		default:
			log.Fatal("invalid type")
			continue
		}

		fmt.Printf("received %#v\n", receivedEvent)
		c.lock.Lock()
		c.events[env.N] = receivedEvent
		c.lock.Unlock()
	}
}

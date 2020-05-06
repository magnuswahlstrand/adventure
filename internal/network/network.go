package network

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"nhooyr.io/websocket"
	"github.com/kyeett/adventure/internal/event"
	"log"
	"net/url"
	"nhooyr.io/websocket/wsjson"
	"sync"
	"time"
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

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	c, _, err := websocket.Dial(ctx, u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	wsConn := &WebsocketConnection{
		conn:   c,
		events: map[int64]event.Event{},
		lock:   &sync.Mutex{},
	}

	go wsConn.Start()

	log.Printf("connected, and started")
	return wsConn
}

type envelope struct {
	Type  string
	Event string
	N     int64
}

func (c *WebsocketConnection) WriteEvent(evt event.Event) error {
	c.lock.Lock()
	n := c.sentN
	c.sentN++
	c.lock.Unlock()

	b, err := json.Marshal(evt)
	if err != nil {
		log.Fatal(err)
	}

	env := envelope{
		Type:  string(evt.Type()),
		Event: base64.StdEncoding.EncodeToString(b),
		N:     n,
	}

	if err := wsjson.Write(context.Background(), c.conn, env); err != nil {
		log.Fatal(err)
	}
	return nil
}

func (c *WebsocketConnection) Broadcast(events []event.Event) {
	fmt.Println("broadcast!!")
	for _, evt := range events {
		c.WriteEvent(evt)
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
	fmt.Println("starting connection")
	for {
		var env envelope
		if err := wsjson.Read(context.Background(), c.conn, &env); err != nil {
			log.Println("read:", err)
			continue
		}

		var receivedEvent event.Event
		switch env.Type {
		case event.TypeMove:
			var evt event.Move
			c.decodeAndUnmarshal(env, &evt)
			receivedEvent = evt
		case event.TypeAttack:
			var evt event.Attack
			c.decodeAndUnmarshal(env, &evt)
			receivedEvent = evt
		case event.TypeOpenChest:
			var evt event.OpenChest
			c.decodeAndUnmarshal(env, &evt)
			receivedEvent = evt
		case event.TypeTakeItem:
			var evt event.TakeItem
			c.decodeAndUnmarshal(env, &evt)
			receivedEvent = evt
		case event.TypeOpenDoor:
			var evt event.OpenDoor
			c.decodeAndUnmarshal(env, &evt)
			receivedEvent = evt
		case event.TypeReachGoal:
			var evt event.ReachGoal
			c.decodeAndUnmarshal(env, &evt)
			receivedEvent = evt

		default:
			log.Println("invalid type")
			continue
		}

		fmt.Printf("received %#v\n", receivedEvent)
		c.lock.Lock()
		c.events[env.N] = receivedEvent
		c.lock.Unlock()
	}
}

func (c *WebsocketConnection) decodeAndUnmarshal(env envelope, v event.Event) {
	b, err := base64.StdEncoding.DecodeString(env.Event)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(b, v); err != nil {
		log.Fatal("unmarshal:", err)
	}
}

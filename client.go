package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

// https://play.google.com/books/reader?id=Er0pIQAAAEAJ&pg=GBS.PA641

// client represents a single chatting user
type client struct {
	// socket is the web socket for this client.
	socket *websocket.Conn
	//send is a channel on which messages are sent.
	send chan *Event
	// channel is the channel this client is currently joined
	channel *channel
}

//go:generate jsonenums -type=Kind

type Kind int

const (
  UI_LAUNCHER_DISPLAY Kind = iota
  UI_LAUNCHER_HIDE
  ERROR
)

// https://eagain.net/articles/go-dynamic-json/

// Event represents an event container to be exchanged between parties
type Event struct {
	// Type represents an action type
	Type Kind `json:"type"` // ISSUE: https://github.com/campoy/jsonenums/issues/28#issue-299906485
	// User represents mandatory field required for all events exchanged
	User string `json:"user,omitempty"`
	// Payload holds actual Event data
	Payload interface{} `json:"payload,omitempty"`
}

// UILauncherDisplayPayload represents actual data payload carried by Event
type UILauncherDisplayPayload struct {
	Stakeholder string `json:"stakeholder,omitempty"`
}

// UILauncherHidePayload represents actual data payload carried by Event
type UILauncherHidePayload struct {
	Stakeholder string `json:"stakeholder,omitempty"`
}

// ErrorPayload represents actual data payload carried by Event
type ErrorPayload struct {
	Message string `json:"message,omitempty"`
}

var kindHandlers = map[Kind]func() interface{}{
	UI_LAUNCHER_DISPLAY: func() interface{} { return &UILauncherDisplayPayload{} },
	UI_LAUNCHER_HIDE: func() interface{} { return &UILauncherHidePayload{} },
	ERROR: func() interface{} { return &ErrorPayload{} },
}


func (c *client) read() {
	defer c.socket.Close()
	for {
    // Read the message from the socket
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
    }

    // fmt.Println("Pure event:", string(msg))

    var raw json.RawMessage
    evt := Event{
      Payload: &raw,
    }
    if err := json.Unmarshal([]byte(msg), &evt); err != nil {
      log.Fatal("Event:",err, msg)
    }

    m := kindHandlers[evt.Type]()
    if err := json.Unmarshal(raw, m); err != nil {
      // TODO: Reply with error message
      er := Event{
        Type: ERROR,
        Payload: &ErrorPayload{
          Message: "Unsupported action signature",
        },
      }
      c.channel.forward <- &er
      // log.Fatal("Payload:",err,raw)
    }
    switch m.(type) {
    case *UILauncherDisplayPayload:
      c.channel.forward <- &evt
    case *UILauncherHidePayload:
      c.channel.forward <- &evt
    }
	}
}

func (c *client) write() {
	defer c.socket.Close()
	for event := range c.send {
    evt, err := json.Marshal(event)
    if err != nil {
        fmt.Println(err)
        return
    }
		err1 := c.socket.WriteMessage(websocket.TextMessage, evt)
		if err1 != nil {
			return
		}
	}
}

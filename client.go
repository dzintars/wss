package main

import (
	"encoding/json"
	"fmt"

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

// Event represents an event to be exchanged between parties
type Event struct {
  // Type represents an action type
  Type string `json:"type,omitempty"`
  // Payload holds actual Event data
	Payload string `json:"payload,omitempty"`
}

func (c *client) read() {
	defer c.socket.Close()
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
    }
    var e Event
    if err := json.Unmarshal(msg, &e); err != nil {
        panic(err)
    }
		c.channel.forward <- &e
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

package main

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type client struct {
	socket *websocket.Conn
	send   chan []byte
	room   *room
}

func (c *client) read() {
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			fmt.Println("client.read", "-", err)
			break
		}
		c.room.forward <- msg
	}
	close(c.send)
}

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
	c.socket.Close()
}

func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			fmt.Println("client.write", "-", err)
			break
		}
	}
	c.socket.Close()
}

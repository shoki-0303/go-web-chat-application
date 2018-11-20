package main

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type client struct {
	socket   *websocket.Conn
	send     chan message
	room     *room
	userData map[string]interface{}
}

func (c *client) read() {
	for {
		var msg message
		if err := c.socket.ReadJSON(&msg); err == nil {
			msg.Time = time.Now()
			msg.Name = c.userData["name"].(string)
			c.room.forward <- msg
		} else {
			log.Println("client", "-", err)
			break
		}
	}
	c.socket.Close()
}

func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteJSON(msg); err != nil {
			log.Println("client.write", "-", err)
			break
		}
	}
	c.socket.Close()
}

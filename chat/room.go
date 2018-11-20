package main

import (
	"log"
	"net/http"

	"github.com/stretchr/objx"

	"github.com/gorilla/websocket"
)

type room struct {
	forward chan message
	join    chan *client
	leave   chan *client
	clients map[*client]bool
}

func newRoom() *room {
	return &room{
		forward: make(chan message),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
	}
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
			log.Println("client joins")
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.send)
			log.Println("client leaves")
		case msg := <-r.forward:
			log.Println("forwardが" + msg.Message + "を受け取りました")
			for client := range r.clients {
				select {
				case client.send <- msg:
				default:
					log.Println("run", "-", "sendにmsgがおくられませんでした")
					delete(r.clients, client)
					close(client.send)
				}
			}
		}
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 256,
}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println("conn err", err)
		return
	}
	cookie, err := req.Cookie("auth")
	if err != nil {
		log.Println("cookie err", err)
		return
	}
	client := &client{
		socket:   socket,
		send:     make(chan message),
		room:     r,
		userData: make(map[string]interface{}),
	}
	client.userData = objx.MustFromBase64(cookie.Value)
	r.join <- client
	defer func() { r.leave <- client }()
	go client.write()
	client.read()
}

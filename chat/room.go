package main

type room struct {
	forward chan []byte
	join    *client
	leave   *client
	clients map[*client]bool
}

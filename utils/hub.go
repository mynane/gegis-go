package utils

import (
		"encoding/json"
	"fmt"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients 	map[*Client]bool

	// Inbound messages from the clients.
	broadcast 	chan []byte

	// Register requests from the clients.
	register 	chan *Client

	// Unregister requests from clients.
	unregister 	chan *Client

	// Linked room
	rooms 		[]string
}

func NewHub(rooms []string) *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		rooms: 		rooms,
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				var someOne string
				json.Unmarshal(message, &someOne)
				fmt.Println(someOne, "123123123123")
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

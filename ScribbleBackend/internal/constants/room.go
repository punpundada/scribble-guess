package constants

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type Room struct {
	ID      string                   `json:"id"`
	Clients map[*websocket.Conn]bool `json:"clients"`
	Mu      sync.RWMutex
}

func NewRoom(id string) *Room {
	return &Room{
		ID:      id,
		Clients: make(map[*websocket.Conn]bool),
	}
}

type Message struct {
	RoomId string `json:"roomId"`
}

func (r *Room) BroadcastMessage(msg Message) {
	r.Mu.RLock()
	clients := make([]*websocket.Conn, 0, len(r.Clients))
	for c := range r.Clients {
		clients = append(clients, c)
	}
	r.Mu.RUnlock()

	for _, client := range clients {
		err := client.WriteMessage(websocket.TextMessage, []byte(`{"message":"received"}`))
		if err != nil {
			log.Println(err)
			client.Close()

			r.Mu.Lock()
			delete(r.Clients, client)
			r.Mu.Unlock()
		}
	}
}

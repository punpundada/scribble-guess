package constants

import (
	"fmt"
	"log"
	"sync"
)

type BroadcastMessage struct {
	Message  []byte `json:"message"`
	ClientId string `json:"clientId"`
}

func NewBroadcastMessage(message []byte, clientId string) *BroadcastMessage {
	return &BroadcastMessage{
		Message:  message,
		ClientId: clientId,
	}
}

type Room struct {
	ID           string           `json:"id"`
	Clients      map[*Client]bool `json:"clients"`
	Mu           sync.RWMutex
	Register     chan *Client
	Unregister   chan *Client
	Broadcast    chan *BroadcastMessage
	BroadcastAll chan *BroadcastMessage
}

func NewRoom(id string) *Room {
	return &Room{
		ID:           id,
		Clients:      make(map[*Client]bool),
		Register:     make(chan *Client),
		Unregister:   make(chan *Client),
		Broadcast:    make(chan *BroadcastMessage),
		BroadcastAll: make(chan *BroadcastMessage),
	}
}

func (r *Room) Run() {
	for {
		select {
		case client := <-r.Register:
			r.Mu.Lock()
			r.Clients[client] = true
			r.Mu.Unlock()
		case client := <-r.Unregister:
			if _, ok := r.Clients[client]; ok {
				r.Mu.Lock()
				delete(r.Clients, client)
				r.Mu.Unlock()
			}
		case msg := <-r.Broadcast:
			r.BroadcastMessage(msg)
		case msg := <-r.BroadcastAll:
			r.BroadcastAllMessage(msg)
		}
	}
}

func (r *Room) BroadcastMessage(msg *BroadcastMessage) {
	r.Mu.RLock()
	clients := make([]*Client, 0, len(r.Clients))
	for c := range r.Clients {
		log.Println("client id", c.Id)
		if msg.ClientId == c.Id {
			log.Println("Skipping broadcast to client", c.Id, "as it is the sender of the message")
			continue
		}
		clients = append(clients, c)
	}
	r.Mu.RUnlock()
	for _, client := range clients {
		select {
		case client.Send <- fmt.Appendf(nil, `{"message": "%s","to":"%s"}`, msg.Message, msg.ClientId):
			// default:
			// 	r.Mu.Lock()
			// 	delete(r.Clients, client)
			// 	r.Mu.Unlock()
			// 	close(client.Send)
		}
	}
}

func (r *Room) BroadcastAllMessage(msg *BroadcastMessage) {
	r.Mu.RLock()
	clients := make([]*Client, 0, len(r.Clients))
	for c := range r.Clients {
		clients = append(clients, c)
	}
	r.Mu.RUnlock()
	for _, client := range clients {
		select {
		case client.Send <- fmt.Appendf(nil, `{"message": "%s","toClient":"%s"}`, msg.Message, msg.ClientId):
			// default:
			// 	r.Mu.Lock()
			// 	delete(r.Clients, client)
			// 	r.Mu.Unlock()
			// 	close(client.Send)
		}
	}
}

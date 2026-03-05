package constants

import (
	"sync"
)

type Meta struct {
	SenderId string  `json:"sender_id"`
	RoomId   string  `json:"room_id"`
	Time     float64 `json:"time"`
}
type BroadcastMessage struct {
	MessageType string `json:"message_type"`
	Meta        Meta   `json:"meta"`
	Data        any    `json:"data"`
}

func NewBroadcastMessage(messageType string, meta *Meta, data interface{}) *BroadcastMessage {
	return &BroadcastMessage{
		MessageType: messageType,
		Meta:        *meta,
		Data:        data,
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
		if msg.Meta.SenderId == c.Id {
			continue
		}
		clients = append(clients, c)
	}
	r.Mu.RUnlock()
	for _, client := range clients {
		client.Send <- msg
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
		client.Send <- msg
	}
}

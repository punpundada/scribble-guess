package websocket

import (
	"encoding/json"
	"log"
	"scribble-backend/internal/constants"
	"sync"
)

type Hub struct {
	Rooms     map[string]*constants.Room
	Broadcast chan []byte
	mu        sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{
		Rooms:     make(map[string]*constants.Room, 10),
		Broadcast: make(chan []byte),
	}
}

func (h *Hub) GetRoomFromId(roomId string) (*constants.Room, bool) {
	h.mu.RLock()
	room, exists := h.Rooms[roomId]
	h.mu.RUnlock()
	if !exists {
		return nil, false
	}
	return room, true
}

type Message struct {
	RoomId string `json:"roomId"`
}

func (h *Hub) Run() {

	for {
		msg := <-h.Broadcast
		var message Message
		err := json.Unmarshal(msg, &message)
		if err != nil {
			log.Println("Error unmarshalling message:", err)
			continue
		}
		log.Println("msg", message)
		room, exists := h.GetRoomFromId(message.RoomId)
		if !exists {
			log.Printf("Room with ID %s not found\n", message.RoomId)
			continue
		}
		room.BroadcastMessage(constants.Message(message))
	}
}

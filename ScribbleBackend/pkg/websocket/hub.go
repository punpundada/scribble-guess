package websocket

import (
	"scribble-backend/internal/constants"
	"sync"

	"github.com/google/uuid"
)

type Hub struct {
	Rooms map[string]*constants.Room
	// Broadcast chan []byte
	mu sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{
		Rooms: make(map[string]*constants.Room, 10),
		// Broadcast: make(chan []byte),
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

func (h *Hub) GetOrCreateRoom(roomId string) *constants.Room {
	id := roomId
	if len(id) == 0 {
		id = uuid.New().String()
	}
	room, ok := h.Rooms[id]
	if !ok {
		room = constants.NewRoom(id)
		h.Rooms[id] = room
		go room.Run()
	}
	return room
}

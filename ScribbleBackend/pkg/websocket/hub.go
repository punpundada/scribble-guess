package websocket

import (
	"math/rand/v2"
	"scribble-backend/internal/constants"
	"strconv"
	"sync"
	"time"
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
	seed := rand.NewPCG(uint64(time.Now().UnixNano()), 10)
	rnd := rand.New(seed)

	id := roomId
	if len(id) == 0 {
		id = strconv.Itoa(rnd.IntN(899999) + 100000)
	}
	room, ok := h.Rooms[id]
	if !ok {
		room = constants.NewRoom(id)
		h.Rooms[id] = room
		go room.Run()
	}
	return room
}

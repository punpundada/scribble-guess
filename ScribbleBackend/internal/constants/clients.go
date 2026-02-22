package constants

import "github.com/gorilla/websocket"

type Client struct {
	Conn   *websocket.Conn
	Send   chan []byte
	RoomID string
	Id     string
}

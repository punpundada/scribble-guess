package server

import (
	"net/http"
	"scribble-backend/internal/handler"
	m "scribble-backend/internal/middleware"
	"scribble-backend/pkg/websocket"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

type Server struct {
	Hub *websocket.Hub
}

func NewServer(hub *websocket.Hub) *Server {
	return &Server{
		Hub: hub,
	}
}

func (s *Server) NewRouter() http.Handler {
	router := chi.NewRouter()

	router.Use(m.RequestID)
	router.Use(m.Logger)
	// router.Use(middleware.Logger)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}))

	router.Get("/health", handler.Health)

	router.Route("/api", s.WebSocketRoutes)
	router.Route("/api/room", s.RoomRoutes)
	return router
}

func (s *Server) WebSocketRoutes(router chi.Router) {
	wsHandler := &handler.WS{Hub: s.Hub}
	router.Get("/ws", wsHandler.WebSocketHandler)
}

func (s *Server) RoomRoutes(router chi.Router) {
	roohHandler := &handler.RoomHandler{Hub: s.Hub}
	router.Post("/create", roohHandler.CreateRoom)
	router.Post("/delete", roohHandler.DeleteRoom)
}

package server

import (
	"net/http"
	"scribble-backend/internal/handler"
	"scribble-backend/internal/middleware"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func NewRouter() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}))

	router.Get("/health", handler.Health)

	router.Route("/api", func(api chi.Router) {
		api.Get("/ws", handler.WebSocketHandler)
	})

	return router
}

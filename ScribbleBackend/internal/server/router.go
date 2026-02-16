package server

import (
	"net/http"
	"scribble-backend/internal/handler"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func NewRouter() http.Handler {
	router := chi.NewRouter()

	// router.Use()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}))

	router.Get("/health", handler.Health)

	// 	r.Route("/api", func(api chi.Router) {
	// 	api.Get("/ws", handler.WebSocketHandler)
	// })

	return router
}

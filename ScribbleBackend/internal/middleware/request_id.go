package middleware

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"net/http"
)

type contextKey string

const RequestIDKey contextKey = "request_id"

func generateID() string {
	b := make([]byte, 12)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := generateID()

		ctx := context.WithValue(r.Context(), RequestIDKey, id)
		r = r.WithContext(ctx)

		w.Header().Set("X-Request-ID", id)

		next.ServeHTTP(w, r)
	})
}

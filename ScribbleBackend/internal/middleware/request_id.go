package middleware

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"scribble-backend/pkg/logger"
)

func generateID() string {
	b := make([]byte, 12)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		id := generateID()

		ctx := logger.WithRequestID(r.Context(), id)
		r = r.WithContext(ctx)

		w.Header().Set("X-Request-ID", id)

		next.ServeHTTP(w, r)
	})
}

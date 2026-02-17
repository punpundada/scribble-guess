package middleware

import (
	"net/http"
	"scribble-backend/pkg/logger"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		next.ServeHTTP(w, r)
		logger.FromCtx(r.Context()).
			Printf("%s %s %s",
				r.Method,
				r.URL.Path,
				time.Since(start),
			)
	})
}

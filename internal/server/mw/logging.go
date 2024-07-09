package mw

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

// wrappedWriter wraps an http.ResponseWriter to capture the status code.
type wrappedResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedResponseWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

// LoggingMiddleware logs the request method, path, and duration.
func LoggingMiddlerware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &wrappedResponseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}
		next.ServeHTTP(wrapped, r)
		slog.Info(fmt.Sprintf("%d %s %s %s", wrapped.statusCode, r.Method, r.URL.Path, time.Since(start)))
	})
}

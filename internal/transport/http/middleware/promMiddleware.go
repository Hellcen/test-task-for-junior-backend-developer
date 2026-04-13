package middleware

import (
	"example.com/taskservice/metrics"
	"net/http"
	"time"
)

func PromMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := &metrics.ResponseWriter{ResponseWriter: w, StatusCode: http.StatusOK}
		next(rw, r)

		duration := time.Since(start).Seconds()

		// Сохраняем метрики
		metrics.HttpRequestsTotal.WithLabelValues(r.Method, r.URL.Path, http.StatusText(rw.StatusCode)).Inc()
		metrics.HttpRequestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(duration)
	}
}

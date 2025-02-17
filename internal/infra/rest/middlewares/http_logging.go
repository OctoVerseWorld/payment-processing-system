package middlewares

import (
	"go.uber.org/zap"
	"net/http"
	"time"
)

// responseWriterWrapper ...
type responseWriterWrapper struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader ...
func (rw *responseWriterWrapper) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// HTTPLoggingMiddleware логирует HTTP-запросы + статус-код
func HTTPLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Оборачиваем ResponseWriter
		rw := &responseWriterWrapper{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(rw, r)

		duration := time.Since(start)

		zap.L().Info("HTTP Request",
			zap.Int("status", rw.statusCode),
			zap.String("method", r.Method),
			zap.String("url", r.URL.Path),
			zap.Duration("duration", duration),
		)
	})
}

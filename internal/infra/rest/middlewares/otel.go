package middlewares

import (
	"github.com/riandyrn/otelchi"
	"net/http"
)

// OTELTraceMiddleware is a middleware that adds OpenTelemetry tracing to the HTTP requests.
func OTELTraceMiddleware(next http.Handler) http.Handler {
	return otelchi.Middleware("octoverse-payments")(next)
}

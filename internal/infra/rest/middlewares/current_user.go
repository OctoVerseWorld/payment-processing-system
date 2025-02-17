package middlewares

import (
	"PaymentProcessingSystem/internal"
	"PaymentProcessingSystem/internal/configs"
	"PaymentProcessingSystem/internal/infra/rest/responses"
	"context"
	"net/http"
	"strconv"
)

type contextKey string
type headerName string

const (
	UserIDHeader headerName = "X-Current-User-ID"
	UserIDKey    contextKey = "currentUserID"
)

// CurrentUserMiddleware добавляет X-Current-User-ID в контекст запроса
func CurrentUserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg := configs.GetConfig()
		var userID int
		var err error
		userIDStr := r.Header.Get(string(UserIDHeader))

		if cfg.AppConfig.Environment.IsProduction() {
			if userIDStr == "" {
				responses.RenderErrorResponse(w, r,
					internal.NewErrorf(internal.UserIDNotProvided, "Header %s not provided", UserIDHeader))
				return
			}
			userID, err = strconv.Atoi(userIDStr)
			if err != nil {
				responses.RenderErrorResponse(w, r,
					internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "validation error"))
			}
		} else {
			if userIDStr == "" {
				userID = cfg.AppConfig.DefaultUserID
			}
		}

		ctx := context.WithValue(r.Context(), UserIDKey, userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

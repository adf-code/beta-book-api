package middleware

import (
	"github.com/adf-code/beta-book-api/internal/delivery/response"
	"github.com/rs/zerolog"
	"net/http"
	"strings"
)

func AuthMiddleware(logger zerolog.Logger) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				logger.Warn().Msg("‼️ Bearer token not found in header")
				response.Failed(w, 401, "authentication", "tryAuthentication", "Unauthorized")
				return
			}

			token := strings.TrimPrefix(authHeader, "Bearer ")
			if token != "connect123" {
				logger.Warn().Msg("‼️ Bearer token not authorized")
				response.Failed(w, 403, "authentication", "tryAuthentication", "Forbidden")
				return
			}

			next(w, r)
		}
	}
}

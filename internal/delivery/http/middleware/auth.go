package middleware

import (
	"beta-book-api/internal/delivery/response"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			response.Failed(w, 401, "authentication", "tryAuthentication", "Unauthorized")
			return
		}
		// You can do further token validation here
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token != "connect123" { // Replace this with real validation
			response.Failed(w, 403, "authentication", "tryAuthentication", "Forbidden")
			return
		}
		next.ServeHTTP(w, r)
	}
}

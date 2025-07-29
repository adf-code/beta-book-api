package middleware

import (
	"github.com/rs/zerolog"
	"net/http"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(logger zerolog.Logger) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
			next(rec, r)
			logger.Info().
				Str("method", r.Method).
				Str("path", r.URL.Path).
				Str("remote", r.RemoteAddr).
				Str("user_agent", r.UserAgent()).
				Int("status", rec.status).
				Msgf("📥 Incoming HTTP request, method: %s, path: %s, status: %d, remote: %s. user_agent: %s", r.Method, r.URL.Path, rec.status, r.RemoteAddr, r.UserAgent())

			//next(w, r)
		}
	}
}

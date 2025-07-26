package http

import (
	"beta-book-api/internal/delivery/http/book"
	"beta-book-api/internal/delivery/http/middleware"
	"beta-book-api/internal/pkg/logger"
	"beta-book-api/internal/repository"
	"beta-book-api/internal/usecase"
	"github.com/swaggo/http-swagger"
	"net/http"
	"strings"

	_ "beta-book-api/docs"
)

func SetupHandler(repo repository.BookRepository) http.Handler {
	telemetryLog := logger.InitLoggerWithTelemetry()

	bookUC := usecase.NewBookUseCase(repo)
	bookHandler := book.NewBookHandler(bookUC, telemetryLog)
	auth := middleware.AuthMiddleware

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		method := r.Method

		// ✅ Swagger
		if strings.HasPrefix(path, "/swagger/") {
			httpSwagger.WrapHandler(w, r)
			return
		}

		// ✅ All API prefixed with /api/v1
		if !strings.HasPrefix(path, "/api/v1") {
			http.NotFound(w, r)
			return
		}

		// Remove prefix for routing logic
		apiPath := strings.TrimPrefix(path, "/api/v1")

		switch {
		case apiPath == "/books" && method == http.MethodGet:
			auth(bookHandler.GetAll)(w, r)

		case strings.HasPrefix(apiPath, "/books/") && method == http.MethodGet:
			auth(bookHandler.GetByID)(w, r)

		case apiPath == "/books" && method == http.MethodPost:
			auth(bookHandler.Create)(w, r)

		case strings.HasPrefix(apiPath, "/books/") && method == http.MethodDelete:
			auth(bookHandler.Delete)(w, r)

		default:
			http.NotFound(w, r)
		}
	})
}

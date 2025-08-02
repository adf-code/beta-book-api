package http

import (
	"github.com/adf-code/beta-book-api/internal/delivery/http/book"
	"github.com/adf-code/beta-book-api/internal/delivery/http/health"
	"github.com/adf-code/beta-book-api/internal/delivery/http/middleware"
	"github.com/adf-code/beta-book-api/internal/delivery/http/router"
	"github.com/adf-code/beta-book-api/internal/usecase"
	"github.com/rs/zerolog"

	"github.com/swaggo/http-swagger"
	"net/http"
)

func SetupHandler(bookUC usecase.BookUseCase, bookCoverUC usecase.BookCoverUseCase, logger zerolog.Logger) http.Handler {
	bookHandler := book.NewBookHandler(bookUC, bookCoverUC, logger)
	healthHandler := health.NewHealthHandler(logger)
	auth := middleware.AuthMiddleware(logger)
	log := middleware.LoggingMiddleware(logger)

	r := router.NewRouter()

	r.HandlePrefix(http.MethodGet, "/swagger/", httpSwagger.WrapHandler)

	r.Handle("GET", "/healthz", middleware.Chain(log)(healthHandler.Check))

	r.Handle("GET", "/api/v1/books/cover/{id}", middleware.Chain(log, auth)(bookHandler.GetCoverByBookID))
	r.Handle("GET", "/api/v1/books/{id}", middleware.Chain(log, auth)(bookHandler.GetByID))
	r.Handle("GET", "/api/v1/books", middleware.Chain(log, auth)(bookHandler.GetAll))
	r.Handle("POST", "/api/v1/books/upload-cover", middleware.Chain(log, auth)(bookHandler.UploadCover))
	r.Handle("POST", "/api/v1/books", middleware.Chain(log, auth)(bookHandler.Create))
	r.Handle("DELETE", "/api/v1/books/{id}", middleware.Chain(log, auth)(bookHandler.Delete))

	return r
}

package http

import (
	"beta-book-api/internal/delivery/http/book"
	"beta-book-api/internal/delivery/http/middleware"
	"beta-book-api/internal/delivery/http/router"
	"beta-book-api/internal/usecase"
	"github.com/rs/zerolog"

	"github.com/swaggo/http-swagger"
	"net/http"
)

func SetupHandler(bookUC usecase.BookUseCase, bookCoverUC usecase.BookCoverUseCase, logger zerolog.Logger) http.Handler {
	bookHandler := book.NewBookHandler(bookUC, bookCoverUC, logger)
	auth := middleware.AuthMiddleware(logger)
	log := middleware.LoggingMiddleware(logger)

	r := router.NewRouter()

	// Swagger
	r.Handle(http.MethodGet, "/swagger/", httpSwagger.WrapHandler)

	// API Routes
	r.Handle(http.MethodGet, "/api/v1/books", middleware.Chain(log, auth)(bookHandler.GetAll))
	r.Handle(http.MethodGet, "/api/v1/books/cover/{id}", middleware.Chain(log, auth)(bookHandler.GetCoverByBookID))
	r.Handle(http.MethodGet, "/api/v1/books/{id}", middleware.Chain(log, auth)(bookHandler.GetByID))
	r.Handle(http.MethodPost, "/api/v1/books/upload-cover", middleware.Chain(log, auth)(bookHandler.UploadCover))
	r.Handle(http.MethodPost, "/api/v1/books", middleware.Chain(log, auth)(bookHandler.Create))
	r.Handle(http.MethodDelete, "/api/v1/books/{id}", middleware.Chain(log, auth)(bookHandler.Delete))

	return r
}

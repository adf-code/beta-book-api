package http

import (
	"beta-book-api/internal/delivery/http/book"
	"beta-book-api/internal/repository"
	"github.com/swaggo/http-swagger"
	"net/http"
	"strings"

	_ "beta-book-api/docs"
)

func SetupHandler(repo repository.BookRepository) http.Handler {
	bookHandler := book.NewBookHandler(repo)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		method := r.Method

		switch {
		// ✅ Swagger docs
		case strings.HasPrefix(path, "/swagger/"):
			httpSwagger.WrapHandler(w, r)

		// ✅ Get all books
		case path == "/books" && method == http.MethodGet:
			bookHandler.GetAll(w, r)

		// ✅ Get book by ID
		case strings.HasPrefix(path, "/books/") && method == http.MethodGet:
			bookHandler.GetByID(w, r)

		// ✅ Create book
		case path == "/books" && method == http.MethodPost:
			bookHandler.Create(w, r)

		// ✅ Delete book
		case strings.HasPrefix(path, "/books/") && method == http.MethodDelete:
			bookHandler.Delete(w, r)

		// ✅ Not found
		default:
			http.NotFound(w, r)
		}
	})
}

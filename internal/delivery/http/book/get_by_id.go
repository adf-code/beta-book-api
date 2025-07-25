package book

import (
	"beta-book-api/internal/delivery/response"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"net/http"
	"strings"
)

func (h *BookHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	if idStr == "" {
		response.Failed(w, 422, "books", "getBookByID", "Missing ID Parameter, Get Book by ID")
		return
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Failed(w, 422, "books", "getBookByID", "Invalid UUID, Get Book by ID")
		return
	}
	book, err := h.Repo.FetchByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.Success(w, 404, "books", "getBookByID", "Book not Found", nil)
			return
		}
		response.Failed(w, 500, "books", "getBookByID", "Error Get Book by ID")
		return
	}
	response.Success(w, 200, "books", "getBookByID", "Success Get Book by ID", book)
}

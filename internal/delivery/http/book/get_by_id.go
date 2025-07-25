package book

import (
	"beta-book-api/internal/delivery/response"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"net/http"
	"strings"
)

// GetBookByID godoc
// @Summary      Get book by ID
// @Description  Retrieve a book entity using its UUID
// @Tags         books
// @Security     BearerAuth
// @Param        id   path      string  true  "UUID of the book"
// @Success      200  {object}  response.APIResponse
// @Failure      400  {object}  response.APIResponse  "Invalid UUID"
// @Failure      401  {object}  response.APIResponse  "Unauthorized"
// @Failure      404  {object}  response.APIResponse  "Book not found"
// @Failure      500  {object}  response.APIResponse  "Internal server error"
// @Router       /books/{id} [get]
func (h *BookHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/books/")
	if idStr == "" {
		response.Failed(w, 422, "books", "getBookByID", "Missing ID Parameter, Get Book by ID")
		return
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Failed(w, 422, "books", "getBookByID", "Invalid UUID, Get Book by ID")
		return
	}
	book, err := h.UseCase.GetByID(id)
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

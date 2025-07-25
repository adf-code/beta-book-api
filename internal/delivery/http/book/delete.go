package book

import (
	"beta-book-api/internal/delivery/response"
	"github.com/google/uuid"
	"net/http"
	"strings"
)

func (h *BookHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	if idStr == "" {
		response.Failed(w, 422, "books", "deleteBookByID", "Missing ID Parameter, Delete Book by ID")
		return
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Failed(w, 422, "books", "deleteBookByID", "Invalid UUID, Delete Book by ID")
		return
	}
	if err := h.Repo.Remove(id); err != nil {
		response.Failed(w, 500, "books", "deleteBookByID", "Error Delete Book")
		return
	}
	response.Success(w, 202, "books", "deleteBookByID", "Success Delete Book", nil)
}

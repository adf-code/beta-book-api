package book

import (
	"beta-book-api/internal/delivery/response"
	"beta-book-api/internal/entity"
	"encoding/json"
	"net/http"
)

func (h *BookHandler) Create(w http.ResponseWriter, r *http.Request) {
	var book entity.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		response.Failed(w, 422, "books", "createBook", "Invalid Data, Create Book")
		return
	}

	if err := h.Repo.Store(&book); err != nil {
		response.Failed(w, 500, "books", "createBook", "Error Create Book")
		return
	}
	response.Success(w, 201, "books", "createBook", "Success Create Book", book)
}

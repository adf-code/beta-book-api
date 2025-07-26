package book

import (
	"beta-book-api/internal/delivery/response"
	"beta-book-api/internal/entity"
	"encoding/json"
	"net/http"
)

// CreateBook godoc
// @Summary      Create a new book
// @Description  Creates a new book with the provided title, author, and year
// @Tags         books
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request  body      entity.Book  true  "Book data to create"
// @Success      201      {object}  response.APIResponse
// @Failure      400      {object}  response.APIResponse
// @Failure      401      {object}  response.APIResponse
// @Failure      422      {object}  response.APIResponse
// @Failure      500      {object}  response.APIResponse
// @Router       /books [post]
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

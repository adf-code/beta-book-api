package book

import (
	"encoding/json"
	"fmt"
	"github.com/adf-code/beta-book-api/internal/delivery/response"
	"github.com/adf-code/beta-book-api/internal/entity"
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
// @Router       /api/v1/books [post]
func (h *BookHandler) Create(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info().Msg("📥 Incoming Create request")
	var book entity.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		h.Logger.Error().Err(err).Msg("❌ Failed to store book, invalid data")
		response.Failed(w, 422, "books", "createBook", "Invalid Data, Create Book")
		return
	}

	newBook, err := h.BookUC.Create(r.Context(), book)
	if err != nil {
		h.Logger.Error().Err(err).Msg("❌ Failed to store book, general")
		response.Failed(w, 500, "books", "createBook", "Error Create Book")
		return
	}
	newBook.BookCover = make([]entity.BookCover, 0)
	h.Logger.Info().Str("data", fmt.Sprint(newBook)).Msg("✅ Successfully stored book")
	response.Success(w, 201, "books", "createBook", "Success Create Book", newBook)
}

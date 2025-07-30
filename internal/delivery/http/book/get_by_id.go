package book

import (
	"beta-book-api/internal/delivery/http/router"
	"beta-book-api/internal/delivery/response"
	"beta-book-api/internal/entity"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"net/http"
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
	h.Logger.Info().Msg("📥 Incoming GetByID request")

	idStr := router.GetParam(r, "id")
	if idStr == "" {
		h.Logger.Error().Msg("❌ Failed to get book by ID, missing ID parameter")
		response.Failed(w, 422, "books", "getBookByID", "Missing ID Parameter, Get Book by ID")
		return
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		h.Logger.Error().Err(err).Msg("❌ Failed to get book by ID, invalid UUID parameter")
		response.Failed(w, 422, "books", "getBookByID", "Invalid UUID, Get Book by ID")
		return
	}
	book, err := h.BookUC.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			h.Logger.Info().Msg("✅ Successfully get book by id, data not found")
			response.Success(w, 404, "books", "getBookByID", "Book not Found", nil)
			return
		}
		h.Logger.Error().Err(err).Msg("❌ Failed to get book by ID, general")
		response.Failed(w, 500, "books", "getBookByID", "Error Get Book by ID")
		return
	}
	book.BookCover = make([]entity.BookCover, 0)
	h.Logger.Info().Str("data", fmt.Sprint(book.ID)).Msg("✅ Successfully get book by id")
	response.Success(w, 200, "books", "getBookByID", "Success Get Book by ID", book)
}

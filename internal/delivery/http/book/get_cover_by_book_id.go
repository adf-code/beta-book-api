package book

import (
	"beta-book-api/internal/delivery/response"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"strings"
)

// GetCoverByBookID godoc
// @Summary      Get cover by book ID
// @Description  Retrieve a book cover entity using book UUID
// @Tags         books
// @Security     BearerAuth
// @Param        id   path      string  true  "UUID of the book"
// @Success      200  {object}  response.APIResponse
// @Failure      400  {object}  response.APIResponse  "Invalid UUID"
// @Failure      401  {object}  response.APIResponse  "Unauthorized"
// @Failure      404  {object}  response.APIResponse  "Book not found"
// @Failure      500  {object}  response.APIResponse  "Internal server error"
// @Router       /books/cover/{id} [get]
func (h *BookHandler) GetCoverByBookID(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info().Msg("📥 Incoming GetByID request")
	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/books/cover/")
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
	booksCover, err := h.BookCoverUC.GetByBookID(r.Context(), id)
	if err != nil {
		h.Logger.Error().Err(err).Msg("❌ Failed to fetch books cover, general")
		response.FailedWithMeta(w, 500, "books", "getAllBooks", "Error Get Book Cover by Book ID", nil)
		return
	}
	book.BookCover = booksCover
	h.Logger.Info().Str("data", fmt.Sprint(book.ID)).Msg("✅ Successfully get book by id")
	response.Success(w, 200, "books", "getBookByID", "Success Get Book by ID", book)
}

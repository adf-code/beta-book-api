package book

import (
	"beta-book-api/internal/delivery/response"
	"github.com/google/uuid"
	"net/http"
	"strings"
)

// DeleteBookByID godoc
// @Summary      Delete a book by ID
// @Description  Deletes a book entity using its UUID
// @Tags         books
// @Security     BearerAuth
// @Param        id   path      string  true  "UUID of the book to delete"
// @Success      202  {object}  response.APIResponse
// @Failure      400  {object}  response.APIResponse  "Invalid UUID"
// @Failure      401  {object}  response.APIResponse  "Unauthorized"
// @Failure      404  {object}  response.APIResponse  "Book not found"
// @Failure      500  {object}  response.APIResponse  "Internal server error"
// @Router       /books/{id} [delete]
func (h *BookHandler) Delete(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info().Msg("📥 Incoming Delete request")
	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/books/")
	if idStr == "" {
		h.Logger.Error().Msg("❌ Failed to remove book, missing ID parameter")
		response.Failed(w, 422, "books", "deleteBookByID", "Missing ID Parameter, Delete Book by ID")
		return
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		h.Logger.Error().Err(err).Msg("❌ Failed to remove book, invalid UUID parameter")
		response.Failed(w, 422, "books", "deleteBookByID", "Invalid UUID, Delete Book by ID")
		return
	}
	if err := h.UseCase.Delete(id); err != nil {
		h.Logger.Error().Err(err).Msg("❌ Failed to remove book, general")
		response.Failed(w, 500, "books", "deleteBookByID", "Error Delete Book")
		return
	}
	h.Logger.Info().Msg("✅ Successfully removed book")
	response.Success(w, 202, "books", "deleteBookByID", "Success Delete Book", nil)
}

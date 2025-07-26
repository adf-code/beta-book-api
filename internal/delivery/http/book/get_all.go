package book

import (
	"beta-book-api/internal/delivery/request"
	"beta-book-api/internal/delivery/response"
	"net/http"
)

// GetAllBooks godoc
// @Summary      Get list of books
// @Description  List all books with filter, search, pagination
// @Tags         books
// @Accept       json
// @Produce      json
//
// --- Search Query ---
// @Param        search_field      query    string   false  "Search field (e.g., title)"
// @Param        search_value      query    string   false  "Search value (e.g., golang)"
//
// --- Filter Search Query ---
// @Param filter_field query []string false "Filter field" collectionFormat(multi) explode(true)
// @Param filter_value query []string false "Filter value" collectionFormat(multi) explode(true)
//
// --- Range Query ---
// @Param range_field query []string false "Range field" collectionFormat(multi) explode(true)
// @Param from        query []string false "Range lower bound" collectionFormat(multi) explode(true)
// @Param to          query []string false "Range upper bound" collectionFormat(multi) explode(true)
//
// --- Pagination & Sort ---
// @Param        sort_field        query    string   false  "Sort field"
// @Param        sort_direction    query    string   false  "Sort direction ASC/DESC"
// @Param        page              query    int      false  "Page number"
// @Param        per_page          query    int      false  "Limit per page"
//
// @Security     BearerAuth
//
// @Success      200     {object}  response.APIResponse
// @Failure      500     {object}  response.APIResponse
// @Router       /books [get]
func (h *BookHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	params := request.ParseBookQueryParams(r)
	books, err := h.UseCase.GetAll(params)
	if err != nil {
		response.FailedWithMeta(w, 500, "books", "getAllBooks", "Error Get All Books", nil)
		return
	}
	response.SuccessWithMeta(w, 200, "books", "getAllBooks", "Success Get All Books", &params, books)
}

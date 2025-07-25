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
// @Param        search_field      query    string   false  "Search field (e.g., title)"
// @Param        search_value      query    string   false  "Search value (e.g., golang)"
// @Param        filter_author[]   query    []string false  "Filter by author"
// @Param        filter_title[] query    []string false  "Filter by title"
// @Param        range_year_min    query    int      false  "Min year"
// @Param        range_year_max    query    int      false  "Max year"
// @Param        range_created_from query   string   false  "Start date (RFC3339)"
// @Param        range_created_to   query   string   false  "End date (RFC3339)"
// @Param        sort_field        query    string   false  "Sort field"
// @Param        sort_direction    query    string   false  "Sort direction ASC/DESC"
// @Param        page              query    int      false  "Page number"
// @Param        per_page          query    int      false  "Limit per page"
// @Success      200     {object}  response.APIResponse
// @Failure      500     {object}  response.APIResponse
// @Router       /api/v1/books [get]
func (h *BookHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	params := request.ParseBookQueryParams(r)
	books, err := h.Repo.FetchWithQueryParams(params)
	if err != nil {
		response.FailedWithMeta(w, 500, "books", "getAllBooks", "Error Get All Books", nil)
		return
	}
	response.SuccessWithMeta(w, 200, "books", "getAllBooks", "Success Get All Books", &params, books)
}

package request

import (
	"net/http"
	"strconv"
	"strings"
)

func ParseBookQueryParams(r *http.Request) BookListQueryParams {
	q := r.URL.Query()

	// Search
	searchField := q.Get("search_field")
	searchValue := q.Get("search_value")

	// Filters
	var filters []QueryFilter
	filterFields := q["filter_field"]
	filterValues := q["filter_value"]
	for i := 0; i < len(filterFields) && i < len(filterValues); i++ {
		values := strings.Split(filterValues[i], ",")
		filters = append(filters, QueryFilter{
			Field: filterFields[i],
			Value: values,
		})
	}

	// Range
	var ranges []QueryRange
	rangeFields := q["range_field"]
	froms := q["from"]
	tos := q["to"]

	for i := 0; i < len(rangeFields); i++ {
		rng := QueryRange{Field: rangeFields[i]}
		if i < len(froms) {
			rng.From = &froms[i]
		}
		if i < len(tos) {
			rng.To = &tos[i]
		}
		ranges = append(ranges, rng)
	}

	// Sort
	sortField := q.Get("sort_field")
	sortDir := strings.ToUpper(q.Get("sort_direction"))

	// Pagination
	page, _ := strconv.Atoi(q.Get("page"))
	per_page, _ := strconv.Atoi(q.Get("per_page"))
	if page <= 0 {
		page = 1
	}
	if per_page <= 0 {
		per_page = 10
	}

	return BookListQueryParams{
		SearchField: searchField,
		SearchValue: searchValue,
		Filter:      filters,
		Range:       ranges,
		SortField:   sortField,
		SortDir:     sortDir,
		Page:        page,
		PerPage:     per_page,
	}
}

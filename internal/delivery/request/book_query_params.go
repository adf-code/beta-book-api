package request

type QueryFilter struct {
	Field string   `json:"field"`
	Value []string `json:"value"`
}

type QueryRange struct {
	Field string  `json:"field"`
	From  *string `json:"from"`
	To    *string `json:"to"`
}

type BookListQueryParams struct {
	SearchField string        `json:"search_field"`
	SearchValue string        `json:"search_value"`
	Filter      []QueryFilter `json:"filter"`
	Range       []QueryRange  `json:"range"`
	SortField   string        `json:"sort_field"`
	SortDir     string        `json:"sort_dir"`
	Page        int           `json:"page"`
	PerPage     int           `json:"per_page"`
}

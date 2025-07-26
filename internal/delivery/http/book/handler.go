package book

import (
	"beta-book-api/internal/usecase"
)

type BookHandler struct {
	UseCase usecase.BookUseCase
}

func NewBookHandler(uc usecase.BookUseCase) *BookHandler {
	return &BookHandler{UseCase: uc}
}

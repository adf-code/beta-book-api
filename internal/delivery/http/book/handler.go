package book

import (
	"beta-book-api/internal/usecase"
	"github.com/rs/zerolog"
)

type BookHandler struct {
	UseCase usecase.BookUseCase
	Logger  zerolog.Logger
}

func NewBookHandler(uc usecase.BookUseCase, logger zerolog.Logger) *BookHandler {
	return &BookHandler{UseCase: uc, Logger: logger}
}

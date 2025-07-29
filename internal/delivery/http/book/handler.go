package book

import (
	"beta-book-api/internal/usecase"
	"github.com/rs/zerolog"
)

type BookHandler struct {
	BookUC      usecase.BookUseCase
	BookCoverUC usecase.BookCoverUseCase
	Logger      zerolog.Logger
	//EmailClient *mail.SendGridClient
}

func NewBookHandler(bookUC usecase.BookUseCase, bookCoverUC usecase.BookCoverUseCase, logger zerolog.Logger) *BookHandler {
	return &BookHandler{BookUC: bookUC, BookCoverUC: bookCoverUC, Logger: logger}
}

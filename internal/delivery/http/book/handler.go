package book

import (
	email "beta-book-api/internal/pkg/mail"
	"beta-book-api/internal/usecase"
	"github.com/rs/zerolog"
)

type BookHandler struct {
	UseCase     usecase.BookUseCase
	Logger      zerolog.Logger
	EmailClient email.SendGridClient
}

func NewBookHandler(uc usecase.BookUseCase, logger zerolog.Logger, emailClient email.SendGridClient) *BookHandler {
	return &BookHandler{UseCase: uc, Logger: logger, EmailClient: emailClient}
}

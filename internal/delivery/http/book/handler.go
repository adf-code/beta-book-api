package book

import (
	"beta-book-api/internal/repository"
)

type BookHandler struct {
	Repo repository.BookRepository
}

func NewBookHandler(repo repository.BookRepository) *BookHandler {
	return &BookHandler{Repo: repo}
}

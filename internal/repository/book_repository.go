package repository

import (
	"beta-book-api/internal/delivery/request"
	"beta-book-api/internal/entity"
	"github.com/google/uuid"
)

type BookRepository interface {
	FetchAll() ([]entity.Book, error)
	FetchWithQueryParams(params request.BookListQueryParams) ([]entity.Book, error)
	FetchByID(id uuid.UUID) (*entity.Book, error)
	Store(book *entity.Book) error
	Remove(id uuid.UUID) error
}

package repository

import (
	"beta-book-api/internal/delivery/request"
	"beta-book-api/internal/entity"
	"database/sql"
	"github.com/google/uuid"
)

type BookRepository interface {
	FetchWithQueryParams(params request.BookListQueryParams) ([]entity.Book, error)
	FetchByID(id uuid.UUID) (*entity.Book, error)
	Store(tx *sql.Tx, book *entity.Book) error
	Remove(id uuid.UUID) error
}

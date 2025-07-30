package mocks

import (
	"beta-book-api/internal/delivery/request"
	"beta-book-api/internal/entity"
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type BookRepository struct {
	mock.Mock
}

func (m *BookRepository) FetchWithQueryParams(ctx context.Context, params request.BookListQueryParams) ([]entity.Book, error) {
	args := m.Called(ctx, params)
	return args.Get(0).([]entity.Book), args.Error(1)
}

func (m *BookRepository) FetchByID(ctx context.Context, id uuid.UUID) (*entity.Book, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entity.Book), args.Error(1)
}

func (m *BookRepository) Store(ctx context.Context, tx *sql.Tx, book *entity.Book) error {
	args := m.Called(ctx, tx, book)
	return args.Error(0)
}

func (m *BookRepository) Remove(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

package mocks

import (
	"beta-book-api/internal/entity"
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type BookCoverRepository struct {
	mock.Mock
}

func (m *BookCoverRepository) Store(ctx context.Context, tx *sql.Tx, cover *entity.BookCover) error {
	args := m.Called(ctx, tx, cover)
	return args.Error(0)
}

func (m *BookCoverRepository) FetchByBookID(ctx context.Context, bookID uuid.UUID) ([]entity.BookCover, error) {
	args := m.Called(ctx, bookID)
	return args.Get(0).([]entity.BookCover), args.Error(1)
}

package mocks

import (
	"beta-book-api/internal/delivery/request"
	"beta-book-api/internal/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type BookRepository struct {
	mock.Mock
}

func (m *BookRepository) FetchWithQueryParams(params request.BookListQueryParams) ([]entity.Book, error) {
	args := m.Called(params)
	return args.Get(0).([]entity.Book), args.Error(1)
}

func (m *BookRepository) FetchByID(id uuid.UUID) (*entity.Book, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Book), args.Error(1)
}

func (m *BookRepository) Store(book *entity.Book) error {
	args := m.Called(book)
	return args.Error(0)
}

func (m *BookRepository) Remove(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}

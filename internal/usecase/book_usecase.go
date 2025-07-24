package usecase

import (
	"beta-book-api/internal/entity"
	"github.com/google/uuid"
)

type BookUseCase interface {
	GetAll() ([]entity.Book, error)
	GetByID(id uuid.UUID) (*entity.Book, error)
	Create(book entity.Book) (*entity.Book, error)
	Delete(id uuid.UUID) error
}

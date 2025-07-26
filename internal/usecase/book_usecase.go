package usecase

import (
	"beta-book-api/internal/delivery/request"
	"beta-book-api/internal/entity"
	"beta-book-api/internal/pkg/logger"
	"beta-book-api/internal/repository"
	"github.com/google/uuid"
)

type BookUseCase interface {
	GetAll(params request.BookListQueryParams) ([]entity.Book, error)
	GetByID(id uuid.UUID) (*entity.Book, error)
	Create(book entity.Book) (*entity.Book, error)
	Delete(id uuid.UUID) error
}

type bookUseCase struct {
	repo repository.BookRepository
}

func NewBookUseCase(repo repository.BookRepository) BookUseCase {
	return &bookUseCase{repo: repo}
}

func (uc *bookUseCase) GetAll(params request.BookListQueryParams) ([]entity.Book, error) {
	logger.Log.Info().Str("usecase", "GetAll").Msg("Fetching books")
	return uc.repo.FetchWithQueryParams(params)
}

func (uc *bookUseCase) GetByID(id uuid.UUID) (*entity.Book, error) {
	logger.Log.Info().Str("usecase", "GetByID").Msg("Fetching book by ID")
	return uc.repo.FetchByID(id)
}

func (uc *bookUseCase) Create(book entity.Book) (*entity.Book, error) {
	logger.Log.Info().Str("usecase", "Create").Msg("Store book")
	err := uc.repo.Store(&book)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (uc *bookUseCase) Delete(id uuid.UUID) error {
	logger.Log.Info().Str("usecase", "Delete").Msg("Remove book")
	return uc.repo.Remove(id)
}

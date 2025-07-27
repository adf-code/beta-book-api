package usecase

import (
	"beta-book-api/internal/delivery/request"
	"beta-book-api/internal/entity"
	"beta-book-api/internal/repository"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type BookUseCase interface {
	GetAll(params request.BookListQueryParams) ([]entity.Book, error)
	GetByID(id uuid.UUID) (*entity.Book, error)
	Create(book entity.Book) (*entity.Book, error)
	Delete(id uuid.UUID) error
}

type bookUseCase struct {
	repo   repository.BookRepository
	logger zerolog.Logger
}

func NewBookUseCase(repo repository.BookRepository, logger zerolog.Logger) BookUseCase {
	return &bookUseCase{
		repo:   repo,
		logger: logger,
	}
}

func (uc *bookUseCase) GetAll(params request.BookListQueryParams) ([]entity.Book, error) {
	uc.logger.Info().Str("usecase", "GetAll").Msg("⚙️ Fetching all books")
	return uc.repo.FetchWithQueryParams(params)
}

func (uc *bookUseCase) GetByID(id uuid.UUID) (*entity.Book, error) {
	uc.logger.Info().Str("usecase", "GetByID").Msg("⚙️ Fetching book by ID")
	return uc.repo.FetchByID(id)
}

func (uc *bookUseCase) Create(book entity.Book) (*entity.Book, error) {
	uc.logger.Info().Str("usecase", "Create").Msg("⚙️ Store book")
	err := uc.repo.Store(&book)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (uc *bookUseCase) Delete(id uuid.UUID) error {
	uc.logger.Info().Str("usecase", "Delete").Msg("⚙️ Remove book")
	return uc.repo.Remove(id)
}

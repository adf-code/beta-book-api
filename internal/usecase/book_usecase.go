package usecase

import (
	"beta-book-api/internal/delivery/request"
	"beta-book-api/internal/entity"
	"beta-book-api/internal/pkg/mail"
	"beta-book-api/internal/repository"
	"database/sql"
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
	repo        repository.BookRepository
	db          *sql.DB
	logger      zerolog.Logger
	emailClient mail.EmailClient
}

func NewBookUseCase(repo repository.BookRepository, db *sql.DB, logger zerolog.Logger, emailClient mail.EmailClient) BookUseCase {
	return &bookUseCase{
		repo:        repo,
		db:          db,
		logger:      logger,
		emailClient: emailClient,
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
	tx, err := uc.db.Begin()
	if err != nil {
		uc.logger.Error().Err(err).Msg("❌ Failed to begin transaction")
		return nil, err
	}

	err = uc.repo.Store(tx, &book)
	if err != nil {
		tx.Rollback()
		uc.logger.Error().Err(err).Msg("❌ Failed to store book, rolling back")
		return nil, err
	}

	err = uc.emailClient.SendBookCreatedEmail(book) // custom wrapper
	if err != nil {
		tx.Rollback()
		uc.logger.Error().Err(err).Msg("❌ Failed to send email, rolling back")
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		uc.logger.Error().Err(err).Msg("❌ Failed to commit transaction")
		return nil, err
	}

	uc.logger.Info().Str("book_id", book.ID.String()).Msg("✅ Book created and email sent successfully")
	return &book, nil
}

func (uc *bookUseCase) Delete(id uuid.UUID) error {
	uc.logger.Info().Str("usecase", "Delete").Msg("⚙️ Remove book")
	return uc.repo.Remove(id)
}

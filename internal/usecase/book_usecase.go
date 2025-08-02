package usecase

import (
	"context"
	"database/sql"
	"github.com/adf-code/beta-book-api/internal/delivery/request"
	"github.com/adf-code/beta-book-api/internal/entity"
	"github.com/adf-code/beta-book-api/internal/pkg/mail"
	"github.com/adf-code/beta-book-api/internal/repository"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type BookUseCase interface {
	GetAll(ctx context.Context, params request.BookListQueryParams) ([]entity.Book, error)
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Book, error)
	Create(ctx context.Context, book entity.Book) (*entity.Book, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type bookUseCase struct {
	bookRepo    repository.BookRepository
	db          *sql.DB
	logger      zerolog.Logger
	emailClient mail.EmailClient
}

func NewBookUseCase(bookRepo repository.BookRepository, db *sql.DB, logger zerolog.Logger, emailClient mail.EmailClient) BookUseCase {
	return &bookUseCase{
		bookRepo:    bookRepo,
		db:          db,
		logger:      logger,
		emailClient: emailClient,
	}
}

func (uc *bookUseCase) GetAll(ctx context.Context, params request.BookListQueryParams) ([]entity.Book, error) {
	uc.logger.Info().Str("usecase", "GetAll").Msg("⚙️ Fetching all books")
	return uc.bookRepo.FetchWithQueryParams(ctx, params)
}

func (uc *bookUseCase) GetByID(ctx context.Context, id uuid.UUID) (*entity.Book, error) {
	uc.logger.Info().Str("usecase", "GetByID").Msg("⚙️ Fetching book by ID")
	return uc.bookRepo.FetchByID(ctx, id)
}

func (uc *bookUseCase) Create(ctx context.Context, book entity.Book) (*entity.Book, error) {
	uc.logger.Info().Str("usecase", "Create").Msg("⚙️ Store book")
	tx, err := uc.db.Begin()
	if err != nil {
		uc.logger.Error().Err(err).Msg("❌ Failed to begin transaction")
		return nil, err
	}

	err = uc.bookRepo.Store(ctx, tx, &book)
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

func (uc *bookUseCase) Delete(ctx context.Context, id uuid.UUID) error {
	uc.logger.Info().Str("usecase", "Delete").Msg("⚙️ Remove book")
	return uc.bookRepo.Remove(ctx, id)
}

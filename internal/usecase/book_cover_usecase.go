package usecase

import (
	"beta-book-api/internal/entity"
	"beta-book-api/internal/pkg/object_storage"
	"beta-book-api/internal/repository"
	"context"
	"database/sql"
	"github.com/rs/zerolog"
	"mime/multipart"
	"strings"
	"time"

	"fmt"
	"github.com/google/uuid"
)

type BookCoverUseCase interface {
	Upload(ctx context.Context, bookID uuid.UUID, file multipart.File, fileHeader *multipart.FileHeader) (*entity.BookCover, error)
	GetByBookID(ctx context.Context, bookID uuid.UUID) ([]entity.BookCover, error)
}

type bookCoverUseCase struct {
	bookCoverRepo repository.BookCoverRepository
	db            *sql.DB
	logger        zerolog.Logger
	objectStorage *object_storage.MinioClient
}

func NewBookCoverUseCase(bookCoverRepo repository.BookCoverRepository, db *sql.DB, logger zerolog.Logger, storage *object_storage.MinioClient) BookCoverUseCase {
	return &bookCoverUseCase{
		bookCoverRepo: bookCoverRepo,
		db:            db,
		logger:        logger,
		objectStorage: storage,
	}
}

func (uc *bookCoverUseCase) Upload(ctx context.Context, bookID uuid.UUID, file multipart.File, fileHeader *multipart.FileHeader) (*entity.BookCover, error) {
	uc.logger.Info().Str("usecase", "UploadCover").Msg("⚙️ Upload book cover")
	timestamp := time.Now().Format(time.RFC3339)
	sanitizedTimestamp := strings.ReplaceAll(timestamp, ":", "-") // replace ":" with "-" to avoid path issues
	objectName := fmt.Sprintf("covers/book_%s_%s_%s", bookID, sanitizedTimestamp, fileHeader.Filename)

	url, err := uc.objectStorage.UploadFile(ctx, file, objectName, fileHeader.Size, fileHeader.Header.Get("Content-Type"))
	if err != nil {
		uc.logger.Error().Err(err).Msg("❌ Failed upload file to object storage")
		return nil, err
	}

	cover := entity.BookCover{
		BookID:   bookID,
		FileName: fileHeader.Filename,
		FileURL:  url,
	}

	tx, err := uc.db.Begin()
	if err != nil {
		uc.logger.Error().Err(err).Msg("❌ Failed to begin transaction")
		return nil, err
	}

	err = uc.bookCoverRepo.Store(ctx, tx, &cover)
	if err != nil {
		tx.Rollback()
		uc.logger.Error().Err(err).Msg("❌ Failed to store book, rolling back")
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		uc.logger.Error().Err(err).Msg("❌ Failed to commit transaction")
		return nil, err
	}

	uc.logger.Info().Str("book_cover_id", cover.ID.String()).Msg("✅ Book Cover uploaded successfully")
	return &cover, nil
}

func (uc *bookCoverUseCase) GetByBookID(ctx context.Context, bookID uuid.UUID) ([]entity.BookCover, error) {
	return uc.bookCoverRepo.FetchByBookID(ctx, bookID)
}

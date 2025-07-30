package usecase_test

import (
	"beta-book-api/internal/entity"
	objectMocks "beta-book-api/internal/pkg/object_storage/mocks"
	repoMocks "beta-book-api/internal/repository/mocks"
	"beta-book-api/internal/usecase"
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"mime/multipart"
	"os"
	"testing"
)

func TestUploadBookCover(t *testing.T) {
	db, sqlMock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	sqlMock.ExpectBegin()
	sqlMock.ExpectCommit()

	mockRepo := new(repoMocks.BookCoverRepository)
	mockStorage := new(objectMocks.MinioClient)
	logger := zerolog.Nop()

	bookID := uuid.New()
	fileContent := []byte("fake image content")

	// Buat file temporer
	tmpFile, err := os.CreateTemp("", "cover*.jpg")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	_, err = tmpFile.Write(fileContent)
	assert.NoError(t, err)
	_, err = tmpFile.Seek(0, 0)
	assert.NoError(t, err)

	file := multipart.File(tmpFile)
	fileHeader := &multipart.FileHeader{
		Filename: "cover.jpg",
		Size:     int64(len(fileContent)),
		Header:   map[string][]string{"Content-Type": {"image/jpeg"}},
	}

	mockStorage.On("UploadFile", mock.Anything, mock.Anything, mock.AnythingOfType("string"), int64(len(fileContent)), "image/jpeg").Return("http://example.com/cover.jpg", nil)
	mockRepo.On("Store", mock.Anything, mock.AnythingOfType("*sql.Tx"), mock.AnythingOfType("*entity.BookCover")).Return(nil)

	bookCoverUC := usecase.NewBookCoverUseCase(mockRepo, db, logger, mockStorage)
	result, err := bookCoverUC.Upload(context.TODO(), bookID, file, fileHeader)

	assert.NoError(t, err)
	assert.Equal(t, "cover.jpg", result.FileName)
	assert.Equal(t, "http://example.com/cover.jpg", result.FileURL)
	assert.Equal(t, bookID, result.BookID)
	assert.NoError(t, sqlMock.ExpectationsWereMet())

	mockRepo.AssertExpectations(t)
	mockStorage.AssertExpectations(t)
}

func TestGetBookCoversByBookID(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mockRepo := new(repoMocks.BookCoverRepository)
	mockStorage := new(objectMocks.MinioClient)
	logger := zerolog.Nop()

	bookID := uuid.New()
	expected := []entity.BookCover{
		{ID: uuid.New(), BookID: bookID, FileName: "cover.jpg", FileURL: "http://example.com/cover.jpg"},
	}

	mockRepo.On("FetchByBookID", mock.Anything, bookID).Return(expected, nil)

	bookCoverUC := usecase.NewBookCoverUseCase(mockRepo, db, logger, mockStorage)
	result, err := bookCoverUC.GetByBookID(context.TODO(), bookID)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

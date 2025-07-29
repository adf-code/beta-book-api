package usecase_test

import (
	"beta-book-api/internal/delivery/request"
	"beta-book-api/internal/entity"
	mailMocks "beta-book-api/internal/pkg/mail/mocks"
	repoMocks "beta-book-api/internal/repository/mocks"
	"beta-book-api/internal/usecase"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/rs/zerolog"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func newTestBookUseCase() usecase.BookUseCase {
	mockRepo := new(repoMocks.BookRepository)
	db := &sql.DB{}
	logger := zerolog.Nop()
	emailClient := new(mailMocks.SendGridClient)
	return usecase.NewBookUseCase(mockRepo, db, logger, emailClient)
}

func TestGetAllBooks(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mockRepo := new(repoMocks.BookRepository)
	mockEmail := new(mailMocks.SendGridClient)
	logger := zerolog.Nop()

	bookUC := usecase.NewBookUseCase(mockRepo, db, logger, mockEmail)

	expected := []entity.Book{
		{ID: uuid.New(), Title: "Go Programming", Author: "Alice", Year: 2020},
	}

	mockRepo.On("FetchWithQueryParams", mock.Anything).Return(expected, nil)
	result, err := bookUC.GetAll(request.BookListQueryParams{})

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestGetBookByID(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mockRepo := new(repoMocks.BookRepository)
	mockEmail := new(mailMocks.SendGridClient)
	logger := zerolog.Nop()

	bookUC := usecase.NewBookUseCase(mockRepo, db, logger, mockEmail)

	id := uuid.New()
	expectedBook := &entity.Book{ID: id, Title: "Clean Code", Author: "Robert C. Martin", Year: 2008}

	mockRepo.On("FetchByID", id).Return(expectedBook, nil)

	result, err := bookUC.GetByID(id)

	assert.NoError(t, err)
	assert.Equal(t, expectedBook, result)
	mockRepo.AssertExpectations(t)
}

func TestCreateBook(t *testing.T) {
	// Step 1: Setup mock DB & transaction
	db, sqlMock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	sqlMock.ExpectBegin()
	sqlMock.ExpectCommit()

	// Step 2: Mock repository & email
	mockRepo := new(repoMocks.BookRepository)
	mockEmail := new(mailMocks.SendGridClient)
	logger := zerolog.Nop()

	book := entity.Book{
		ID:     uuid.New(),
		Title:  "Test Book",
		Author: "Test Author",
		Year:   2024,
	}

	// Setup repo expectations (ignore actual DB op)
	mockRepo.On("Store", mock.AnythingOfType("*sql.Tx"), &book).Return(nil)
	mockEmail.On("SendBookCreatedEmail", book).Return(nil)

	// Step 3: Call usecase
	bookUC := usecase.NewBookUseCase(mockRepo, db, logger, mockEmail)
	result, err := bookUC.Create(book)

	// Step 4: Assertions
	assert.NoError(t, err)
	assert.Equal(t, book.Title, result.Title)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	mockRepo.AssertExpectations(t)
	mockEmail.AssertExpectations(t)
}

func TestDeleteBook(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mockRepo := new(repoMocks.BookRepository)
	mockEmail := new(mailMocks.SendGridClient)
	logger := zerolog.Nop()

	bookUC := usecase.NewBookUseCase(mockRepo, db, logger, mockEmail)

	id := uuid.New()
	mockRepo.On("Remove", id).Return(nil)

	err = bookUC.Delete(id)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

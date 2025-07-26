package usecase_test

import (
	"beta-book-api/internal/delivery/request"
	"beta-book-api/internal/entity"
	"beta-book-api/internal/repository/mocks"
	"beta-book-api/internal/usecase"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllBooks(t *testing.T) {
	mockRepo := new(mocks.BookRepository)
	bookUC := usecase.NewBookUseCase(mockRepo)

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
	mockRepo := new(mocks.BookRepository)
	bookUC := usecase.NewBookUseCase(mockRepo)

	id := uuid.New()
	expectedBook := &entity.Book{ID: id, Title: "Clean Code", Author: "Robert C. Martin", Year: 2008}

	mockRepo.On("FetchByID", id).Return(expectedBook, nil)

	result, err := bookUC.GetByID(id)

	assert.NoError(t, err)
	assert.Equal(t, expectedBook, result)
	mockRepo.AssertExpectations(t)
}

func TestCreateBook(t *testing.T) {
	mockRepo := new(mocks.BookRepository)
	bookUC := usecase.NewBookUseCase(mockRepo)

	newBook := entity.Book{Title: "The Pragmatic Programmer", Author: "Andy Hunt", Year: 1999}
	mockRepo.On("Store", mock.AnythingOfType("*entity.Book")).Return(nil)

	result, err := bookUC.Create(newBook)

	assert.NoError(t, err)
	assert.Equal(t, newBook.Title, result.Title)
	assert.Equal(t, newBook.Author, result.Author)
	assert.Equal(t, newBook.Year, result.Year)
	mockRepo.AssertExpectations(t)
}

func TestDeleteBook(t *testing.T) {
	mockRepo := new(mocks.BookRepository)
	bookUC := usecase.NewBookUseCase(mockRepo)

	id := uuid.New()
	mockRepo.On("Remove", id).Return(nil)

	err := bookUC.Delete(id)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

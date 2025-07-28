package mocks

import (
	"beta-book-api/internal/entity"
	"github.com/stretchr/testify/mock"
)

type SendGridClient struct {
	mock.Mock
}

func (s *SendGridClient) SendBookCreatedEmail(book entity.Book) error {
	args := s.Called(book)
	return args.Error(0)
}

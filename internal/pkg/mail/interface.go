package mail

import "beta-book-api/internal/entity"

type EmailClient interface {
	SendBookCreatedEmail(book entity.Book) error
}

package mail

import (
	"beta-book-api/config"
	"beta-book-api/internal/entity"
	"errors"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailClient interface {
	SendBookCreatedEmail(book entity.Book) error
}

type SendGridClient struct {
	apiKey      string
	senderName  string
	senderEmail string
	logger      zerolog.Logger
}

func NewSendGridClient(cfg *config.AppConfig, logger zerolog.Logger) *SendGridClient {
	return &SendGridClient{
		apiKey:      cfg.SendGridAPIKey,
		senderName:  "Beta Book API",
		senderEmail: cfg.SendGridSenderEmail,
		logger:      logger,
	}
}

func (s *SendGridClient) InitSendGrid() *SendGridClient {
	return s
}

func (s *SendGridClient) SendBookCreatedEmail(book entity.Book) error {
	from := mail.NewEmail(s.senderName, s.senderEmail)
	subject := fmt.Sprintf("New Book Created: %s", book.Title)
	to := mail.NewEmail("Recipient", "arief.dfaltah@gmail.com") // You can make this dynamic
	plainTextContent := fmt.Sprintf("A new book has been created:\n\nTitle: %s\nAuthor: %s\nYear: %d", book.Title, book.Author, book.Year)
	htmlContent := fmt.Sprintf("<h1>New Book Created</h1><p><strong>Title:</strong> %s<br><strong>Author:</strong> %s<br><strong>Year:</strong> %d</p>", book.Title, book.Author, book.Year)

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(s.apiKey)

	resp, err := client.Send(message)
	if err != nil {
		s.logger.Error().Err(err).Msgf("❌ Failed to send email")
		return err
	}

	if resp.StatusCode >= 400 {
		s.logger.Error().Err(err).Msgf("❌ Sendgrid error: %s", resp.Body)
		return errors.New("Sendgrid response code not 2**")
	}

	s.logger.Info().Msgf("✅ Email sent with status: %d", resp.StatusCode)
	return nil
}

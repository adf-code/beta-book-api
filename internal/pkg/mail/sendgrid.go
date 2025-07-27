package email

import (
	"beta-book-api/config"
	"github.com/rs/zerolog"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendGridClient struct {
	apiKey string
	from   string
	logger zerolog.Logger
}

func NewSendGridClient(cfg *config.AppConfig, logger zerolog.Logger) SendGridClient {
	return SendGridClient{
		apiKey: cfg.SendGridAPIKey,
		from:   cfg.SendGridSenderEmail,
		logger: logger,
	}
}

func (s *SendGridClient) SendEmail(to, subject, plainText, htmlContent string) error {
	from := mail.NewEmail("Beta Book", s.from)
	toEmail := mail.NewEmail("User", to)
	message := mail.NewSingleEmail(from, subject, toEmail, plainText, htmlContent)

	client := sendgrid.NewSendClient(s.apiKey)
	response, err := client.Send(message)
	if err != nil {
		s.logger.Fatal().Err(err).Msgf("❌ Failed to send emai")
		return err
	}

	s.logger.Info().Msgf("✅ Email sent with status: %d", response.StatusCode)
	return nil
}

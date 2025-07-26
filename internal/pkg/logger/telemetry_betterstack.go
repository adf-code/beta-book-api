package logger

import (
	"bytes"
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog"
)

type TelemetryWriter struct {
	client   *http.Client
	apiKey   string
	endpoint string
}

func NewTelemetryWriter(apiKey string, endpoint string) *TelemetryWriter {
	return &TelemetryWriter{
		client:   &http.Client{Timeout: 5 * time.Second},
		apiKey:   apiKey,
		endpoint: endpoint,
	}
}

func (w *TelemetryWriter) Write(p []byte) (n int, err error) {
	req, err := http.NewRequest("POST", w.endpoint, bytes.NewBuffer(p))
	if err != nil {
		return 0, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+w.apiKey)

	resp, err := w.client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	return len(p), nil
}

func InitLoggerWithTelemetry() zerolog.Logger {
	apiKey := os.Getenv("TELEMETRY_API_KEY")
	if apiKey == "" {
		panic("TELEMETRY_API_KEY not set")
	}

	endpoint := os.Getenv("TELEMETRY_ENDPOINT")
	if endpoint == "" {
		panic("TELEMETRY_ENDPOINT not set")
	}

	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	logtailWriter := NewTelemetryWriter(apiKey, endpoint)
	multiWriter := zerolog.MultiLevelWriter(consoleWriter, logtailWriter)

	return zerolog.New(multiWriter).With().Timestamp().Logger()
}

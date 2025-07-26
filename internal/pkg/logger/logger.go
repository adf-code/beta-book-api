package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var Log zerolog.Logger

func InitLogger(env string) {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if env == "production" {
		Log = zerolog.New(os.Stdout).With().Timestamp().Logger()
	} else {
		Log = zerolog.New(output).With().Timestamp().Logger()
	}
}

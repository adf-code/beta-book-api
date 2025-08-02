package health

import (
	"github.com/rs/zerolog"
)

type HealthHandler struct {
	Logger zerolog.Logger
}

func NewHealthHandler(logger zerolog.Logger) *HealthHandler {
	return &HealthHandler{Logger: logger}
}

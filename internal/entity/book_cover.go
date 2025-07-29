package entity

import (
	"github.com/google/uuid"
	"time"
)

type BookCover struct {
	ID        uuid.UUID  `json:"id"`
	BookID    uuid.UUID  `json:"book_id"`
	FileName  string     `json:"file_name"`
	FileURL   string     `json:"file_url"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

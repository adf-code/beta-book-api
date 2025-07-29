package entity

import (
	"github.com/google/uuid"
	"time"
)

type Book struct {
	ID        uuid.UUID   `json:"id"`
	Title     string      `json:"title"`
	Author    string      `json:"author"`
	Year      int         `json:"year"`
	BookCover []BookCover `json:"cover"`
	CreatedAt *time.Time  `json:"created_at"`
	UpdatedAt *time.Time  `json:"updated_at"`
}

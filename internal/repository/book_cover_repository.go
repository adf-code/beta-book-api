package repository

import (
	"beta-book-api/internal/entity"
	"context"
	"database/sql"
	"github.com/google/uuid"
)

type BookCoverRepository interface {
	Store(ctx context.Context, tx *sql.Tx, cover *entity.BookCover) error
	FetchByBookID(ctx context.Context, bookID uuid.UUID) ([]entity.BookCover, error)
}

type bookCoverRepo struct {
	DB *sql.DB
}

func NewBookCoverRepo(db *sql.DB) BookCoverRepository {
	return &bookCoverRepo{DB: db}
}

func (r *bookCoverRepo) Store(ctx context.Context, tx *sql.Tx, cover *entity.BookCover) error {
	return tx.QueryRowContext(
		ctx,
		"INSERT INTO book_covers (book_id, file_name, file_url) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at",
		cover.BookID, cover.FileName, cover.FileURL,
	).Scan(&cover.ID, &cover.CreatedAt, &cover.UpdatedAt)
}

func (r *bookCoverRepo) FetchByBookID(ctx context.Context, bookID uuid.UUID) ([]entity.BookCover, error) {
	query := `SELECT id, book_id, file_name, file_url, created_at, updated_at
	          FROM book_covers
	          WHERE book_id = $1
	          ORDER BY created_at DESC`

	rows, err := r.DB.QueryContext(ctx, query, bookID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var covers []entity.BookCover
	for rows.Next() {
		var c entity.BookCover
		if err := rows.Scan(&c.ID, &c.BookID, &c.FileName, &c.FileURL, &c.CreatedAt, &c.UpdatedAt); err != nil {
			return nil, err
		}
		covers = append(covers, c)
	}
	return covers, nil
}

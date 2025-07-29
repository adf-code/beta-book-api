package repository

import (
	"beta-book-api/internal/delivery/request"
	"beta-book-api/internal/entity"
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type bookRepo struct {
	DB *sql.DB
}

type BookRepository interface {
	FetchWithQueryParams(ctx context.Context, params request.BookListQueryParams) ([]entity.Book, error)
	FetchByID(ctx context.Context, id uuid.UUID) (*entity.Book, error)
	Store(ctx context.Context, tx *sql.Tx, book *entity.Book) error
	Remove(ctx context.Context, id uuid.UUID) error
}

func NewBookRepo(db *sql.DB) BookRepository {
	return &bookRepo{DB: db}
}

func (r *bookRepo) FetchWithQueryParams(ctx context.Context, params request.BookListQueryParams) ([]entity.Book, error) {
	query := "SELECT id, title, author, year, created_at, updated_at FROM books WHERE 1=1"
	args := []interface{}{}
	argIndex := 1

	// Search
	if params.SearchField != "" && params.SearchValue != "" {
		query += fmt.Sprintf(" AND %s ILIKE $%d", params.SearchField, argIndex)
		args = append(args, "%"+params.SearchValue+"%")
		argIndex++
	}

	// Filters
	for _, f := range params.Filter {
		if len(f.Value) > 0 {
			query += fmt.Sprintf(" AND %s = ANY($%d)", f.Field, argIndex)
			args = append(args, pq.Array(f.Value))
			argIndex++
		}
	}

	// Range
	for _, r := range params.Range {
		if r.From != nil {
			query += fmt.Sprintf(" AND %s >= $%d", r.Field, argIndex)
			args = append(args, *r.From)
			argIndex++
		}
		if r.To != nil {
			query += fmt.Sprintf(" AND %s <= $%d", r.Field, argIndex)
			args = append(args, *r.To)
			argIndex++
		}
	}

	// Sort
	if params.SortField != "" && (params.SortDir == "ASC" || params.SortDir == "DESC") {
		query += fmt.Sprintf(" ORDER BY %s %s", params.SortField, params.SortDir)
	}

	// Pagination
	if params.Page > 0 && params.PerPage > 0 {
		offset := (params.Page - 1) * params.PerPage
		query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argIndex, argIndex+1)
		args = append(args, params.PerPage, offset)
	}
	rows, err := r.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []entity.Book
	for rows.Next() {
		var b entity.Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.Year, &b.CreatedAt, &b.UpdatedAt); err != nil {
			return nil, err
		}
		books = append(books, b)
	}

	return books, nil
}

func (r *bookRepo) FetchByID(ctx context.Context, id uuid.UUID) (*entity.Book, error) {
	var b entity.Book
	err := r.DB.QueryRowContext(ctx, "SELECT id, title, author, year, created_at, updated_at FROM books WHERE id = $1 AND deleted_at is null", id).
		Scan(&b.ID, &b.Title, &b.Author, &b.Year, &b.CreatedAt, &b.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func (r *bookRepo) Store(ctx context.Context, tx *sql.Tx, book *entity.Book) error {
	return tx.QueryRowContext(
		ctx,
		"INSERT INTO books (title, author, year) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at",
		book.Title, book.Author, book.Year,
	).Scan(&book.ID, &book.CreatedAt, &book.UpdatedAt)
}

func (r *bookRepo) Remove(ctx context.Context, id uuid.UUID) error {
	_, err := r.DB.ExecContext(ctx, "DELETE FROM books WHERE id = $1", id)
	return err
}

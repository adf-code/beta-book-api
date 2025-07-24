package repository

import (
	"beta-book-api/internal/entity"
	"database/sql"
	"github.com/google/uuid"
	"log"
)

type bookRepo struct {
	DB *sql.DB
}

func NewBookRepo(db *sql.DB) BookRepository {
	return &bookRepo{DB: db}
}

func (r *bookRepo) FetchAll() ([]entity.Book, error) {
	rows, err := r.DB.Query("SELECT id, title, author, year FROM books")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var books []entity.Book
	for rows.Next() {
		var b entity.Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.Year); err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}

func (r *bookRepo) FetchByID(id uuid.UUID) (*entity.Book, error) {
	var b entity.Book
	err := r.DB.QueryRow("SELECT id, title, author, year FROM books WHERE id = $1", id).
		Scan(&b.ID, &b.Title, &b.Author, &b.Year)
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func (r *bookRepo) Store(book *entity.Book) error {
	return r.DB.QueryRow(
		"INSERT INTO books (title, author, year) VALUES ($1, $2, $3) RETURNING id",
		book.Title, book.Author, book.Year,
	).Scan(&book.ID)
}

func (r *bookRepo) Remove(id uuid.UUID) error {
	_, err := r.DB.Exec("DELETE FROM books WHERE id = $1", id)
	return err
}

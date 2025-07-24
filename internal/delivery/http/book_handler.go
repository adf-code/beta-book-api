package http

import (
	"beta-book-api/internal/entity"
	"beta-book-api/internal/repository"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"strings"
)

type BookHandler struct {
	Repo repository.BookRepository
}

func NewBookHandler(repo repository.BookRepository) *BookHandler {
	return &BookHandler{Repo: repo}
}

func (h *BookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch {
	case r.Method == "GET" && r.URL.Path == "/books":
		h.getAll(w, r)
	case r.Method == "GET" && strings.HasPrefix(r.URL.Path, "/books/"):
		h.getByID(w, r)
	case r.Method == "POST" && r.URL.Path == "/books":
		h.create(w, r)
	case r.Method == "DELETE" && strings.HasPrefix(r.URL.Path, "/books/"):
		h.delete(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (h *BookHandler) getAll(w http.ResponseWriter, r *http.Request) {
	books, err := h.Repo.FetchAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(books)
}

func (h *BookHandler) getByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	if idStr == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid UUID format", http.StatusBadRequest)
		return
	}
	book, err := h.Repo.FetchByID(id)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(book)
}

func (h *BookHandler) create(w http.ResponseWriter, r *http.Request) {
	var book entity.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := h.Repo.Store(&book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(book)
}

func (h *BookHandler) delete(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	if idStr == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid UUID format", http.StatusBadRequest)
		return
	}
	if err := h.Repo.Remove(id); err != nil {
		http.Error(w, "Failed to delete", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

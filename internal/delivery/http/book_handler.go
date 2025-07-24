package http

import (
	"beta-book-api/internal/delivery/response"
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
		response.Failed(w, 500, "books", "getAllBooks", "Error Get All Books")
		return
	}
	response.Success(w, 200, "books", "getAllBooks", "Success Get All Books", books)
}

func (h *BookHandler) getByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	if idStr == "" {
		response.Failed(w, 422, "books", "getBookByID", "Missing ID Parameter, Get Book by ID")
		return
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Failed(w, 422, "books", "getBookByID", "Invalid UUID, Get Book by ID")
		return
	}
	book, err := h.Repo.FetchByID(id)
	if err != nil {
		response.Failed(w, 500, "books", "getBookByID", "Error Get Book by ID")
		return
	}
	response.Success(w, 200, "books", "getBookByID", "Success Get Book by ID", book)
}

func (h *BookHandler) create(w http.ResponseWriter, r *http.Request) {
	var book entity.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		response.Failed(w, 422, "books", "createBook", "Invalid Data, Create Book")
		return
	}

	if err := h.Repo.Store(&book); err != nil {
		response.Failed(w, 500, "books", "createBook", "Error Create Book")
		return
	}
	response.Success(w, 201, "books", "createBook", "Success Create Book", book)
}

func (h *BookHandler) delete(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	if idStr == "" {
		response.Failed(w, 422, "books", "deleteBookByID", "Missing ID Parameter, Delete Book by ID")
		return
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Failed(w, 422, "books", "deleteBookByID", "Invalid UUID, Delete Book by ID")
		return
	}
	if err := h.Repo.Remove(id); err != nil {
		response.Failed(w, 500, "books", "deleteBookByID", "Error Delete Book")
		return
	}
	response.Success(w, 202, "books", "deleteBookByID", "Success Delete Book", nil)
}

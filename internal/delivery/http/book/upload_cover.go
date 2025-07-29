package book

import (
	"beta-book-api/internal/delivery/response"
	"github.com/google/uuid"
	"net/http"
)

func (h *BookHandler) UploadCover(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, "Invalid form", http.StatusBadRequest)
		return
	}

	bookIDStr := r.FormValue("book_id")
	bookID, err := uuid.Parse(bookIDStr)
	if err != nil {
		http.Error(w, "Invalid book_id", http.StatusBadRequest)
		return
	}

	file, fileHeader, err := r.FormFile("cover")
	if err != nil {
		http.Error(w, "Missing file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	cover, err := h.BookCoverUC.Upload(r.Context(), bookID, file, fileHeader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Success(w, 201, "book_covers", "uploadBookCover", "Success Upload Book Cover", cover)
}

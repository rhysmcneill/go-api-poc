package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rhysmcneill/go-api-poc/internal/models"
	"github.com/rhysmcneill/go-api-poc/internal/services"
)

// Handler to get all books
func GetBooksHandler(s *services.BookService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		books := s.GetAllBooks() // fetch all books

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(books)
	}
}

// Handler to create a new book
func CreateBookHandler(s *services.BookService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var b models.Books
		if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		added, err := s.AddBook(b)
		if err != nil {
			http.Error(w, "Failed to add book", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(added)
	}
}

// Handler to get a book by ID
func GetBookByIDHandler(s *services.BookService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		getId := chi.URLParam(r, "id")
		cnvId, err := strconv.Atoi(getId)

		if err != nil {
			http.Error(w, "Failed to retrieve book", http.StatusNotFound)
			fmt.Println("Error converting ID:", err)
			return
		}

		// Call the service to get the book by ID
		book, err := s.GetBookById(cnvId)

		if err != nil {
			http.Error(w, "Failed to retrieve book", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(book)
	}
}

// Handler to update a book by ID
func UpdateBookHandler(s *services.BookService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		getId := chi.URLParam(r, "id")
		cnvId, err := strconv.Atoi(getId)

		if err != nil {
			http.Error(w, "Failed to retrieve book", http.StatusNotFound)
			fmt.Println("Error converting ID:", err)
			return
		}

		var updatedBook models.Books
		if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		// Call the service to update the book by ID
		book, err := s.UpdateBook(cnvId, updatedBook)

		if err != nil {
			http.Error(w, "Failed to update book", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(book)
	}
}

// Handler to delete a book by ID
func DeleteBookHandler(s *services.BookService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		getId := chi.URLParam(r, "id")
		cnvId, err := strconv.Atoi(getId)

		if err != nil {
			http.Error(w, "Failed to retrieve book", http.StatusNotFound)
			fmt.Println("Error converting ID:", err)
			return
		}

		// Call the service to delete the book by ID
		book, err := s.DeleteBook(cnvId)

		if err != nil {
			http.Error(w, "Failed to delete book", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(book)
	}
}

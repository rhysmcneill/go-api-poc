package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/rhysmcneill/go-api-poc/internal/models"
	"github.com/rhysmcneill/go-api-poc/internal/services"
)

func TestGetBooksHandler(t *testing.T) {
	svc := services.BookDefinition()
	h := GetBooksHandler(svc)

	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	w := httptest.NewRecorder()

	h(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, res.StatusCode)
	}

	if ct := res.Header.Get("Content-Type"); ct != "application/json" {
		t.Errorf("expected Content-Type application/json, got %q", ct)
	}

	var books []models.Books
	if err := json.NewDecoder(res.Body).Decode(&books); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if len(books) != 3 {
		t.Errorf("expected 3 books, got %d", len(books))
	}
}

func TestGetBookByIdHandler(t *testing.T) {
	svc := services.BookDefinition()

	r := chi.NewRouter()
	r.Get("/books/{id}", GetBookByIDHandler(svc))

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/books/2", nil)
	r.ServeHTTP(rr, req)

	res := rr.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, res.StatusCode)
	}

	if ct := res.Header.Get("Content-Type"); ct != "application/json" {
		t.Errorf("expected Content-Type application/json, got %q", ct)
	}

	var book models.Books
	if err := json.NewDecoder(res.Body).Decode(&book); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if book.Id != 2 {
		t.Errorf("expected book ID to be 2, got %d", book.Id)
	}
}

func TestAddBookHandler(t *testing.T) {
	svc := services.BookDefinition()

	r := chi.NewRouter()
	r.Post("/books", CreateBookHandler(svc))

	newBook := models.Books{Id: 4, Title: "New Book", Author: "New Author"}
	body, _ := json.Marshal(newBook)

	// httptest.NewRequest expects an io.Reader for the body
	req := httptest.NewRequest("POST", "/books", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	res := rr.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, res.StatusCode)
	}

	if ct := res.Header.Get("Content-Type"); ct != "application/json" {
		t.Errorf("expected Content-Type application/json, got %q", ct)
	}

	var addedBook models.Books
	if err := json.NewDecoder(res.Body).Decode(&addedBook); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if addedBook != newBook {
		t.Errorf("expected added book to be %v, got %v", newBook, addedBook)
	}
}

func TestHealthHandler(t *testing.T) {
	h := HealthHandler()

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	h(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, res.StatusCode)
	}
	body := make([]byte, 2)
	res.Body.Read(body)
	if string(body) != "OK" {
		t.Errorf("expected body to be 'OK', got %q", string(body))
	}
}

func TestUpdateBookHandler(t *testing.T) {
	svc := services.BookDefinition()

	r := chi.NewRouter()
	r.Put("/books/{id}", UpdateBookHandler(svc))
	updatedBook := models.Books{Id: 2, Title: "Updated Book", Author: "Updated Author"}
	body, _ := json.Marshal(updatedBook)

	// httptest.NewRequest expects an io.Reader for the body
	req := httptest.NewRequest("PUT", "/books/2", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	res := rr.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, res.StatusCode)
	}

	if ct := res.Header.Get("Content-Type"); ct != "application/json" {
		t.Errorf("expected Content-Type application/json, got %q", ct)
	}

	var updateBook models.Books
	if err := json.NewDecoder(res.Body).Decode(&updateBook); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

}

func TestDeleteBookHandler(t *testing.T) {
	svc := services.BookDefinition()

	r := chi.NewRouter()
	r.Delete("/books/{id}", DeleteBookHandler(svc))

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/books/2", nil)
	r.ServeHTTP(rr, req)

	res := rr.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, res.StatusCode)
	}

	if ct := res.Header.Get("Content-Type"); ct != "application/json" {
		t.Errorf("expected Content-Type application/json, got %q", ct)
	}

	var deletedBook models.Books
	if err := json.NewDecoder(res.Body).Decode(&deletedBook); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if deletedBook.Id != 2 {
		t.Errorf("expected deleted book ID to be 2, got %d", deletedBook.Id)
	}
}

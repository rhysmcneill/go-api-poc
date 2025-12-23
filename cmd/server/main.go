package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rhysmcneill/go-api-poc/internal/handlers"
	"github.com/rhysmcneill/go-api-poc/internal/services"
)

var bookService = services.BookDefinition()

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/health", handlers.HealthHandler())
	r.Get("/books", handlers.GetBooksHandler(bookService))
	r.Get("/books/{id}", handlers.GetBookByIDHandler(bookService))
	r.Post("/books", handlers.CreateBookHandler(bookService))
	r.Put("/books/{id}", handlers.UpdateBookHandler(bookService))
	r.Delete("/books/{id}", handlers.DeleteBookHandler(bookService))
	http.ListenAndServe(":3000", r)
}

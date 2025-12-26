package services

import (
	"errors"

	"github.com/rhysmcneill/go-api-poc/internal/models"
)

// BookService struct to manage book data
type BookService struct {
	books []models.Books // References Book struct, and stores in memory
}

// BookDefinition initializes the BookService with some sample data
func BookDefinition() *BookService {
	return &BookService{ // defines a slice of book data
		books: []models.Books{
			{Id: 1, Title: "1988", Author: "Rhys McNeill"},
			{Id: 2, Title: "The Thursday Murder Club", Author: "Richard Osman"},
			{Id: 3, Title: "Micahel Cassidies Dump", Author: "Bill Gates"},
		},
	}
}

// GetAllBooks returns all books in the service
func (s *BookService) GetAllBooks() []models.Books {
	return s.books
}

// AddBook adds a new book to the service
func (s *BookService) AddBook(newBook models.Books) (models.Books, error) {
	s.books = append(s.books, newBook)
	return newBook, nil
}

// GetBookById returns a book by its ID
func (s *BookService) GetBookById(id int) (models.Books, error) {
	for _, book := range s.books {
		if book.Id == id {
			return book, nil
		}
	}
	return models.Books{}, errors.New("Could not find book")
}

// UpdateBook updates a book by its ID
func (s *BookService) UpdateBook(id int, updatedBook models.Books) (models.Books, error) {
	for i, book := range s.books {
		if book.Id == id {
			s.books[i] = updatedBook
			return updatedBook, nil
		}
	}
	return models.Books{}, errors.New("Could not find book")
}

// DeleteBook deletes a book by its ID
func (s *BookService) DeleteBook(id int) (models.Books, error) {
	for i, book := range s.books {
		if book.Id == id {
			s.books = append(s.books[:i], s.books[i+1:]...)
			return book, nil
		}
	}
	return models.Books{}, errors.New("Could not find book")
}

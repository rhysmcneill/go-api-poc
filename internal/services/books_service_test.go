package services

import (
	"testing"

	"github.com/rhysmcneill/go-api-poc/internal/models"
)

// TestGetAllBooks tests the GetAllBooks method of BookService
func TestGetAllBooks(t *testing.T) {
	service := BookDefinition()
	books := service.GetAllBooks()

	if len(books) != 3 {
		t.Errorf("Expected 3 books, got %d", len(books))
	}

}

// TestAddBook tests the AddBook method of BookService
func TestAddBook(t *testing.T) {
	service := BookDefinition()
	newBook := models.Books{Id: 4, Title: "New Book", Author: "New Author"}
	addedBook, err := service.AddBook(newBook)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if addedBook != newBook {
		t.Errorf("Expected added book to be %v, got %v", newBook, addedBook)
	}

	if len(service.GetAllBooks()) != 4 {
		t.Errorf("Expected 4 books after addition, got %d", len(service.GetAllBooks()))
	}

}

func TestGetBookById(t *testing.T) {
	service := BookDefinition()
	book, err := service.GetBookById(2)

	if err != nil {
		t.Fail()
	}

	if book.Id != 2 {
		t.Errorf("Expected book ID to be 2, got %d", book.Id)
	}
}

func TestUpdateBook(t *testing.T) {
	service := BookDefinition()
	updatedBook := models.Books{Id: 2, Title: "I Updated the Title", Author: "I also Updated the Author"}
	book, err := service.UpdateBook(2, updatedBook)

	if err != nil {
		t.Fail()
	}

	if book.Title != "I Updated the Title" {
		t.Errorf("Expected book Title to be 'I Updated the Title', got %s", book.Title)
	}

	if book.Author != "I also Updated the Author" {
		t.Errorf("Expected book Author to be 'I also Updated the Author', got %s", book.Author)
	}
}

func TestDeleteBook(t *testing.T) {
	service := BookDefinition()
	deletedBook, err := service.DeleteBook(3)

	if err != nil {
		t.Fail()
	}

	if deletedBook.Id != 3 {
		t.Errorf("Expected deleted book ID to be 3, got %d", deletedBook.Id)
	}

	if len(service.GetAllBooks()) != 2 {
		t.Errorf("Expected 2 books after deletion, got %d", len(service.GetAllBooks()))
	}
}

package services

import (
	"github.com/somnidev/go-fiber/model"
)

type BookService struct {
	books map[string]model.Book
}

func NewBookService() (*BookService, error) {
	bs := map[string]model.Book{
		"001": {Title: "Learning Go: An Idiomatic Approach to Real-World Go Programming", Author: "Jon Bodner"},
		"002": {Title: "Introduction to Algorithms, fourth edition 4th", Author: "Thomas H. Cormen"},
		"003": {Title: "Clean Code: A Handbook of Agile Software Craftsmanship", Author: "Robert C. Martin"},
	}
	return &BookService{books: bs}, nil
}

func ListBooks(bookService *BookService) []model.Book {
	books := make([]model.Book, 0, len(bookService.books))
	for _, value := range bookService.books {
		books = append(books, value)
	}
	return books
}

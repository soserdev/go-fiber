package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/somnidev/go-fiber/controllers"
	"github.com/somnidev/go-fiber/services"
)

var (
	bookService *services.BookService
)

func main() {
	app := fiber.New()
	bookService, _ = services.NewBookService()
	bookController, _ := controllers.NewBookController(bookService)

	app.Get("/books/:id", bookController.GetBookById).Name("books.id")
	app.Get("/books", bookController.GetBooks)
	app.Post("/books", bookController.CreateBook)
	app.Delete("/books/:id", bookController.DeleteBookById)
	app.Put("/books/:id", bookController.UpdateBookById)
	app.Listen(":3000")
}

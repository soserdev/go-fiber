package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/somnidev/go-fiber/model"
	"github.com/somnidev/go-fiber/services"
)

var (
	bookService *services.BookService
)

func GetBooks(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(services.ListBooks(bookService))
}

func CreateBooks(c *fiber.Ctx) error {
	b := new(model.Book)
	if err := c.BodyParser(b); err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(b)
}

func main() {
	app := fiber.New()
	bookService, _ = services.NewBookService()

	app.Get("/books", GetBooks)
	app.Post("/books", CreateBooks)
	app.Listen(":3000")
}

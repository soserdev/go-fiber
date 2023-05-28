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
	return c.Status(fiber.StatusOK).JSON(bookService.ListBooks())
}

func CreateBook(c *fiber.Ctx) error {
	b := new(model.Book)
	if err := c.BodyParser(b); err != nil {
		return err
	}
	nb := bookService.CreateBook(*b)
	return c.Status(fiber.StatusCreated).JSON(nb)
}

func main() {
	app := fiber.New()
	bookService, _ = services.NewBookService()

	app.Get("/books", GetBooks)
	app.Post("/books", CreateBook)
	app.Listen(":3000")
}

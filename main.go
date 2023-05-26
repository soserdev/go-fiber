package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/somnidev/go-fiber/model"
)

func GetBook(c *fiber.Ctx) error {
	book := model.Book{
		Title:  "Learning Go: An Idiomatic Approach to Real-World Go Programming",
		Author: "Jon Bodner",
	}
	return c.Status(fiber.StatusOK).JSON(book)
}

func CreateBook(c *fiber.Ctx) error {
	b := new(model.Book)
	if err := c.BodyParser(b); err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(b)
}

func main() {
	app := fiber.New()

	app.Get("/books", GetBook)
	app.Post("/books", CreateBook)

	app.Listen(":3000")
}

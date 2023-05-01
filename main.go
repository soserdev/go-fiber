package main

import (
	"github.com/gofiber/fiber/v2"
)

func GetBook(c *fiber.Ctx) error {
	book := Book{
		Title:  "Learning Go: An Idiomatic Approach to Real-World Go Programming",
		Author: "Jon Bodner",
	}
	return c.Status(fiber.StatusOK).JSON(book)
}

func main() {
	app := fiber.New()

	app.Get("/", GetBook)

	app.Listen(":3000")
}

package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/soserdev/go-fiber/model"
	"github.com/soserdev/go-fiber/services"
)

type BookController struct {
	bookService *services.BookService
}

func NewBookController(bs *services.BookService) (*BookController, error) {
	return &BookController{bookService: bs}, nil
}

func (bc *BookController) GetBookById(c *fiber.Ctx) error {
	id := c.Params("id")
	b, found := bc.bookService.GetBookById(id)
	if !found {
		c.Status(fiber.StatusNotFound)
		return nil
	}
	return c.Status(fiber.StatusOK).JSON(b)
}

func (bc *BookController) GetBooks(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(bc.bookService.ListBooks())
}

func (bc *BookController) CreateBook(c *fiber.Ctx) error {
	b := new(model.Book)
	if err := c.BodyParser(b); err != nil {
		return err
	}
	nb := bc.bookService.CreateBook(*b)

	location, _ := c.GetRouteURL("books.id", fiber.Map{"id": nb.ID})
	c.Location(location)
	c.Status(fiber.StatusCreated)
	return nil
}

func (bc *BookController) DeleteBookById(c *fiber.Ctx) error {
	id := c.Params("id")
	bc.bookService.DeleteBookById(id)
	c.Status(fiber.StatusNoContent)
	return nil
}

func (bc *BookController) UpdateBookById(c *fiber.Ctx) error {
	b := new(model.Book)
	if err := c.BodyParser(b); err != nil {
		return err
	}
	id := c.Params("id")
	bc.bookService.UpdateBookById(id, *b)
	c.Status(fiber.StatusNoContent)
	return nil
}

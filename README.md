# Go Fiber

[Fiber](https://docs.gofiber.io) is an Express inspired web framework built on top of Fasthttp, the fastest HTTP engine for Go. Designed to ease things up for fast development with zero memory allocation and performance in mind.

## Installation​

First of all, download and install Go. 1.17 or higher is required.

Installation is done using the go get command:

```bash
go get github.com/gofiber/fiber/v2
```

## Introduction

Create file `main.go`  and a hello-world.

```go
package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Listen(":3000")
}
```

```bash
go run .
```

Browse to `http://localhost:3000` and you should see `Hello, World!` on the page.

## Get method returns JSON

Create a `book.go`.

```go
package main

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}
```

Update the `main.go`

```go
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
```

Now we can get the result.

```bash
$ curl localhost:3000
{"title":"Learning Go: An Idiomatic Approach to Real-World Go Programming","author":"Jon Bodner"}
```

### Refactoring to packages

Rename package from `go-fiber` to `github.com/somnidev/go-fiber`.

```bash
go mod edit -module github.com/somnidev/go-fiber
```

Move `book.go` to new subdirectory `model` and rename package to `model`.

```go
package model

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}
```

Now import the new package `github.com/somnidev/go-fiber/model` and use the **package name** `model` to access the `Book` - use `model.Book`.

```go
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
func main() {
	app := fiber.New()
	app.Get("/", GetBook)
	app.Listen(":3000")
}
```

### Create a new Book

We need a [`BodyParser`](https://docs.gofiber.io/api/ctx/#bodyparser) to get the JSON in the body.

```go
func CreateBook(c *fiber.Ctx) error {
	b := new(model.Book)
	if err := c.BodyParser(b); err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(b)
}
```

Now we can test the new endpoint.

```bash
curl -X POST -H "Content-Type: application/json" --data "{\"title\":\"book-title\",\"author\":\"john doe\"}" localhost:3000/books
```

### Create a `BookService`

Now we create a `BookService`.

```go
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
```

We also have to init our `BookServie` in `main.go`.

```go
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
...

func main() {
	app := fiber.New()
	bookService, _ = services.NewBookService()

	app.Get("/books", GetBooks)
	app.Post("/books", CreateBooks)
	app.Listen(":3000")
}
```

Now let's check if it works.

```bash
$ curl -s localhost:3000/books | jq
[
  {
    "title": "Learning Go: An Idiomatic Approach to Real-World Go Programming",
    "author": "Jon Bodner"
  },
  {
    "title": "Introduction to Algorithms, fourth edition 4th",
    "author": "Thomas H. Cormen"
  },
  {
    "title": "Clean Code: A Handbook of Agile Software Craftsmanship",
    "author": "Robert C. Martin"
  }
]
```

### Add a `UUID`

Let's add a `UUID`.

### Refactor to Dependency Injection

In order to understand how you properly do Dependency Injection in Go visit [Dependency Injection Explained](https://markphelps.me/posts/dependency-injection-explained/).

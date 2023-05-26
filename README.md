# Go Fiber

[Fiber](https://docs.gofiber.io) is an Express inspired web framework built on top of Fasthttp, the fastest HTTP engine for Go. Designed to ease things up for fast development with zero memory allocation and performance in mind.

## Installation​

First of all, download and install Go. 1.17 or higher is required.

Installation is done using the go get command:

```bash
go get github.com/gofiber/fiber/v2
```

## Intorduction

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


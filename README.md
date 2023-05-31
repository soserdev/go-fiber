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

### Rest endoint to Create a new Book

Now we add a new method called `CreateBook` to our `main.go` file. We need a [`BodyParser`](https://docs.gofiber.io/api/ctx/#bodyparser) to get the JSON in the body.

```go
func CreateBook(c *fiber.Ctx) error {
	b := new(model.Book)
	if err := c.BodyParser(b); err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(b)
}
```

We also add a new route in our `main.go` file.

```go
app.Post("/books", CreateBook)
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

We also have to init our `BookServie` in `main.go` and change our `GetBooks` method to return a list of books.

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

Let's add an `ID` to our `model.Book`, where we store the `UUID`.

```go
package model

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
```

Now update the `BookService`.

```go
package services

import (
	"github.com/google/uuid"
	"github.com/somnidev/go-fiber/model"
)

type BookService struct {
	books map[string]model.Book
}

func NewBookService() (*BookService, error) {
	uuid1 := uuid.New().String()
	uuid2 := uuid.New().String()
	uuid3 := uuid.New().String()
	bs := map[string]model.Book{
		uuid1: {ID: uuid1, Title: "Learning Go: An Idiomatic Approach to Real-World Go Programming", Author: "Jon Bodner"},
		uuid2: {ID: uuid2, Title: "Introduction to Algorithms, fourth edition 4th", Author: "Thomas H. Cormen"},
		uuid3: {ID: uuid3, Title: "Clean Code: A Handbook of Agile Software Craftsmanship", Author: "Robert C. Martin"},
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

Add a new method called `CreateBook`.

```go
func CreateBook(bookService *BookService, book model.Book) model.Book {
	uuid := uuid.New().String()
	book.ID = uuid
	bookService.books[uuid] = book
	return book
}
```

Finally we update our `CreateBook` method in our `main.go` to use the service.

```go
func CreateBook(c *fiber.Ctx) error {
	b := new(model.Book)
	if err := c.BodyParser(b); err != nil {
		return err
	}
	nb := services.CreateBook(bookService, *b)
	return c.Status(fiber.StatusCreated).JSON(nb)
}
```

That's it. Now we can add a book which will be store in our service.

```bash
$ curl -X POST -H "Content-Type: application/json" --data "{\"title\":\"The miracle of John Doe\",\"author\":\"John Doe\"}" localhost:3000/books
{"id":"8f00940d-3361-4043-a356-27d9b3001c46","title":"The miracle of John Doe","author":"John Doe"
```

And check the result.

```bash
curl -s localhost:3000/books | jq
[
  {
    "id": "8f00940d-3361-4043-a356-27d9b3001c46",
    "title": "The miracle of John Doe",
    "author": "John Doe"
  },
  {
    "id": "819d58e1-99dd-4fbd-ad58-6febcd9fee15",
    "title": "Learning Go: An Idiomatic Approach to Real-World Go Programming",
    "author": "Jon Bodner"
  },
  {
    "id": "f93ff1d7-b8dc-44f4-b716-f4548aa8b046",
    "title": "Introduction to Algorithms, fourth edition 4th",
    "author": "Thomas H. Cormen"
  },
  {
    "id": "a2e34b28-0a81-4a99-8155-a0a854c0efa0",
    "title": "Clean Code: A Handbook of Agile Software Craftsmanship",
    "author": "Robert C. Martin"
  }
]
```

### Refactor to use method receivers

First we refactor the `BookService`.

```go
--- a/services/bookService.go
+++ b/services/bookService.go

-func ListBooks(bookService *BookService) []model.Book {
+func (bookService *BookService) ListBooks() []model.Book {
 	books := make([]model.Book, 0, len(bookService.books))
	for _, value := range bookService.books {
		books = append(books, value)
	}
	return books
}
-func CreateBook(bookService *BookService, book model.Book) model.Book {
+func (bookService *BookService) CreateBook(book model.Book) model.Book {
        uuid := uuid.New().String()
        book.ID = uuid
        bookService.books[uuid] = book
	return book
}
```

Then we have to refactor `main.go`.

```go
--- a/main.go
+++ b/main.go
 func GetBooks(c *fiber.Ctx) error {
-       return c.Status(fiber.StatusOK).JSON(services.ListBooks(bookService))
+       return c.Status(fiber.StatusOK).JSON(bookService.ListBooks())
 }

 func CreateBook(c *fiber.Ctx) error {
        if err := c.BodyParser(b); err != nil {
                return err
        }
-       nb := services.CreateBook(bookService, *b)
+       nb := bookService.CreateBook(*b)
        return c.Status(fiber.StatusCreated).JSON(nb)
 }
```

### Get a book using its `UUID`

First we add a new method the `BookService`that gets the book.

```go
func (bookService *BookService) GetBookById(id string) (model.Book, bool) {
	b, found := bookService.books[id]
	return b, found
}
```

Now we add a new handler to `main.go`.

```go

func GetBookById(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Println(id)
	b, found := bookService.GetBookById(id)
	if !found {
		c.Status(fiber.StatusNotFound)
		return nil
	}
	return c.Status(fiber.StatusOK).JSON(b)
}
```

The [Params Method](https://docs.gofiber.io/api/ctx/#params) `c.Params("id")` can be used to get the route parameters, you could pass an optional default value that will be returned if the param key does not exist. The signature of the Params Method looks like this.

```go
func (c *Ctx) Params(key string, defaultValue ...string) string
```

And finally we need to add a new route.

```go
app.Get("/books/:id", GetBookById)
```

If we try to get a book that does not exist, we get a `404`.

```bash
curl -v localhost:3000/books/XYZ
*   Trying 127.0.0.1:3000...
* Connected to localhost (127.0.0.1) port 3000 (#0)
> GET /books/XYZ HTTP/1.1
> Host: localhost:3000
> User-Agent: curl/7.86.0
> Accept: */*
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 404 Not Found
< Date: Mon, 29 May 2023 10:31:18 GMT
< Content-Length: 0
<
* Connection #0 to host localhost left intact
```

### Add a `Location` header to the Create Request

In order to follow best practices, the `CreateBook` method should return a `Location` response header. Therefore we update the route in `main.go` to create the location url.

```go
app.Get("/books/:id", GetBookById).Name("books.id")
```

Now we can update the `CreateBook` method. The `GetRouteURL` uses the name we defined above.

```go
func CreateBook(c *fiber.Ctx) error {
	b := new(model.Book)
	if err := c.BodyParser(b); err != nil {
		return err
	}
	nb := bookService.CreateBook(*b)

	location, _ := c.GetRouteURL("books.id", fiber.Map{"id": nb.ID})
	c.Location(location)
	c.Status(fiber.StatusCreated)
	return nil
}
```

Now let's check the Create Request.

```bash
curl -v -X POST -H "Content-Type: application/json" --data "{\"title\":\"The miracle of John Doe\",\"author\":\"John Doe\"}" localhost:3000/books
Note: Unnecessary use of -X or --request, POST is already inferred.
*   Trying 127.0.0.1:3000...
* Connected to localhost (127.0.0.1) port 3000 (#0)
> POST /books HTTP/1.1
> Host: localhost:3000
> User-Agent: curl/7.86.0
> Accept: */*
> Content-Type: application/json
> Content-Length: 55
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 201 Created
< Date: Tue, 30 May 2023 17:28:27 GMT
< Content-Length: 0
< Location: /books/e6364fbd-1724-4c4e-b807-f5d4be222293
<
* Connection #0 to host localhost left intact
```


### Delete a book

In order to delete a book we need to add a new method called `` to the `BookService`.

```go
func (bookService *BookService) DeleteBookById(id string) {
	delete(bookService.books, id)
}
```

Now we can add a new controller method to our `main.go`. We return a `204` even if the book is not found.

```go
func DeleteBookById(c *fiber.Ctx) error {
	id := c.Params("id")
	bookService.DeleteBookById(id)
	c.Status(fiber.StatusNoContent)
	return nil
}
```

We add a new route to our `main.go`.

```go
app.Delete("/books/:id", DeleteBookById)
```


### Update a book

### Refactor to Dependency Injection

In order to understand how you properly do Dependency Injection in Go visit [Dependency Injection Explained](https://markphelps.me/posts/dependency-injection-explained/).

First we add a new package `controllers` for our controllers and a file `controllers.go`.

### Use a Concurrent Map

Golang Maos are not thread-safe, see [
Concurrent Map Writing and Reading in Go, or how to deal with the data races](https://webdevstation.com/posts/concurrent-map-writing-and-reading-in-go/). There may occur some racing conditions we should use a [concurrent map](https://pkg.go.dev/github.com/orcaman/concurrent-map#section-readme).



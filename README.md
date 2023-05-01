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
go run main.go
```

Browse to `http://localhost:3000` and you should see `Hello, World!` on the page.


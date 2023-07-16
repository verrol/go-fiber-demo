package main

import (
	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Use(logMiddleware)
	app.Get("/", hello)

	app.Use(addHeaderMiddleware)
	app.Get("/api", anApi)


	log.Info("starting server", "port", 8080)
	err := app.Listen(":8080")
	if err != nil {
		log.Info("failed to start server", "error", err)
	}
}

// logMiddleware is a simple logging middleware
func logMiddleware(c *fiber.Ctx) error {
	log.Info("my middleware called", "path", c.Path(),  "method", c.Method())
	return c.Next()
}

// addHeaderMiddleware adds a header to the response
func addHeaderMiddleware(c *fiber.Ctx) error {
	c.Set("X-Header", "Custom header")
	return c.Next()
}

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func anApi(c *fiber.Ctx) error {
	return c.SendString("this endpoint called after middleware")
}

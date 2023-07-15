package main

import (
	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", hello)

	log.Info("starting server", "port", 8080)
	err := app.Listen(":8080")
	if err != nil {
		log.Info("failed to start server", "error", err)
	}
}

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello World!")
}

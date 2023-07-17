package handlers

import "github.com/gofiber/fiber/v2"

type HelloHandler struct{}

func NewHelloWorldHandler() *HelloHandler {
	return &HelloHandler{}
}

func (h *HelloHandler) Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

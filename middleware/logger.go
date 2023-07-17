package middleware

import (
	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
)

// LoggingMiddleware is a simple logging middleware
type LoggingMiddleWare struct{}

func NewLoggingMiddleWare() *LoggingMiddleWare {
	return &LoggingMiddleWare{}
}

func (m *LoggingMiddleWare) Next(c *fiber.Ctx) error {
	log.Info("my middleware called", "path", c.Path(), "method", c.Method())
	return c.Next()
}

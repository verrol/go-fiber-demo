package middleware

import "github.com/gofiber/fiber/v2"

// AddHeaderMiddleware adds a header to the response
type AddHeaderMiddleware struct{}

func NewAddheaderMiddleWare() *AddHeaderMiddleware {
	return &AddHeaderMiddleware{}
}

func (m *AddHeaderMiddleware) Next(c *fiber.Ctx) error {
	c.Set("X-Header", "Custom header")
	return c.Next()
}

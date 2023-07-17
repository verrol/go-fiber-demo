package handlers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type DemoApi struct {
	bar string
}

func NewDemoApi() *DemoApi {
	return &DemoApi{
		bar: "bar",
	}
}

func (d *DemoApi) GetBar(c *fiber.Ctx) error {
	msg := fmt.Sprintf("Bar is '%v'", d.bar)
	return c.SendString(msg)
}

func (d *DemoApi) PostBar(c *fiber.Ctx) error {
	d.bar = string(c.Body())
	return c.SendStatus(http.StatusOK)
}

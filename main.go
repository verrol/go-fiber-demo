package main

import (
	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/verrol/go-fiber-demo/internal/handlers"
	"github.com/verrol/go-fiber-demo/internal/middleware"
)

func main() {
	app := fiber.New()

	// register middleware
	app.Use("/", middleware.NewLoggingMiddleWare().Next)      // log all requests
	app.Use("/api", middleware.NewAddheaderMiddleWare().Next) // add a custom header for /api/* requests

	// register routes
	app.Get("/hello", handlers.NewHelloWorldHandler().Hello)
	demoApi := handlers.NewDemoApi()
	app.Get("/api/bar", demoApi.GetBar)
	app.Post("/api/bar", demoApi.PostBar)

	// start server
	log.Info("starting server", "port", 8080)
	err := app.Listen(":8080")
	if err != nil {
		log.Info("failed to start server", "error", err)
	}
}

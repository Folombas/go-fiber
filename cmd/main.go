package main

import (
	"folocode/go-fiber/internal/home"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	home.NewHandler(app)

	app.Listen(":8080")
}

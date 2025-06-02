package main

import (
	"folocode/go-fiber/config"
	"folocode/go-fiber/internal/home"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.Init()
	dbConf := config.NewDatabaseConfig()
	log.Println(dbConf)
	app := fiber.New()

	home.NewHandler(app)

	app.Listen(":8080")
}

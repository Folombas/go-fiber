package main

import (
	"folocode/go-fiber/config"
	"folocode/go-fiber/internal/home"
	"folocode/go-fiber/internal/vacancy"
	"folocode/go-fiber/pkg/database"
	"folocode/go-fiber/pkg/logger"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.Init()
	config.NewDatabaseConfig()
	logConfig := config.NewLogConfig()
	dbConfig := config.NewDatabaseConfig()
	customLogger := logger.NewLogger(logConfig)

	app := fiber.New()
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: customLogger,
	}))
	app.Use(recover.New())
	app.Static("/public", "./public")
	dbpool := database.CreateDbPool(dbConfig, customLogger)
	defer dbpool.Close()

	// Repositories
	vacancyRepo := vacancy.NewVacancyRepository(dbpool, customLogger)

	// Handler
	home.NewHandler(app, customLogger)
	vacancy.NewHandler(app, customLogger, vacancyRepo)

	app.Listen(":8080")
}

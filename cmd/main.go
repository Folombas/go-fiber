package main

import (
	"folocode/go-fiber/config"
	"folocode/go-fiber/internal/home"
	"folocode/go-fiber/internal/sitemap"
	"folocode/go-fiber/internal/vacancy"
	"folocode/go-fiber/pkg/database"
	"folocode/go-fiber/pkg/logger"
	"folocode/go-fiber/pkg/middleware"
	"time"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/postgres/v3"
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
	storage := postgres.New(postgres.Config{
		DB:         dbpool,
		Table:      "sessions",
		Reset:      false,
		GCInterval: 10 * time.Second,
	})
	store := session.New(session.Config{
		Storage: storage,
	})
	app.Use(middleware.AuthMiddleware(store))

	// Repositories
	vacancyRepo := vacancy.NewVacancyRepository(dbpool, customLogger)

	// Handler
	home.NewHandler(app, customLogger, vacancyRepo, store)
	vacancy.NewHandler(app, customLogger, vacancyRepo)
	sitemap.NewHandler(app)

	app.Listen(":8080")
}

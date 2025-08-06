package main

import (
	"go-fiber-api-template/config"
	"go-fiber-api-template/internals/entity"
	"go-fiber-api-template/internals/handler"
	"go-fiber-api-template/internals/repository"
	"go-fiber-api-template/internals/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func main() {

	// Initialize Layers
	pgdb := config.InitPostgresDB()
	pgdb.AutoMigrate(&entity.ForexFactory{})

	// Initialize Fiber App
	app := fiber.New()

	// API Routes
	// repositories
	forexFactoryRepo := repository.NewForexFactoryRepository(pgdb)

	// usecases
	forexFactoryUsecase := usecase.NewForexfactoryUsecase(forexFactoryRepo)

	// handlers
	forexFactoryHandler := handler.NewForexFactoryHandler(forexFactoryUsecase)

	app.Get("/hc", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	//Routes
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/forexfactory", forexFactoryHandler.GetForexFactory)
	v1.Post("/forexfactory", forexFactoryHandler.UpsertForexFactory)

	app.Use(limiter.New(config.LimitConfigDefault))

	app.Listen(":3400")
}

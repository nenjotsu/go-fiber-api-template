package config

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgresDB() *gorm.DB {
	var err error
	err = godotenv.Load()
	if err != nil {
		panic(err)
	}
	dsn := os.Getenv("DATABASE_POSTGRES")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	if err = sqlDB.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Database connected!")
	return db
}

var LimitConfigDefault = limiter.Config{
	Max:        30,
	Expiration: 30 * time.Second,
	KeyGenerator: func(c *fiber.Ctx) string {
		return c.IP()
	},
	LimitReached: func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusTooManyRequests)
	},
	SkipFailedRequests:     false,
	SkipSuccessfulRequests: false,
	LimiterMiddleware:      limiter.FixedWindow{},
}

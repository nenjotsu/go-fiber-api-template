package initializers

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectPgDb() {
	var err error
	err = godotenv.Load()
	if err != nil {
		panic(err)
	}

	dsn := os.Getenv("DATABASE_POSTGRES")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

package database

import (
	"fmt"
	"os"

	"example.com/ormproject/models"
	"github.com/lpernett/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()

	if err != nil {
		panic("Could not load env")
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	database := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s&TrustServerCertificate=true",
		user, password, host, database)

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to database")
	}

	DB = db

	db.AutoMigrate(
		&models.Samochod{},
		&models.PosSilnik{},
		&models.TypSilnika{},
		&models.Dealer{},
		&models.Model{},
	)
}

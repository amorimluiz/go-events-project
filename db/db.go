package db

import (
	"amorimluiz/events/models"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()

	if err != nil {
		panic(fmt.Sprintf("Error loading .env file:\n%v", err))
	}

	dbUsername := os.Getenv("DB_USERNAME")
	dbRootPassword := os.Getenv("DB_ROOT_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUsername, dbRootPassword, dbHost, dbPort, dbName)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("failed to connect to the database:\n%v", err))
	}

	createTables()
}

func createTables() {
	DB.AutoMigrate(
		&models.User{},
		&models.Event{},
		&models.Registration{},
	)
}

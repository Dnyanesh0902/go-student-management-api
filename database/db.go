package database

import (
	"fmt"
	"log"
	"os"

	"student-management-api/models"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf(
		"sqlserver://%s:%s@%s:%s?database=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to Connect database:", err)
	}
	DB = db

	err = DB.AutoMigrate(&models.User{}, &models.Student{})

	if err != nil {
		log.Fatal("Migration Failed:", err)
	}

	fmt.Println("SQL SERVER Connected Successfully.")
}
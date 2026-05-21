package main

import (
	"log"

	"student-management-api/database"
	"student-management-api/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect Database
	database.ConnectDB()

	// Create Router
	router := gin.Default()

	// Setup Routes
	routes.SetupRoutes(router)

	// Run Server
	router.Run(":8080")
}
package main

import (
	"log"
	"price-alert-system/internal/handlers"
	"price-alert-system/pkg/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Connect to the database
	if err := database.Connect(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Setup Gin router
	r := gin.Default()
	r.GET("/alerts", handlers.GetAlerts)
	r.POST("/alerts", handlers.CreateAlert)

	// Start the server
	log.Println("Starting server on :8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

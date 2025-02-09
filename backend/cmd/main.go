package main
import "backend/internal/handlers"
import "backend/pkg/database"

import (
	"log"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	if err := database.Connect(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Setup Gin router
	r := gin.Default()
	r.GET("/alerts", handlers.GetAlerts)
	r.POST("/alerts", handlers.CreateAlert)

	// Start server
	log.Println("Server is running on port 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

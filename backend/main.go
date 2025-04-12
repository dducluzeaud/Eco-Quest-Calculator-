package main

import (
	"eco-quest-calculator/backend/models"
	"eco-quest-calculator/backend/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database connection
	models.ConnectDatabase()

	// Create a new Gin router instance
	router := gin.Default()

	// Set up authentication routes
	routes.AuthRoutes(router)

	// Start the server on port 8080
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

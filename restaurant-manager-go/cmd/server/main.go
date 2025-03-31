package main

import (
	"log"
	"restaurant-manager/internal/db"
	"restaurant-manager/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	database := db.NewDatabase()
	if err := database.Connect(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Declare the router in main
	router := gin.Default()

	// Pass the router to setup routes
	routerHandler := routes.NewHandlerRoutes(router, database)
	routerHandler.SetupRoutes()

	// Start the Gin server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

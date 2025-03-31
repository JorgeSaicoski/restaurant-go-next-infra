package routes

import (
	"github.com/gin-gonic/gin"

	"restaurant-manager/internal/db"
	"restaurant-manager/internal/handlers"
)

// RegisterHealthyRoutes sets up health check-related routes
func RegisterHealthyRoutes(router *gin.Engine, database *db.Database) {
	healthyHandler := handlers.NewHealthyHandler(database)
	router.GET("/healthy", healthyHandler.Healthy)
}

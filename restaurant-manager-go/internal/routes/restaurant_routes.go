package routes

import (
	"github.com/gin-gonic/gin"

	"restaurant-manager/internal/db"
	"restaurant-manager/internal/handlers"
)

// RegisterRestaurantRoutes sets up health check-related routes
func RegisterRestaurantRoutes(router *gin.Engine, database *db.Database) {
	restaurantHandler := handlers.NewRestaurantHandler(database)
	restaurantRoutes := router.Group("/restaurant")
	{
		restaurantRoutes.GET("/:id", restaurantHandler.GetRestaurant)
		restaurantRoutes.PATCH("/:id", restaurantHandler.UpdateRestaurant)
		restaurantRoutes.PUT("/", restaurantHandler.CreateRestaurant)
	}
}

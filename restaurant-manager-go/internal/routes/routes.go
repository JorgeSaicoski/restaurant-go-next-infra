package routes

import (
	"log"
	"restaurant-manager/internal/db"

	"github.com/gin-gonic/gin"
)

type HandlerRoutes struct {
	router   *gin.Engine
	database *db.Database
}

func NewHandlerRoutes(router *gin.Engine, database *db.Database) *HandlerRoutes {
	h := &HandlerRoutes{
		router:   router,
		database: database,
	}
	return h
}

func (h *HandlerRoutes) SetupRoutes() {
	if h.router == nil {
		log.Fatal("Failed to set up routes: router is nil")
	}

	if h.database == nil {
		log.Fatal("Failed to set up routes: database instance is nil")
	}

	// Recover from panic to prevent app crashes
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("Panic occurred while setting up routes: %v", r)
		}
	}()

	// Register routes (assuming this function could fail)
	RegisterHealthyRoutes(h.router, h.database)

	log.Println("Routes successfully registered")
}

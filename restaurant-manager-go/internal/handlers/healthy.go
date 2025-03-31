package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"restaurant-manager/internal/db"
)

type HealthyHandler struct {
	database *db.Database
}

// NewHealthyHandler initializes HealthyHandler with a database instance
func NewHealthyHandler(database *db.Database) *HealthyHandler {
	return &HealthyHandler{database: database}
}

// Healthy checks the database connection and responds with the status
func (h *HealthyHandler) Healthy(c *gin.Context) {
	db := h.database.GetDB()

	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
		"db":     db != nil,
	})
}

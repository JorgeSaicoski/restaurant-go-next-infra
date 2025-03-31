package handlers

import (
	"net/http"
	"restaurant-manager/internal/models"
	"restaurant-manager/internal/repository"

	"github.com/gin-gonic/gin"
)

type ExpectedCreateRequest struct {
	RestaurantID uint    `json:"restaurant_id" binding:"required"`
	Name         string  `json:"name" binding:"required"`
	Description  string  `json:"description" binding:"required"`
	Price        float64 `json:"price" binding:"required"`
}

type MenuItemHandler struct {
	repo       repository.MenuItemRepository
	restaurant models.Restaurant
}

func (h *MenuItemHandler) CreateMenuItem(c *gin.Context) {
	var request ExpectedCreateRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var newMenuItem models.MenuItem

	newMenuItem.RestaurantID = request.RestaurantID
	newMenuItem.Name = request.Name
	newMenuItem.Description = request.Description
	newMenuItem.Price = request.Price

	if err := h.repo.Create(&newMenuItem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Menu Item created successfully", "menu_item": newMenuItem})
}

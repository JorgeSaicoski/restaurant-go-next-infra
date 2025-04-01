package handlers

import (
	"net/http"
	"restaurant-manager/internal/models"
	"restaurant-manager/internal/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ExpectedCreateRequest struct {
	RestaurantID uint    `json:"restaurant_id" binding:"required"`
	Name         string  `json:"name" binding:"required"`
	Description  string  `json:"description" binding:"required"`
	Price        float64 `json:"price" binding:"required"`
}

type UpdateMenuItemRequest struct {
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Price       *float64 `json:"price"`
	Available   *bool    `json:"available"`
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

func (h *MenuItemHandler) UpdateMenuItem(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Menu Item ID"})
		return
	}

	var req UpdateMenuItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := make(map[string]any)

	if req.Name != nil {
		updates["name"] = *req.Name
	}
	if req.Description != nil {
		updates["description"] = *req.Description
	}
	if req.Price != nil {
		updates["price"] = *req.Price
	}
	if req.Available != nil {
		updates["available"] = *req.Available
	}

	if err := h.repo.Update(uint(id), updates); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Menu Item updated successfully"})
}

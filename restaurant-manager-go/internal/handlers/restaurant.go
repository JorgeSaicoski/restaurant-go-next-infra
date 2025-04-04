package handlers

import (
	"net/http"
	"restaurant-manager/internal/db"
	"restaurant-manager/internal/models"
	"restaurant-manager/internal/repository"

	"github.com/gin-gonic/gin"
)

type RestaurantHandler struct {
	repo *repository.RestaurantRepository
}

type CreateRestaurantRequest struct {
	Name         string `json:"name" binding:"required"`
	Address      string `json:"address" binding:"required"`
	LimitUsers   bool   `json:"limit_users"`
	MaxCustomers int    `json:"max_customers" binding:"required"`
}

func NewRestaurantHandler(database *db.Database) *RestaurantHandler {
	return &RestaurantHandler{
		repo: repository.NewRestaurantRepository(database),
	}
}

func (h *RestaurantHandler) CreateRestaurant(c *gin.Context) {
	var request CreateRestaurantRequest

	//Parse JSON into request var
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var newRestaurant models.Restaurant

	newRestaurant.Name = request.Name
	newRestaurant.Address = request.Address
	newRestaurant.LimitUsers = request.LimitUsers
	newRestaurant.MaxCustomers = request.MaxCustomers

	if err := h.repo.Create(&newRestaurant); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Restaurant created successfully", "restaurant": newRestaurant})
}

func (h *RestaurantHandler) GetRestaurant(c *gin.Context) {
	// Assume the ID is passed as a URL parameter, e.g. /restaurants/:id
	var idParam struct {
		ID uint `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&idParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	restaurant, err := h.repo.GetByID(idParam.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"restaurant": restaurant})
}

// UpdateRestaurant is an example of updating a restaurant by its ID with a JSON payload.
func (h *RestaurantHandler) UpdateRestaurant(c *gin.Context) {
	var idParam struct {
		ID uint `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&idParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updates map[string]any
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.Update(idParam.ID, updates); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Restaurant updated successfully"})
}

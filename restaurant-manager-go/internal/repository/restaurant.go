package repository

import (
	"errors"
	"restaurant-manager/internal/db"
	"restaurant-manager/internal/models"
)

type RestaurantRepository struct {
	DB *db.Database
}

func NewRestaurantRepository(database *db.Database) *RestaurantRepository {
	return &RestaurantRepository{DB: database}
}

func (r *RestaurantRepository) GetByID(id uint) (*models.Restaurant, error) {
	var restaurant models.Restaurant
	if err := r.DB.GetDB().First(&restaurant, id).Error; err != nil {
		return nil, errors.New("restaurant not found")
	}
	return &restaurant, nil
}

func (r *RestaurantRepository) Create(restaurant *models.Restaurant) error {
	if err := r.DB.GetDB().Create(restaurant).Error; err != nil {
		return errors.New("failed to create restaurant")
	}
	return nil
}

func (r *RestaurantRepository) Update(id uint, updates map[string]any) error {
	var restaurant models.Restaurant

	if err := r.DB.GetDB().First(&restaurant, id).Error; err != nil {
		return errors.New("restaurant not found")
	}

	if err := r.DB.GetDB().Model(&restaurant).Updates(updates).Error; err != nil {
		return errors.New("failed to update restaurant")
	}
	return nil
}

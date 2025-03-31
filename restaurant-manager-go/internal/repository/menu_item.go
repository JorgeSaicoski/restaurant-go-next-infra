package repository

import (
	"errors"
	"restaurant-manager/internal/db"
	"restaurant-manager/internal/models"
)

type MenuItemRepository struct {
	DB *db.Database
}

func NewMenuItemRepository(database *db.Database) *MenuItemRepository {
	return &MenuItemRepository{DB: database}
}

func (r *MenuItemRepository) GetByID(id uint) (*models.MenuItem, error) {
	var menuItem models.MenuItem
	if err := r.DB.GetDB().First(&menuItem, id).Error; err != nil {
		return nil, errors.New("menuItem not found")
	}
	return &menuItem, nil
}

func (r *MenuItemRepository) Create(menuItem *models.MenuItem) error {
	if err := r.DB.GetDB().Create(menuItem).Error; err != nil {
		return errors.New("failed to create menuItem")
	}
	return nil
}

func (r *MenuItemRepository) Update(id uint, updates map[string]any) error {
	var menuItem models.MenuItem

	if err := r.DB.GetDB().First(&menuItem, id).Error; err != nil {
		return errors.New("menuItem not found")
	}

	if err := r.DB.GetDB().Model(&menuItem).Updates(updates).Error; err != nil {
		return errors.New("failed to update menuItem")
	}
	return nil
}

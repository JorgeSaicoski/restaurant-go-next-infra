package models

import "time"

type Restaurant struct {
	ID           uint   `gorm:"primaryKey"`
	Name         string `gorm:"index"`
	Address      string
	LimitUsers   bool
	MaxCustomers int // Allowed concurrent customers
	CreatedAt    time.Time
	UpdatedAt    time.Time

	// Relations
	Staff     []User     `gorm:"many2many:restaurant_user_restaurants;"`
	Orders    []Order    `gorm:"foreignKey:RestaurantID"`
	MenuItems []MenuItem `gorm:"foreignKey:RestaurantID"`
}

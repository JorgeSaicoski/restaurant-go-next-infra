package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	SSOUserID string `gorm:"uniqueIndex"` // Reference to the user in SSO
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time

	// Relations
	Restaurants []Restaurant `gorm:"many2many:restaurant_user_restaurants;"`
}

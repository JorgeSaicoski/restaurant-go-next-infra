package models

import "time"

type RestaurantUserRestaurant struct {
	ID               uint   `gorm:"primaryKey"`
	RestaurantID     uint   `gorm:"index"`
	RestaurantUserID uint   `gorm:"index"`
	Role             string `gorm:"type:varchar(50)"` // Manager, Chef, Waiter, etc.
	CanCreateOrder   bool   `gorm:"default:false"`
	CreatedAt        time.Time
	UpdatedAt        time.Time

	Restaurant     Restaurant `gorm:"foreignKey:RestaurantID"`
	RestaurantUser User       `gorm:"foreignKey:RestaurantUserID"`
}

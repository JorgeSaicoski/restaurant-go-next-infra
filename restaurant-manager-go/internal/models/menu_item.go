package models

import "time"

type MenuItem struct {
	ID           uint `gorm:"primaryKey"`
	RestaurantID uint `gorm:"index"`
	Name         string
	Description  string
	Price        float64
	Available    bool `gorm:"default:true"`
	CreatedAt    time.Time
	UpdatedAt    time.Time

	Restaurant Restaurant `gorm:"foreignKey:RestaurantID"`
}

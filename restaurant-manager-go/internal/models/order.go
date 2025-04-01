package models

import "time"

type Order struct {
	ID           uint `gorm:"primaryKey"`
	RestaurantID uint `gorm:"index"`
	CreatedByID  uint `gorm:"index"` // Who created the order (e.g., a waiter)
	CustomerName string
	TotalAmount  float64
	Status       string `gorm:"type:varchar(20)"` // Pending, Completed, Canceled, etc.
	CreatedAt    time.Time
	UpdatedAt    time.Time

	// Relations
	Restaurant Restaurant  `gorm:"foreignKey:RestaurantID"`
	CreatedBy  User        `gorm:"foreignKey:CreatedByID"`
	Items      []OrderItem `gorm:"foreignKey:OrderID"`
}

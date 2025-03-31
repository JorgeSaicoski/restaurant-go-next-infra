package models

import "time"

type OrderItem struct {
	ID         uint `gorm:"primaryKey"`
	OrderID    uint `gorm:"index"`
	MenuItemID uint `gorm:"index"`
	Quantity   int
	Price      float64
	CreatedAt  time.Time
	UpdatedAt  time.Time

	Order    Order    `gorm:"foreignKey:OrderID"`
	MenuItem MenuItem `gorm:"foreignKey:MenuItemID"`
}

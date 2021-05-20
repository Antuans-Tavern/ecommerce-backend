package model

import "time"

type ShoppingCartProduct struct {
	ShoppingCartID uint `gorm:"primaryKey;"`
	ProductID      int  `gorm:"primaryKey;"`
	Quantity       uint `gorm:"not null;"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

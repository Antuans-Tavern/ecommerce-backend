package model

import (
	"time"
)

type ShoppingCart struct {
	ID        uint  `gorm:"primaryKey"`
	Status    uint8 `gorm:"default:1;"`
	UserID    uint  `gorm:"foreingKey;not null;"`
	User      *User
	Products  []Product `gorm:"many2many:user_languages;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

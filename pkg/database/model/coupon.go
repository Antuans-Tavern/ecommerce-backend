package model

import (
	"time"

	"gorm.io/gorm"
)

type Coupon struct {
	gorm.Model
	Discount     float64 `gorm:"not null;"`
	Type         uint8   `gorm:"not null;"`
	Status       bool    `gorm:"not null;"`
	UserID       uint    `gorm:"foreignKey;"`
	ExpiracyDate time.Time
}

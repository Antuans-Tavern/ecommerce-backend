package model

import (
	"time"

	"gorm.io/gorm"
)

type Coupon struct {
	gorm.Model
	Discount     float64
	Type         uint8
	Status       bool
	UserID       uint `gorm:"foreignKey;"`
	ExpiracyDate time.Time
}

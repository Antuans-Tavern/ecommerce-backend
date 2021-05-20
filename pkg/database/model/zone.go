package model

import "gorm.io/gorm"

type Zone struct {
	gorm.Model
	Name  string  `gorm:"not null;"`
	Price float64 `gorm:"not null;"`
}

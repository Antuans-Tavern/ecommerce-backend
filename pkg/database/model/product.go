package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `gorm:"size:128"`
	Description string
	Price       float64
	Stock       uint
	CategoryID  uint
	Category    *Category
}

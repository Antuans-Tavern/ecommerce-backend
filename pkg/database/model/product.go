package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `gorm:"size:128;unique;not null;"`
	Description string
	Price       float64 `gorm:"not null;"`
	Stock       uint
	CategoryID  uint `gorm:"foreingKey;"`
	Category    *Category
	Images      []Image `gorm:"polymorphic:Imageable;"`
}

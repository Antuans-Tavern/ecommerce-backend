package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string `gorm:"size:128;unique;not null;"`
	Description string
	Price       float64 `gorm:"not null;"`
	Stock       uint
	CategoryID  uint
	Category    *Category
	Images      []Image `gorm:"polymorphic:Imageable;"`
}

func JoinCategories(id int) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Joins("INNER JOIN (?) as all_categories", db.Raw("? UNION ?",
			db.Model(&Category{}).Where("id = ?", id),
			db.Model(&Category{}).Where("parent_id = ?", id),
		))
	}
}

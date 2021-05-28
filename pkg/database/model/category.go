package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name     string `gorm:"size:128;unique;"`
	Status   bool   `gorm:"default:1;"`
	ParentID *uint
	Parent   *Category
	Children []Category `gorm:"foreignKey:ParentID"`
	Products []Product
}

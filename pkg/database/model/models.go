package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Email    string
	Password string
	Status   bool
	Type     uint8
	Profile  Profile
	gorm.Model
}

type Profile struct {
	UserID    uint
	User      *User
	Name      string
	Lastname  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Category struct {
	gorm.Model
	Name     string
	Status   uint8
	Products []Product
}

type Product struct {
	gorm.Model
	Name        string
	Description string
	Price       float64
	Stock       uint
	CategoryID  uint
	Category    Category
}

type ShoppingCart struct {
	ID        uint `gorm:"primaryKey"`
	Status    uint8
	UserID    uint
	User      *User
	Products  []Product `gorm:"many2many:user_languages;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

package database

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string
	Password string
	Status   bool
	Type     int8
	Profile  Profile `gorm:"constraint:OnDelete:CASCADE"`
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

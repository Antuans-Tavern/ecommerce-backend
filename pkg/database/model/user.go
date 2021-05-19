package model

import "gorm.io/gorm"

const (
	USER_TYPE_CLIENT = iota
	USER_TYPE_ADMIN
)

type User struct {
	Email    string  `gorm:"size:64;unique;not null;"`
	Password string  `gorm:"size:64;not null;"`
	Status   bool    `gorm:"default:1;not null;"`
	Type     uint8   `gorm:"default:0;size:1;not null;"`
	Profile  Profile `gorm:"constraint:OnDelete:CASCADE;"`
	Image    Image   `gorm:"polymorphic:Imageable;"`
	gorm.Model
}

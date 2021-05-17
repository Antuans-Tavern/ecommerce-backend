package model

import "gorm.io/gorm"

type User struct {
	Email    string  `gorm:"size:64;"`
	Password string  `gorm:"size:64;"`
	Status   bool    `gorm:"DEFAULT:1"`
	Type     uint8   `gorm:"DEFAULT:0;size:1"`
	Profile  Profile `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	gorm.Model
}

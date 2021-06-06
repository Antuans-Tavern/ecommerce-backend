package model

import (
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	UserID    uint `gorm:"foreingKey;not null;"`
	User      *User
	Name      string `gorm:"size:64;not null;" json:"name" validate:"required,max=64"`
	Lastname  string `gorm:"size:64;not null;" json:"lastname" validate:"required,max=64"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

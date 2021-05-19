package model

import (
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	UserID    uint `gorm:"foreingKey;not null;"`
	User      *User
	Name      string `gorm:"size:64;not null;"`
	Lastname  string `gorm:"size:64;not null;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

package model

import (
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	UserID    uint
	User      *User
	Name      string `gorm:"size:64"`
	Lastname  string `gorm:"size:64"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

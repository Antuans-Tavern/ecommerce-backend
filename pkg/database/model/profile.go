package model

import (
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	UserID    uint   `gorm:"foreingKey;not null;" json:"userID"`
	User      *User  `json:"user"`
	Name      string `gorm:"size:64;not null;" json:"name"`
	Lastname  string `gorm:"size:64;not null;" json:"lastname"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

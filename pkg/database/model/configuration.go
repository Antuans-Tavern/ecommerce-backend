package model

import "time"

type Configuration struct {
	ID        uint   `gorm:"primarykey"`
	name      string `gorm:"size:64;unique;not null;"`
	value     string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

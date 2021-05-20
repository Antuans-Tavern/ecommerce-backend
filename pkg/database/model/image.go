package model

import "time"

type Image struct {
	ID            uint   `gorm:"primarykey;"`
	Path          string `gorm:"not null;"`
	Disk          uint8  `gorm:"not null;"`
	ImageableType string `gorm:"not null;"`
	ImageableID   uint   `gorm:"not null;"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

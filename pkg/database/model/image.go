package model

import (
	"time"

	"gorm.io/gorm"
)

// Disks
const (
	IMAGE_DISK_LOCAL = iota
	IMAGE_DISK_SEED
)

// Types
const (
	IMAGE_TYPE_PRODUCT = iota
	IMAGE_TYPE_PRODUCT_MAIN
)

type Image struct {
	ID            uint   `gorm:"primarykey;"`
	BasePath      string `gorm:"-"`
	Path          string `gorm:"not null;"`
	Type          uint
	Disk          uint   `gorm:"not null;type:smallint"`
	ImageableType string `gorm:"not null;"`
	ImageableID   uint   `gorm:"not null;"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (i *Image) AfterFind(tx *gorm.DB) (err error) {
	if i.Disk == IMAGE_DISK_SEED {
		i.BasePath = "https://picsum.photos"
	}

	return
}

package model

import "time"

type Delivery struct {
	ID        uint    `gorm:"primaryKey"`
	Zone      string  `gorm:"not null;"`
	Price     float64 `gorm:"not null;"`
	Address   string  `gorm:"not null;"`
	InvoiceID uint    `gorm:"foreignKey;not null;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

package model

import "time"

type InvoiceItem struct {
	ID          uint    `gorm:"primarykey;"`
	Name        string  `gorm:"size:128;not null;"`
	Description string  `gorm:"not null;"`
	Price       float64 `gorm:"not null;"`
	InvoiceID   uint    `gorm:"foreingKey;not null;"`
	Invoice     *Invoice
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

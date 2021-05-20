package model

import "time"

type Payment struct {
	ID        uint    `gorm:"primarykey;"`
	Type      uint8   `gorm:"not null;"`
	Amount    float64 `gorm:"not null;"`
	Reference string
	InvoiceID uint `gorm:"foreingKey;not null;"`
	Invoice   *Invoice
	CreatedAt time.Time
	UpdatedAt time.Time
}

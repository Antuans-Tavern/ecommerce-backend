package model

import "time"

type Invoice struct {
	ID        uint    `gorm:"primarykey;"`
	UserID    uint    `gorm:"foreingKey;not null;"`
	Subtotal  float64 `gorm:"not null;"`
	IVA       float64 `gorm:"not null;"`
	Total     float64 `gorm:"not null;"`
	Items     []InvoiceItem
	Payments  []Payment
	CreatedAt time.Time
	UpdatedAt time.Time
}

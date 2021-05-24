package model

import "time"

type Invoice struct {
	ID        uint    `gorm:"primarykey;"`
	UserID    uint    `gorm:"foreingKey;not null;"`
	Subtotal  float64 `gorm:"not null;"`
	IVA       float64 `gorm:"not null;"`
	Total     float64 `gorm:"not null;"`
	Type      uint8   `gorm:"not null;"`
	Status    uint8   `gorm:"not null;"`
	Items     []InvoiceItem
	Payments  []Payment
	User      *User
	CreatedAt time.Time
	UpdatedAt time.Time
}

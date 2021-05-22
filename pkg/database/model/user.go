package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

const (
	USER_TYPE_CLIENT = iota
	USER_TYPE_ADMIN
)

type User struct {
	Email     string  `gorm:"size:64;unique;not null;"`
	Password  string  `gorm:"size:64;not null;"`
	Status    bool    `gorm:"default:1;not null;"`
	Type      int     `gorm:"default:0;size:1;not null;"`
	Profile   Profile `gorm:"constraint:OnDelete:CASCADE;"`
	Image     Image   `gorm:"polymorphic:Imageable;"`
	Invoices  Invoice
	ID        string `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u *User) CreateAccressToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: 15000,
		Id:        u.ID,
	})

	return token.SignedString([]byte(viper.GetString("app_secret")))
}

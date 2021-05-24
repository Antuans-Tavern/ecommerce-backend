package model

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

const (
	USER_TYPE_CLIENT = iota
	USER_TYPE_ADMIN
)

type User struct {
	gorm.Model
	Email    string  `gorm:"size:64;unique;not null;"`
	Password string  `gorm:"size:64;not null;"`
	Status   bool    `gorm:"default:1;not null;"`
	Type     int     `gorm:"default:0;size:1;not null;"`
	Profile  Profile `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
	Image    Image   `gorm:"polymorphic:Imageable;"`
	Invoices Invoice
	Tokens   []AccessToken
}

func (u *User) CreateAccressToken(db *gorm.DB) (string, error) {
	accessToken := &AccessToken{
		UUID:   uuid.New().String(),
		UserID: u.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Id: fmt.Sprint(u.ID),
	})

	if token, err := token.SignedString([]byte(viper.GetString("app_secret"))); err != nil {
		return "", err
	} else {
		accessToken.Token = token
	}

	if err := db.Create(accessToken).Error; err != nil {
		return "", err
	}

	return accessToken.Token, nil
}

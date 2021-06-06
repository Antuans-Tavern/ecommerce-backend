package model

import (
	"context"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// User type constant
const (
	USER_TYPE_CLIENT = iota
	USER_TYPE_ADMIN
)

// User model
type User struct {
	gorm.Model
	Email    string  `gorm:"size:64;unique;not null;" validate:"required,email,db_unique=users,max=64"`
	Password string  `gorm:"size:64;not null;" validate:"required,min=6,max=64"`
	Status   bool    `gorm:"default:1;not null;"`
	Type     int     `gorm:"default:0;size:1;not null;"`
	Profile  Profile `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
	// Image    *Image  //`gorm:"polymorphic:Imageable;"`
	// Invoices Invoice
	Tokens []AccessToken
}

// Register stores a new user in the database. If `authenticate` is true, an authentication token will be created
func (u *User) Register(ctx context.Context, db *gorm.DB, authenticate bool) (string, error) {
	var accessToken string

	err := db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var err error
		if err = u.Save(db); err != nil {
			return err
		}

		if !authenticate {
			return nil
		}

		if accessToken, err = u.CreateAccressToken(db, ""); err != nil {
			return err
		}

		return nil
	})

	return accessToken, err
}

// Save stores the user in the database, if the user already has a primary key, it will be updated
func (u *User) Save(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Omit(clause.Associations).Save(u).Error; err != nil {
			return err
		}

		u.Profile.UserID = u.ID

		if err := tx.Create(&u.Profile).Error; err != nil {
			return err
		}

		return nil
	})
}

// CreateAccressToken creates and store a new access token for the calling user
func (u *User) CreateAccressToken(db *gorm.DB, userAgent string) (string, error) {
	accessToken := &AccessToken{
		UUID:      uuid.New().String(),
		UserID:    u.ID,
		UserAgent: userAgent,
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

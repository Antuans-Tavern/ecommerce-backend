package model

import (
	"gorm.io/gorm"
)

func SetPivots(db *gorm.DB) {
	db.SetupJoinTable(&ShoppingCart{}, "Products", &ShoppingCartProduct{})
}

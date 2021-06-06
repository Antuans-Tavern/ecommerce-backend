package graph

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB        *gorm.DB
	Validator *validator.Validate
}

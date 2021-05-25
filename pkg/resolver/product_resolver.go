package resolver

import (
	"context"

	"github.com/Antuans-Tavern/ecommerce-backend/pkg/database/model"
	"gorm.io/gorm"
)

func ProductResolver(db *gorm.DB, ctx context.Context, pagination int, page int) ([]*model.Product, error) {
	products := []*model.Product{}
	err := db.WithContext(ctx).Preload("Images").Limit(pagination).Offset((page - 1) * pagination).Find(&products).Error
	return products, err
}

func ProductSearch(db *gorm.DB, ctx context.Context, search string, pagination, page int) ([]*model.Product, error) {
	products := []*model.Product{}

	err := db.WithContext(ctx).
		Preload("Images").
		Limit(pagination).
		Offset((page-1)*pagination).
		Where("to_tsvector(name) @@ plainto_tsquery('?')", search).
		Or("to_tsvector(description) @@ plainto_tsquery('?')", search).
		Find(&products).Error

	return products, err
}

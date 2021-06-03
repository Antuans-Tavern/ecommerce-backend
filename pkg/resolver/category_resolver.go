package resolver

import (
	"context"

	"github.com/Antuans-Tavern/ecommerce-backend/pkg/database/model"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/graph/types"
	"gorm.io/gorm"
)

func CategoriesResolver(db *gorm.DB, ctx context.Context, search *string, pagination, page, productsPagination, productsPage int) (*types.CategoryCollection, error) {
	categories := []*model.Category{}

	tx := db.WithContext(ctx).
		Preload("Products", func(db *gorm.DB) *gorm.DB {
			return db.Limit(productsPagination).Offset((productsPage - 1) * productsPagination)
		}).
		Limit(pagination).Offset((page - 1) * pagination)

	if search != nil {
		tx = tx.Where(
			db.Where("categories.name ILIKE ?", "%"+*search+"%"),
		)
	}

	err := tx.Find(&categories).Error

	var count int64
	db.Table("categories").Count(&count)

	return &types.CategoryCollection{
		Data:  categories,
		Total: int(count),
	}, err
}

func CategoryResolver(db *gorm.DB, ctx context.Context, id int, productsPagination, productsPage int) (*model.Category, error) {
	category := &model.Category{}

	err := db.Preload("Products", func(db *gorm.DB) *gorm.DB {
		return db.Limit(productsPagination).Offset((productsPage - 1) * productsPagination)
	}).First(category, "id = ?", id).Error

	return category, err
}

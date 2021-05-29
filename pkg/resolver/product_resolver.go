package resolver

import (
	"context"

	"github.com/Antuans-Tavern/ecommerce-backend/pkg/database/model"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/graph/types"
	"gorm.io/gorm"
)

func ProductsResolver(db *gorm.DB, ctx context.Context, search *string, pagination int, page int, category *int) (*types.ProductCollection, error) {
	products := []*model.Product{}

	tx := db.WithContext(ctx).
		Preload("Images").
		Preload("Category").
		Limit(pagination).Offset((page - 1) * pagination)

	if search != nil {
		tx = tx.Where(
			db.Where("products.name ILIKE ?", "%"+*search+"%").
				Or("products.description ILIKE ?", "%"+*search+"%"),
		)
	}

	if category != nil {
		tx = tx.Joins("INNER JOIN categories ON categories.id = products.category_id").Where("categories.id = ?", *category)
	}

	err := tx.Find(&products).Error

	var count int64

	db.Table("products").Count(&count)

	return &types.ProductCollection{
		Data:  products,
		Total: int(count),
	}, err
}

func ProductResolver(db *gorm.DB, ctx context.Context, id int) (*model.Product, error) {
	product := &model.Product{}

	err := db.Joins("Category").First(product, "products.id = ?", id).Error

	return product, err
}

package resolver

import (
	"context"

	"github.com/Antuans-Tavern/ecommerce-backend/pkg/database/model"
	"gorm.io/gorm"
)

func ProductResolver(db *gorm.DB, ctx context.Context, search *string, pagination int, page int, category *int) ([]*model.Product, error) {
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
	return products, err
}

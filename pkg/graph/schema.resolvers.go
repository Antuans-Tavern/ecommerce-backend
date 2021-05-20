package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Antuans-Tavern/ecommerce-backend/pkg/database/model"
	generated1 "github.com/Antuans-Tavern/ecommerce-backend/pkg/graph/generated"
)

func (r *queryResolver) Users(ctx context.Context, pagination int, page int) ([]*model.User, error) {
	users := []*model.User{}

	err := r.DB.WithContext(ctx).Joins("Profile").Find(&users).Limit(pagination).Offset(page).Error

	return users, err
}

func (r *queryResolver) User(ctx context.Context, id int) (*model.User, error) {
	user := &model.User{}
	err := r.DB.WithContext(ctx).Joins("Profile").Find(user, "users.id = ?", id).Error

	return user, err
}

func (r *queryResolver) Categories(ctx context.Context, pagination int, page int) ([]*model.Category, error) {
	categories := []*model.Category{}
	err := r.DB.WithContext(ctx).Preload("Products").Find(&categories).Limit(pagination).Offset(page).Error

	return categories, err
}

func (r *queryResolver) Category(ctx context.Context, id int) (*model.Category, error) {
	category := &model.Category{}
	err := r.DB.WithContext(ctx).Preload("Products").Find(category, "categories.id = ?", id).Error

	return category, err
}

func (r *queryResolver) Products(ctx context.Context, pagination int, page int) ([]*model.Product, error) {
	products := []*model.Product{}
	err := r.DB.WithContext(ctx).Joins("Category").Find(&products).Error

	return products, err
}

func (r *queryResolver) Product(ctx context.Context, id int) (*model.Product, error) {
	product := &model.Product{}

	err := r.DB.WithContext(ctx).Joins("Category").Find(product, "products.id = ?", id).Error

	return product, err
}

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

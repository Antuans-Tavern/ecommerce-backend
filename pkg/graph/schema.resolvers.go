package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Antuans-Tavern/ecommerce-backend/pkg/database/model"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/graph/generated"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/graph/types"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/resolver"
)

func (r *mutationResolver) Register(ctx context.Context, data types.Register) (*types.Login, error) {
	return resolver.Register(r.DB, r.Validator, ctx, data)
}

func (r *mutationResolver) Login(ctx context.Context, email string, password string) (*types.Login, error) {
	return resolver.Login(r.DB, r.Validator, ctx, email, password)
}

func (r *queryResolver) Categories(ctx context.Context, search *string, pagination int, page int, productsPagination int, productsPage int) (*types.CategoryCollection, error) {
	return resolver.CategoriesResolver(r.DB, ctx, search, pagination, page, productsPagination, productsPage)
}

func (r *queryResolver) Category(ctx context.Context, id int, productsPagination int, productsPage int) (*model.Category, error) {
	return resolver.CategoryResolver(r.DB, ctx, id, productsPagination, productsPage)
}

func (r *queryResolver) Products(ctx context.Context, search *string, pagination int, page int, category *int) (*types.ProductCollection, error) {
	return resolver.ProductsResolver(r.DB, ctx, search, pagination, page, category)
}

func (r *queryResolver) Product(ctx context.Context, id int) (*model.Product, error) {
	return resolver.ProductResolver(r.DB, ctx, id)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

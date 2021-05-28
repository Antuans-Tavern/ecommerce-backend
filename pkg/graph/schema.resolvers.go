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
	return resolver.Register(r.DB, ctx, data)
}

func (r *queryResolver) Login(ctx context.Context, email string, password string) (*types.Login, error) {
	return resolver.Login(r.DB, ctx, email, password)
}

func (r *queryResolver) Products(ctx context.Context, search *string, pagination int, page int, category *int) ([]*model.Product, error) {
	return resolver.ProductResolver(r.DB, ctx, search, pagination, page, category)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

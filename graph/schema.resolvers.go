package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Antuans-Tavern/ecommerce-backend/graph/generated"
	"github.com/Antuans-Tavern/ecommerce-backend/graph/model"
)

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	users := []*model.User{}

	err := r.DB.Find(&users).Error

	return users, err
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

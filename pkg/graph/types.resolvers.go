package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Antuans-Tavern/ecommerce-backend/pkg/database/model"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/graph/generated"
)

func (r *categoryResolver) ID(ctx context.Context, obj *model.Category) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *productResolver) ID(ctx context.Context, obj *model.Product) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *productResolver) Stock(ctx context.Context, obj *model.Product) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *productResolver) CategoryID(ctx context.Context, obj *model.Product) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *profileResolver) UserID(ctx context.Context, obj *model.Profile) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *shoppingCartResolver) ID(ctx context.Context, obj *model.ShoppingCart) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *shoppingCartResolver) Status(ctx context.Context, obj *model.ShoppingCart) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) ID(ctx context.Context, obj *model.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) Type(ctx context.Context, obj *model.User) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) ShoppingCarts(ctx context.Context, obj *model.User) ([]*model.ShoppingCart, error) {
	panic(fmt.Errorf("not implemented"))
}

// Category returns generated.CategoryResolver implementation.
func (r *Resolver) Category() generated.CategoryResolver { return &categoryResolver{r} }

// Product returns generated.ProductResolver implementation.
func (r *Resolver) Product() generated.ProductResolver { return &productResolver{r} }

// Profile returns generated.ProfileResolver implementation.
func (r *Resolver) Profile() generated.ProfileResolver { return &profileResolver{r} }

// ShoppingCart returns generated.ShoppingCartResolver implementation.
func (r *Resolver) ShoppingCart() generated.ShoppingCartResolver { return &shoppingCartResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type categoryResolver struct{ *Resolver }
type productResolver struct{ *Resolver }
type profileResolver struct{ *Resolver }
type shoppingCartResolver struct{ *Resolver }
type userResolver struct{ *Resolver }

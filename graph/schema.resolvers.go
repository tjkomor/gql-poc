package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"gql-poc/graph/generated"
	"gql-poc/graph/model"
	"gql-poc/graph/database"
)

var db = database.Connect()


func (r *mutationResolver) BakePizza(ctx context.Context, input *model.NewPizza) (*model.Pizza, error) {
	return db.Save(input), nil
}

func (r *queryResolver) Pizza(ctx context.Context, id string) (*model.Pizza, error) {
	return db.FindByID(id), nil
}

func (r *queryResolver) Pizzas(ctx context.Context) ([]*model.Pizza, error) {
	return db.All(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

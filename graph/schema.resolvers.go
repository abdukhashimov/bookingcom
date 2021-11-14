package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"abdukhashimov/mybron.uz/graph/generated"
	"abdukhashimov/mybron.uz/graph/model"
	"abdukhashimov/mybron.uz/storage/sqlc"
	"context"
	"fmt"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context, limit *int, offset *int) ([]*model.User, error) {
	return r.Services.UserService.GetAll(ctx, sqlc.GetUsersParams{
		Offset: int32(*offset),
		Limit:  int32(*limit),
	})
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

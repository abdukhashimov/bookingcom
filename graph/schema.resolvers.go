package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"abdukhashimov/mybron.uz/graph/generated"
	"abdukhashimov/mybron.uz/graph/model"
	"abdukhashimov/mybron.uz/pkg/logger"
	"abdukhashimov/mybron.uz/storage/sqlc"
	"context"
	"fmt"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	r.log.Info("create user request", logger.Any("payload", input))
	res, err := r.services.UserService.Create(ctx, input)
	r.logErrorAndInfo(res, err)
	return res, err
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id *string, input *model.NewUser) (*model.User, error) {
	r.log.Info("update user request", logger.Any("payload", input))
	res, err := r.services.UserService.UpdateUser(ctx, id, input)
	r.logErrorAndInfo(res, err)
	return res, err
}

func (r *queryResolver) Users(ctx context.Context, limit *int, offset *int) ([]*model.User, error) {
	return r.services.UserService.GetAll(ctx, sqlc.GetUsersParams{
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) logErrorAndInfo(res interface{}, err error) {
	if err != nil {
		r.log.Error("create user request failed", logger.Error(err))
	} else {
		r.log.Info("create user request success", logger.Any("response", res))
	}
}
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

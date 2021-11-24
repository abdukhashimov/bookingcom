package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"abdukhashimov/mybron.uz/graph/generated"
	"abdukhashimov/mybron.uz/graph/model"
	"abdukhashimov/mybron.uz/pkg/jwt"
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

func (r *mutationResolver) Login(ctx context.Context, input *model.LoginParams) (*model.LoginResponse, error) {
	r.log.Info("login request", logger.Any("payload", input))
	res, err := r.services.UserService.Login(ctx, input)
	r.logErrorAndInfo(res, err)
	return res, err
}

func (r *queryResolver) Users(ctx context.Context, limit *int, offset *int) ([]*model.User, error) {
	return r.services.UserService.GetAll(ctx, sqlc.GetUsersParams{
		Offset: int32(*offset),
		Limit:  int32(*limit),
	})
}

func (r *queryResolver) UserMe(ctx context.Context) (*model.User, error) {
	userInfo, _ := ctx.Value("key").(jwt.TokenPayload)
	r.log.Info("user me request", logger.Any("user_id", userInfo))
	res, err := r.services.UserService.GetUserByID(ctx)
	r.logErrorAndInfo(res, err)
	return res, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

func (r *mutationResolver) logErrorAndInfo(res interface{}, err error) {
	if err != nil {
		r.log.Error("request failed", logger.Error(err))
	} else {
		r.log.Info("request success", logger.Any("response", res))
	}
}

func (r *queryResolver) logErrorAndInfo(res interface{}, err error) {
	if err != nil {
		r.log.Error("request failed", logger.Error(err))
	} else {
		r.log.Info("request success", logger.Any("response", res))
	}
}
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

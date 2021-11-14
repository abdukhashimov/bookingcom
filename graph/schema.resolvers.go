package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"abdukhashimov/mybron.uz/graph/generated"
	"abdukhashimov/mybron.uz/graph/model"
	"abdukhashimov/mybron.uz/storage/sqlc"
	"context"
)

func (r *queryResolver) Users(ctx context.Context, limit *int, offset *int) ([]*model.GetUser, error) {
	return r.Services.UserService.GetAll(
		ctx,
		sqlc.GetUsersParams{
			Limit:  int32(*limit),
			Offset: int32(*offset),
		},
	)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

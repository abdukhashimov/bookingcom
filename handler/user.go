package handler

import (
	"abdukhashimov/mybron.uz/storage/sqlc"
	"context"
)

type UserService struct {
	Db *sqlc.Queries
}

func NewUserHandler(db *sqlc.Queries) *UserService {
	return &UserService{
		Db: db,
	}
}

func (u *UserService) GetAll(ctx context.Context, req sqlc.GetUsersParams) ([]sqlc.User, error) {
	return u.Db.GetUsers(ctx, req)
}

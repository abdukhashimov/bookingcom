package handler

import (
	"abdukhashimov/mybron.uz/storage/sqlc"
	"context"
)

type UserService struct {
	db *sqlc.Queries
}

func NewUserHandler(db *sqlc.Queries) *UserService {
	return &UserService{
		db: db,
	}
}

func (u *UserService) GetAll(ctx context.Context, req sqlc.GetUsersParams) ([]sqlc.User, error) {
	return u.db.GetUsers(ctx, req)
}

func (u *UserService) Create(ctx context.Context, req sqlc.CreateUserParams) (sqlc.User, error) {
	return u.db.CreateUser(ctx, req)
}

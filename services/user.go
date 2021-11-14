package services

import (
	"abdukhashimov/mybron.uz/graph/model"
	"abdukhashimov/mybron.uz/storage/sqlc"
	"context"
	"encoding/json"
	"time"
)

type UserService struct {
	db *sqlc.Queries
}

func NewUserService(db *sqlc.Queries) *UserService {
	return &UserService{
		db: db,
	}
}

func (u *UserService) GetAll(ctx context.Context, req sqlc.GetUsersParams) ([]*model.User, error) {
	var (
		res []*model.User
	)

	users, err := u.db.GetUsers(ctx, req)
	if err != nil {
		return res, err
	}

	bytes, err := json.Marshal(users)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(bytes, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (u *UserService) Create(ctx context.Context, req sqlc.CreateUserParams) (sqlc.User, error) {
	createdAt := time.Now()
	updatedAt := time.Now()
	req.CreatedAt = &createdAt
	req.UpdatedAt = &updatedAt
	return u.db.CreateUser(ctx, req)
}

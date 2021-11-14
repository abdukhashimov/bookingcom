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

func (u *UserService) Create(ctx context.Context, req *model.NewUser) (*model.User, error) {
	var (
		payload  sqlc.CreateUserParams
		response model.User
	)
	createdAt := time.Now()
	updatedAt := time.Now()
	payload.CreatedAt = &createdAt
	payload.UpdatedAt = &updatedAt
	err := modelToStruct(req, &payload)
	if err != nil {
		return nil, err
	}

	res, err := u.db.CreateUser(ctx, payload)
	if err != nil {
		return nil, err
	}

	err = modelToStruct(res, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

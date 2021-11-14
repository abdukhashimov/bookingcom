package services

import (
	"abdukhashimov/mybron.uz/graph/model"
	"abdukhashimov/mybron.uz/storage/sqlc"
	"context"
	"time"

	"github.com/google/uuid"
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

	err = modelToStruct(users, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (u *UserService) Create(ctx context.Context, req model.NewUser) (*model.User, error) {
	var (
		payload  sqlc.CreateUserParams
		response model.User
	)
	createdAt := time.Now()
	updatedAt := time.Now()
	payload.CreatedAt = &createdAt
	payload.UpdatedAt = &updatedAt
	payload.ID = uuid.NewString()

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

func (u *UserService) UpdateUser(ctx context.Context, id *string, req *model.NewUser) (*model.User, error) {
	var (
		payload  sqlc.UpdateUserParams
		response model.User
	)

	payload.ID = *id
	err := modelToStruct(req, &payload)
	if err != nil {
		return nil, err
	}

	err = u.db.UpdateUser(ctx, payload)
	if err != nil {
		return nil, err
	}

	userDb, err := u.db.GetUser(ctx, payload.ID)
	if err != nil {
		return nil, err
	}

	err = modelToStruct(userDb, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

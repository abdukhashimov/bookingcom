package services

import (
	"abdukhashimov/mybron.uz/graph/model"
	"abdukhashimov/mybron.uz/pkg/jwt"
	"abdukhashimov/mybron.uz/pkg/messages"
	"abdukhashimov/mybron.uz/storage/sqlc"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type bookObjecService struct {
	db  *sqlc.Queries
	jwt jwt.Jwt
}

func NewBookObjectService(db *sqlc.Queries, jwt jwt.Jwt) *bookObjecService {
	return &bookObjecService{
		db:  db,
		jwt: jwt,
	}
}

func (u *bookObjecService) Get(ctx context.Context) (*model.User, error) {
	var (
		res model.User
	)
	userInfo, ok := ctx.Value("auth").(jwt.TokenPayload)
	if !ok {
		return &res, errors.New(messages.ErrorAuthFailed)
	}
	userDb, err := u.db.GetUser(ctx, userInfo.UserID)
	if err != nil {
		return &res, err
	}

	err = modelToStruct(userDb, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

func (u *bookObjecService) GetAll(ctx context.Context, req sqlc.GetUsersParams) ([]*model.User, error) {
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

func (u *bookObjecService) Create(ctx context.Context, req model.NewUser) (*model.User, error) {
	var (
		payload  sqlc.CreateUserParams
		response model.User
	)

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

func (u *bookObjecService) UpdateUser(ctx context.Context, id *string, req *model.NewUser) (*model.User, error) {
	var (
		payload  sqlc.UpdateUserParams
		response model.User
	)

	payload.ID = *id

	err := modelToStruct(req, &payload)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%+v", payload)
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

package services

import (
	"abdukhashimov/mybron.uz/graph/model"
	"abdukhashimov/mybron.uz/pkg/jwt"
	"abdukhashimov/mybron.uz/pkg/messages"
	"abdukhashimov/mybron.uz/storage/sqlc"
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/google/uuid"
)

type UserService struct {
	db  *sqlc.Queries
	jwt jwt.Jwt
}

func NewUserService(db *sqlc.Queries, jwt jwt.Jwt) *UserService {
	return &UserService{
		db:  db,
		jwt: jwt,
	}
}

func (u *UserService) UpdateMe(ctx context.Context, req model.UpdateUser) (*model.UpdateResponse, error) {
	var (
		res     model.UpdateResponse
		payload sqlc.UpdateUserParams
	)

	userInfo, ok := ctx.Value("auth").(jwt.TokenPayload)
	if !ok {
		return &res, errors.New(messages.ErrorAuthFailed)
	}

	err := modelToStruct(req, &payload)
	if err != nil {
		return &res, errors.New(messages.ErrorFailedToParseJSON)
	}

	// userDB, err := u.db.GetUser(ctx, userInfo.UserID)
	// if err != nil {
	// 	return &res, errors.New(messages.ErrorFailedToRetreiveFromDB)
	// }

	// TODO: Need to find the way to update the object without hard codedly writing it!
	v := reflect.ValueOf(payload)
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
	}
	payload.ID = userInfo.UserID
	// err = u.db.UpdateUser(ctx, payload)
	// if err != nil {
	// 	return &res, errors.New(messages.ErrorFailedToUpdateDBObject)
	// }

	res.ID = userInfo.UserID
	return &res, nil
}

func (u *UserService) GetUserByID(ctx context.Context) (*model.User, error) {
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

func (u *UserService) Login(ctx context.Context, req *model.LoginParams) (*model.LoginResponse, error) {
	var (
		res model.LoginResponse
		err error
	)
	userDB, err := u.db.GetUserByPhoneNumber(ctx, req.PhoneNumber)
	if err != nil {
		return &res, errors.New("invalid credentials")
	}

	token, err := u.jwt.GenerateToken(jwt.TokenPayload{
		UserID: userDB.ID,
	})

	if err != nil {
		return &res, err
	}

	res.Token = token
	res.RefreshToken = "no refresh token for now!"

	return &res, err
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

	updatedAt := time.Now()
	payload.ID = *id
	payload.UpdatedAt = &updatedAt

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

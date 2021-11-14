package handler_test

import (
	"abdukhashimov/mybron.uz/handler"
	"abdukhashimov/mybron.uz/storage/sqlc"
	"context"
	"fmt"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

var (
	userHandler *handler.UserService
)

func createUser(t *testing.T) sqlc.User {
	user, err := userHandler.Create(
		context.Background(),
		sqlc.CreateUserParams{
			ID: faker.UUIDHyphenated(),
		},
	)
	if err != nil {
		panic(err)
	}

	return user
}

func TestGetAllUsers(t *testing.T) {
	userHandler = handler.NewUserHandler(queries)
	users, err := userHandler.GetAll(context.Background(), sqlc.GetUsersParams{
		Offset: 0,
		Limit:  10,
	})
	assert.NoError(t, err, "Failed to retreive all users")
	fmt.Println(users)
}

package services_test

import (
	"abdukhashimov/mybron.uz/graph/model"
	"abdukhashimov/mybron.uz/services"
	"abdukhashimov/mybron.uz/storage/sqlc"
	"context"
	"fmt"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

var (
	userHandler *services.UserService
)

func createUser(t *testing.T) *model.User {
	firstName := faker.FirstName()
	lastName := faker.LastName()
	long, lat := faker.Longitude(), faker.Latitude()

	user, err := userHandler.Create(
		context.Background(),
		model.NewUser{
			FirstName: firstName,
			LastName:  lastName,
			Long:      &long,
			Lat:       &lat,
		},
	)

	if err != nil {
		t.Error("Failed to create user")
		panic(err)
	}

	return user
}

func TestGetAllUsers(t *testing.T) {
	userHandler = services.NewUserService(queries)
	createUser(t)
	users, err := userHandler.GetAll(context.Background(), sqlc.GetUsersParams{
		Offset: 0,
		Limit:  10,
	})
	assert.NoError(t, err, "Failed to retreive all users")
	fmt.Println(users)
}

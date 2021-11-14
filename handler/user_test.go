package handler_test

import (
	"abdukhashimov/mybron.uz/handler"
	"abdukhashimov/mybron.uz/storage/sqlc"
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllUsers(t *testing.T) {
	userHandler := handler.NewUserHandler(queries)
	users, err := userHandler.GetAll(context.Background(), sqlc.GetUsersParams{
		Offset: 0,
		Limit:  10,
	})
	assert.NoError(t, err, "Failed to retreive all users")
	fmt.Println(users)
}

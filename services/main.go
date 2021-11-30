package services

import (
	"abdukhashimov/mybron.uz/pkg/jwt"
	"abdukhashimov/mybron.uz/storage/sqlc"
	"encoding/json"
)

type Services struct {
	userService *userService
}

func NewServices(db *sqlc.Queries, jwt jwt.Jwt) *Services {
	return &Services{
		userService: NewUserService(db, jwt),
	}
}

func (s *Services) UserService() *userService {
	return s.userService
}

func modelToStruct(input interface{}, output interface{}) error {
	bytes, err := json.Marshal(input)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, output)
	if err != nil {
		return err
	}

	return nil
}

package services

import (
	"abdukhashimov/mybron.uz/storage/sqlc"
	"encoding/json"
)

type Services struct {
	UserService *UserService
}

func NewServices(db *sqlc.Queries) *Services {
	return &Services{
		UserService: NewUserService(db),
	}
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

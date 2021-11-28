package services

import (
	"abdukhashimov/mybron.uz/pkg/jwt"
	"abdukhashimov/mybron.uz/storage/sqlc"
	"encoding/json"
	"reflect"
)

type Services struct {
	UserService *UserService
}

func NewServices(db *sqlc.Queries, jwt jwt.Jwt) *Services {
	return &Services{
		UserService: NewUserService(db, jwt),
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

func makeObjectPatch(input interface{}, output interface{}) error {
	value := reflect.ValueOf(input)
	types := value.Type()
	reflectValue := reflect.ValueOf(&output).Elem()

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i).Interface()
		if reflect.TypeOf(field).Kind().String() == "ptr" {
			if reflect.ValueOf(field).IsNil() {
				// need to assign it from output
			} else {
				// just go on
			}
		} else {
			if reflect.ValueOf(field).IsZero() {
				// need to assign it from output
			} else {
				// just go on
			}
		}
	}
	return nil
}

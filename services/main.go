package services

import "abdukhashimov/mybron.uz/storage/sqlc"

type Services struct {
	UserService *UserService
}

func NewServices(db *sqlc.Queries) *Services {
	return &Services{
		UserService: NewUserService(db),
	}
}

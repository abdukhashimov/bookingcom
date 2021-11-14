package handler

import "abdukhashimov/mybron.uz/storage/sqlc"

type Handlers struct {
	UserHandler *UserService
}

func NewHandler(db *sqlc.Queries) *Handlers {
	return &Handlers{
		UserHandler: NewUserHandler(db),
	}
}

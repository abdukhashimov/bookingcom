package services

import (
	"abdukhashimov/mybron.uz/pkg/jwt"
	"abdukhashimov/mybron.uz/storage/sqlc"
)

type faqService struct {
	db  *sqlc.Queries
	jwt jwt.Jwt
}

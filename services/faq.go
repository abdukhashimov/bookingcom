package services

import (
	"abdukhashimov/mybron.uz/graph/model"
	"abdukhashimov/mybron.uz/pkg/jwt"
	"abdukhashimov/mybron.uz/storage/sqlc"
	"context"
)

type faqService struct {
	db  *sqlc.Queries
	jwt jwt.Jwt
}

func NewFaqService(db *sqlc.Queries, jwt jwt.Jwt) *faqService {
	return &faqService{
		db:  db,
		jwt: jwt,
	}
}

func (f *faqService) CreateFaq(ctx context.Context, req model.CreateFaq) (*model.Faq, error) {
	return nil, nil
}

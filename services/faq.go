package services

import (
	"abdukhashimov/mybron.uz/graph/model"
	"abdukhashimov/mybron.uz/pkg/jwt"
	"abdukhashimov/mybron.uz/storage/sqlc"
	"context"

	"github.com/google/uuid"
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
	var (
		payload  sqlc.CreateFaqParams
		response model.Faq
	)

	payload.ID = uuid.NewString()
	err := modelToStruct(req, payload)
	if err != nil {
		return &response, err
	}

	res, err := f.db.CreateFaq(ctx, payload)
	if err != nil {
		return &response, err
	}

	err = modelToStruct(res, &response)

	return &response, err
}

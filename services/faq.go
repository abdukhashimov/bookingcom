package services

import (
	"abdukhashimov/mybron.uz/config"
	"abdukhashimov/mybron.uz/graph/model"
	"abdukhashimov/mybron.uz/pkg/jwt"
	"abdukhashimov/mybron.uz/pkg/utils"
	"abdukhashimov/mybron.uz/storage/sqlc"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/teris-io/shortid"
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
		res      sqlc.Faq
		err      error
	)

	shortId, err := shortid.Generate()
	if err != nil {
		return &response, errors.New("internal server error")
	}

	payload.Slug = slug.Make(fmt.Sprintf("%s-%s", utils.FirstN(req.Question), shortId))

	err = modelToStruct(req, &payload)
	if err != nil {
		return &response, err
	}

	for _, lang := range config.Langs {
		payload.ID = uuid.NewString()
		payload.Lang = lang
		res, err = f.db.CreateFaq(ctx, payload)
		if err != nil {
			return &response, err
		}
	}

	err = modelToStruct(res, &response)
	return &response, err
}

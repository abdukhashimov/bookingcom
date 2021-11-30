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

func (f *faqService) GetFAQ(ctx context.Context, slug, lang string) (*model.Faq, error) {
	var (
		response model.Faq
	)

	res, err := f.db.GetFaq(ctx, sqlc.GetFaqParams{
		Slug: slug,
		Lang: lang,
	})

	if err != nil {
		return &response, err
	}

	err = modelToStruct(res, &response)

	return &response, err
}

func (f *faqService) GetAllFAQ(ctx context.Context, limit, offset *int, lang string) (*model.GetAllResp, error) {
	var (
		response model.GetAllResp
		err      error
	)

	res, err := f.db.GetAllFaq(ctx, sqlc.GetAllFaqParams{
		Lang:   lang,
		Limit:  int32(*limit),
		Offset: int32(*offset),
	})

	if err != nil {
		return &response, err
	}

	err = modelToStruct(res, &response.Faqs)

	return &response, err
}

func (f *faqService) UpdateFaq(ctx context.Context, req model.UpdateFaq) (*model.Faq, error) {
	var (
		response model.Faq
		payload  sqlc.UpdateFaqParams
	)

	err := modelToStruct(req, &payload)
	if err != nil {
		return &response, err
	}

	err = f.db.UpdateFaq(ctx, payload)
	if err != nil {
		return &response, err
	}

	faqDb, err := f.db.GetFaq(context.Background(), sqlc.GetFaqParams{
		Slug: req.Slug,
		Lang: req.Lang,
	})

	if err != nil {
		return &response, err
	}

	err = modelToStruct(faqDb, &response)
	return &response, err
}

func (f *faqService) DeleteFaq(ctx context.Context, slug string) (string, error) {
	err := f.db.DeleteFaq(ctx, slug)
	return "", err
}

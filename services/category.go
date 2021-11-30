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

type categoryService struct {
	db  *sqlc.Queries
	jwt jwt.Jwt
}

func NewCategoryService(db *sqlc.Queries, jwt jwt.Jwt) *categoryService {
	return &categoryService{
		db:  db,
		jwt: jwt,
	}
}

func (c *categoryService) CreateCategory(ctx context.Context, req model.CreateCategory) (*model.Category, error) {
	var (
		payload  sqlc.CreateCategoryParams
		response model.Category
		res      sqlc.Category
		err      error
	)

	shortId, err := shortid.Generate()
	if err != nil {
		return &response, errors.New("internal server error")
	}

	payload.Slug = slug.Make(fmt.Sprintf("%s-%s", utils.FirstN(req.Name), shortId))

	err = modelToStruct(req, &payload)
	if err != nil {
		return &response, err
	}

	for _, lang := range config.Langs {
		payload.ID = uuid.NewString()
		payload.Lang = lang
		res, err = c.db.CreateCategory(ctx, payload)
		if err != nil {
			return &response, err
		}
	}

	err = modelToStruct(res, &response)
	return &response, err
}

func (c *categoryService) UpdateCategory(ctx context.Context, req model.UpdateCategory) (*model.Category, error) {
	var (
		response model.Category
		payload  sqlc.UpdateCategoryParams
	)

	err := modelToStruct(req, &payload)
	if err != nil {
		return &response, err
	}

	err = c.db.UpdateCategory(ctx, payload)
	if err != nil {
		return &response, err
	}

	faqDb, err := c.db.GetCategory(context.Background(), sqlc.GetCategoryParams{
		Slug: req.Slug,
		Lang: req.Lang,
	})

	if err != nil {
		return &response, err
	}

	err = modelToStruct(faqDb, &response)
	return &response, err
}

func (c *categoryService) DeleteCategory(ctx context.Context, slug string) (string, error) {
	err := c.db.DeleteCategory(ctx, slug)
	return "", err
}

func (c *categoryService) GetCategory(ctx context.Context, slug, lang string) (*model.Category, error) {
	var (
		response model.Category
	)

	res, err := c.db.GetCategory(ctx, sqlc.GetCategoryParams{
		Slug: slug,
		Lang: lang,
	})

	if err != nil {
		return &response, err
	}

	err = modelToStruct(res, &response)

	return &response, err
}

func (c *categoryService) GetAllCategory(ctx context.Context, lang string, limit, offset *int) (*model.GetAllCategory, error) {
	var (
		response model.GetAllCategory
		err      error
	)

	res, err := c.db.GetAllFaq(ctx, sqlc.GetAllFaqParams{
		Lang:   lang,
		Limit:  int32(*limit),
		Offset: int32(*offset),
	})

	if err != nil {
		return &response, err
	}

	err = modelToStruct(res, &response.Categories)

	return &response, err
}

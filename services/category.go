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
	return nil, nil
}

func (c *categoryService) DeleteCategory(ctx context.Context, slug string) (string, error) {
	return "", nil
}

func (c *categoryService) GetCategory(ctx context.Context, slug, lang string) (*model.Category, error) {
	return nil, nil
}

func (c *categoryService) GetAllCategory(ctx context.Context, lang string, limit, offset *int) (*model.GetAllCategory, error) {
	return nil, nil
}

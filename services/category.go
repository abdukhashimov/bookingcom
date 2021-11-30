package services

import (
	"abdukhashimov/mybron.uz/graph/model"
	"abdukhashimov/mybron.uz/pkg/jwt"
	"abdukhashimov/mybron.uz/storage/sqlc"
	"context"
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
	return nil, nil
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

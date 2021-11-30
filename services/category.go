package services

import (
	"abdukhashimov/mybron.uz/graph/model"
	"abdukhashimov/mybron.uz/pkg/jwt"
	"abdukhashimov/mybron.uz/storage/sqlc"
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

func (c *categoryService) CreateCategory() (*model.Category, error) {
	return nil, nil
}

func (c *categoryService) UpdateCategory() (*model.Category, error) {
	return nil, nil
}

func (c *categoryService) DeleteCategory() (string, error) {
	return "", nil
}

func (c *categoryService) GetCategory() (*model.Category, error) {
	return nil, nil
}

func (c *categoryService) GetAllCategory() (*model.GetAllCategory, error) {
	return nil, nil
}

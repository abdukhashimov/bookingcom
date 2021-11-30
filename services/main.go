package services

import (
	"abdukhashimov/mybron.uz/pkg/jwt"
	"abdukhashimov/mybron.uz/storage/sqlc"
	"encoding/json"
)

type Services struct {
	userService     *userService
	faqService      *faqService
	categoryService *categoryService
}

func NewServices(db *sqlc.Queries, jwt jwt.Jwt) *Services {
	return &Services{
		userService:     NewUserService(db, jwt),
		faqService:      NewFaqService(db, jwt),
		categoryService: NewCategoryService(db, jwt),
	}
}

func (s *Services) UserService() *userService {
	return s.userService
}

func (s *Services) FaqService() *faqService {
	return s.faqService
}

func (s *Services) CategoryService() *categoryService {
	return s.categoryService
}

func modelToStruct(input interface{}, output interface{}) error {
	bytes, err := json.Marshal(input)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, output)
	if err != nil {
		return err
	}

	return nil
}

package services

import (
	"abdukhashimov/mybron.uz/graph/model"
	"abdukhashimov/mybron.uz/pkg/jwt"
	"abdukhashimov/mybron.uz/storage/sqlc"
	"context"

	"github.com/google/uuid"
)

type bookObjecService struct {
	db  *sqlc.Queries
	jwt jwt.Jwt
}

func NewBookObjectService(db *sqlc.Queries, jwt jwt.Jwt) *bookObjecService {
	return &bookObjecService{
		db:  db,
		jwt: jwt,
	}
}

func (u *bookObjecService) Get(ctx context.Context, id string) (*model.BookObject, error) {
	var (
		res model.BookObject
	)

	BookObjectDb, err := u.db.GetBookObject(ctx, id)
	if err != nil {
		return &res, err
	}

	err = modelToStruct(BookObjectDb, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

func (u *bookObjecService) GetAll(ctx context.Context, req sqlc.GetAllBookObjectParams) (*model.GetAllBookObject, error) {
	var (
		res model.GetAllBookObject
	)

	bookObjects, err := u.db.GetAllBookObject(ctx, req)
	if err != nil {
		return &res, err
	}

	err = modelToStruct(bookObjects, &res.Objects)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

func (u *bookObjecService) Create(ctx context.Context, req model.BookObject) (*model.BookObject, error) {
	var (
		payload  sqlc.CreateBookObjectParams
		response model.BookObject
	)

	payload.ID = uuid.NewString()

	err := modelToStruct(req, &payload)
	if err != nil {
		return nil, err
	}

	res, err := u.db.CreateBookObject(ctx, payload)
	if err != nil {
		return nil, err
	}

	err = modelToStruct(res, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (u *bookObjecService) UpdateBookObject(ctx context.Context, id *string, req *model.UpdateBookObject) (*model.BookObject, error) {
	var (
		payload  sqlc.UpdateBookObjectParams
		response model.BookObject
	)

	payload.ID = *id

	err := modelToStruct(req, &payload)
	if err != nil {
		return nil, err
	}

	err = u.db.UpdateBookObject(ctx, payload)
	if err != nil {
		return nil, err
	}

	bookObjectDb, err := u.db.GetBookObject(ctx, payload.ID)
	if err != nil {
		return nil, err
	}

	err = modelToStruct(bookObjectDb, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (b *bookObjecService) Delete(ctx context.Context, id string) (string, error) {
	err := b.db.DeleteBookObject(ctx, id)
	return "", err
}

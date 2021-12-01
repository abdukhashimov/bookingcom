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
		res           model.GetAllBookObject
		defaultStatus int32 = 8
	)

	if req.Status == nil {
		req.Status = &defaultStatus
	}

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

func (u *bookObjecService) Create(ctx context.Context, req model.CreateBookObject) (*model.BookObject, error) {
	var (
		payload  sqlc.CreateBookObjectParams
		response model.BookObject
		status   int32 = 1
	)

	payload.ID = uuid.NewString()
	payload.Status = &status
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

func (u *bookObjecService) UpdateBookObject(ctx context.Context, req model.UpdateBookObject) (*model.BookObject, error) {
	var (
		payload  sqlc.UpdateBookObjectParams
		response model.BookObject
	)

	statusObj, err := u.db.GetStatusByName(ctx, *req.Status)
	if err != nil {
		return &response, err
	}

	err = modelToStruct(req, &payload)
	if err != nil {
		return nil, err
	}

	payload.Status = statusObj.ID
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
	var (
		status string = "deactivated"
	)

	statusObj, err := b.db.GetStatusByName(ctx, status)
	if err != nil {
		return "", err
	}

	err = b.db.UpdateStatus(ctx, sqlc.UpdateStatusParams{
		Status: &statusObj.ID,
		ID:     id,
	})

	return "deleted...", err
}

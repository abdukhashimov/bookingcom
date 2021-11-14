// Code generated by sqlc. DO NOT EDIT.
// source: user.sql

package sqlc

import (
	"context"
	"time"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
        id,
        first_name,
        last_name,
        phone_number,
        is_verified,
        long,
        lat,
        user_type,
        created_at,
        updated_at
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING id, first_name, last_name, phone_number, is_verified, long, lat, user_type, created_at, updated_at
`

type CreateUserParams struct {
	ID          string     `json:"id"`
	FirstName   *string    `json:"first_name"`
	LastName    *string    `json:"last_name"`
	PhoneNumber string     `json:"phone_number"`
	IsVerified  *bool      `json:"is_verified"`
	Long        *float64   `json:"long"`
	Lat         *float64   `json:"lat"`
	UserType    int32      `json:"user_type"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.FirstName,
		arg.LastName,
		arg.PhoneNumber,
		arg.IsVerified,
		arg.Long,
		arg.Lat,
		arg.UserType,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.PhoneNumber,
		&i.IsVerified,
		&i.Long,
		&i.Lat,
		&i.UserType,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, first_name, last_name, phone_number, is_verified, long, lat, user_type, created_at, updated_at
FROM users
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.PhoneNumber,
		&i.IsVerified,
		&i.Long,
		&i.Lat,
		&i.UserType,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT id, first_name, last_name, phone_number, is_verified, long, lat, user_type, created_at, updated_at
FROM users
ORDER BY created_at desc OFFSET $1
LIMIT $2
`

type GetUsersParams struct {
	Offset int32 `json:"offset"`
	Limit  int32 `json:"limit"`
}

func (q *Queries) GetUsers(ctx context.Context, arg GetUsersParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getUsers, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.PhoneNumber,
			&i.IsVerified,
			&i.Long,
			&i.Lat,
			&i.UserType,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users
SET first_name = $1,
    last_name = $2,
    phone_number = $3,
    is_verified = $4,
    long = $5,
    lat = $6,
    user_type = $7,
    updated_at = $8
WHERE id = $9
`

type UpdateUserParams struct {
	FirstName   *string    `json:"first_name"`
	LastName    *string    `json:"last_name"`
	PhoneNumber string     `json:"phone_number"`
	IsVerified  *bool      `json:"is_verified"`
	Long        *float64   `json:"long"`
	Lat         *float64   `json:"lat"`
	UserType    int32      `json:"user_type"`
	UpdatedAt   *time.Time `json:"updated_at"`
	ID          string     `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser,
		arg.FirstName,
		arg.LastName,
		arg.PhoneNumber,
		arg.IsVerified,
		arg.Long,
		arg.Lat,
		arg.UserType,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}

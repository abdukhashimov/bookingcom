// Code generated by sqlc. DO NOT EDIT.
// source: book_object.sql

package sqlc

import (
	"context"
	"time"
)

const createBookObject = `-- name: CreateBookObject :one
INSERT INTO book_object (
        id,
        category,
        title,
        location,
        about,
        status,
        opens_at,
        long,
        lat,
        closes_at,
        created_at,
        updated_at
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        $10,
        $11,
        $12
    )
RETURNING id, category, title, location, long, lat, about, status, opens_at, closes_at, created_at, updated_at
`

type CreateBookObjectParams struct {
	ID        string    `json:"id"`
	Category  string    `json:"category"`
	Title     string    `json:"title"`
	Location  string    `json:"location"`
	About     string    `json:"about"`
	Status    *int32    `json:"status"`
	OpensAt   string    `json:"opens_at"`
	Long      float64   `json:"long"`
	Lat       float64   `json:"lat"`
	ClosesAt  string    `json:"closes_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (q *Queries) CreateBookObject(ctx context.Context, arg CreateBookObjectParams) (BookObject, error) {
	row := q.db.QueryRowContext(ctx, createBookObject,
		arg.ID,
		arg.Category,
		arg.Title,
		arg.Location,
		arg.About,
		arg.Status,
		arg.OpensAt,
		arg.Long,
		arg.Lat,
		arg.ClosesAt,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i BookObject
	err := row.Scan(
		&i.ID,
		&i.Category,
		&i.Title,
		&i.Location,
		&i.Long,
		&i.Lat,
		&i.About,
		&i.Status,
		&i.OpensAt,
		&i.ClosesAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteBookObject = `-- name: DeleteBookObject :exec
DELETE FROM book_object
WHERE id = $1
`

func (q *Queries) DeleteBookObject(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteBookObject, id)
	return err
}

const getAllBookObject = `-- name: GetAllBookObject :many
SELECT id, category, title, location, long, lat, about, status, opens_at, closes_at, created_at, updated_at
FROM book_object
ORDER BY created_at desc OFFSET $1
LIMIT $2
`

type GetAllBookObjectParams struct {
	Offset int32 `json:"offset"`
	Limit  int32 `json:"limit"`
}

func (q *Queries) GetAllBookObject(ctx context.Context, arg GetAllBookObjectParams) ([]BookObject, error) {
	rows, err := q.db.QueryContext(ctx, getAllBookObject, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []BookObject
	for rows.Next() {
		var i BookObject
		if err := rows.Scan(
			&i.ID,
			&i.Category,
			&i.Title,
			&i.Location,
			&i.Long,
			&i.Lat,
			&i.About,
			&i.Status,
			&i.OpensAt,
			&i.ClosesAt,
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

const getBookObject = `-- name: GetBookObject :one
SELECT id, category, title, location, long, lat, about, status, opens_at, closes_at, created_at, updated_at
FROM book_object
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetBookObject(ctx context.Context, id string) (BookObject, error) {
	row := q.db.QueryRowContext(ctx, getBookObject, id)
	var i BookObject
	err := row.Scan(
		&i.ID,
		&i.Category,
		&i.Title,
		&i.Location,
		&i.Long,
		&i.Lat,
		&i.About,
		&i.Status,
		&i.OpensAt,
		&i.ClosesAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateBookObject = `-- name: UpdateBookObject :exec
UPDATE book_object
SET category = COALESCE(NULLIF($1, ''), category),
    title = COALESCE(NULLIF($2, ''), title),
    location = COALESCE(NULLIF($3, ''), location),
    long = COALESCE($4, long),
    lat = COALESCE($5, lat),
    about = COALESCE(NULLIF($6, ''), about),
    status = COALESCE(NULLIF($7, 0), status),
    opens_at = COALESCE(NULLIF($8, ''), opens_at),
    closes_at = COALESCE(NULLIF($9, ''), closes_at),
    updated_at = COALESCE($10, updated_at)
WHERE id = $11
`

type UpdateBookObjectParams struct {
	Category  interface{} `json:"category"`
	Title     interface{} `json:"title"`
	Location  interface{} `json:"location"`
	Long      float64     `json:"long"`
	Lat       float64     `json:"lat"`
	About     interface{} `json:"about"`
	Status    interface{} `json:"status"`
	OpensAt   interface{} `json:"opens_at"`
	ClosesAt  interface{} `json:"closes_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	ID        string      `json:"id"`
}

func (q *Queries) UpdateBookObject(ctx context.Context, arg UpdateBookObjectParams) error {
	_, err := q.db.ExecContext(ctx, updateBookObject,
		arg.Category,
		arg.Title,
		arg.Location,
		arg.Long,
		arg.Lat,
		arg.About,
		arg.Status,
		arg.OpensAt,
		arg.ClosesAt,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}

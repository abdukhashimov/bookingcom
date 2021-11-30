// Code generated by sqlc. DO NOT EDIT.
// source: category.sql

package sqlc

import (
	"context"
	"database/sql"
)

const createCategory = `-- name: CreateCategory :one
INSERT INTO category (
        id,
        parent_id,
        image,
        active,
        slug,
        lang,
        information,
        created_at,
        updated_at
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id, parent_id, name, image, active, slug, lang, information, created_at, updated_at
`

type CreateCategoryParams struct {
	ID          string       `json:"id"`
	ParentID    *string      `json:"parent_id"`
	Image       *string      `json:"image"`
	Active      *bool        `json:"active"`
	Slug        string       `json:"slug"`
	Lang        string       `json:"lang"`
	Information *string      `json:"information"`
	CreatedAt   sql.NullTime `json:"created_at"`
	UpdatedAt   sql.NullTime `json:"updated_at"`
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) (Category, error) {
	row := q.db.QueryRowContext(ctx, createCategory,
		arg.ID,
		arg.ParentID,
		arg.Image,
		arg.Active,
		arg.Slug,
		arg.Lang,
		arg.Information,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.ParentID,
		&i.Name,
		&i.Image,
		&i.Active,
		&i.Slug,
		&i.Lang,
		&i.Information,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteCategory = `-- name: DeleteCategory :exec
DELETE FROM category
WHERE slug = $1
`

func (q *Queries) DeleteCategory(ctx context.Context, slug string) error {
	_, err := q.db.ExecContext(ctx, deleteCategory, slug)
	return err
}

const getAllCategory = `-- name: GetAllCategory :many
SELECT id, parent_id, name, image, active, slug, lang, information, created_at, updated_at
FROM category
WHERE lang = $1
ORDER BY created_at desc
LIMIT $3 OFFSET $2
`

type GetAllCategoryParams struct {
	Lang   string `json:"lang"`
	Offset int32  `json:"offset_"`
	Limit  int32  `json:"limit_"`
}

func (q *Queries) GetAllCategory(ctx context.Context, arg GetAllCategoryParams) ([]Category, error) {
	rows, err := q.db.QueryContext(ctx, getAllCategory, arg.Lang, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Category
	for rows.Next() {
		var i Category
		if err := rows.Scan(
			&i.ID,
			&i.ParentID,
			&i.Name,
			&i.Image,
			&i.Active,
			&i.Slug,
			&i.Lang,
			&i.Information,
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

const getCategory = `-- name: GetCategory :one
SELECT id, parent_id, name, image, active, slug, lang, information, created_at, updated_at
FROM category
WHERE slug = $1 and lang = $2
LIMIT 1
`

type GetCategoryParams struct {
	Slug string `json:"slug"`
	Lang string `json:"lang"`
}

func (q *Queries) GetCategory(ctx context.Context, arg GetCategoryParams) (Category, error) {
	row := q.db.QueryRowContext(ctx, getCategory, arg.Slug, arg.Lang)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.ParentID,
		&i.Name,
		&i.Image,
		&i.Active,
		&i.Slug,
		&i.Lang,
		&i.Information,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateCategory = `-- name: UpdateCategory :exec
UPDATE category
SET parent_id = COALESCE(NULLIF($1, ''), parent_id),
    image = COALESCE(NULLIF($2, ''), image),
    active = COALESCE($3, active),
    information = COALESCE($4, information),
    created_at = COALESCE($5, created_at),
    updated_at = COALESCE($6, updated_at)
WHERE slug = $7 and lang = $8
`

type UpdateCategoryParams struct {
	ParentID    interface{}  `json:"parent_id"`
	Image       interface{}  `json:"image"`
	Active      *bool        `json:"active"`
	Information *string      `json:"information"`
	CreatedAt   sql.NullTime `json:"created_at"`
	UpdatedAt   sql.NullTime `json:"updated_at"`
	Slug        string       `json:"slug"`
	Lang        string       `json:"lang"`
}

func (q *Queries) UpdateCategory(ctx context.Context, arg UpdateCategoryParams) error {
	_, err := q.db.ExecContext(ctx, updateCategory,
		arg.ParentID,
		arg.Image,
		arg.Active,
		arg.Information,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Slug,
		arg.Lang,
	)
	return err
}

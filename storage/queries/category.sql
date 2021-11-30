-- name: GetCategory :one
SELECT *
FROM category
WHERE slug = @slug and lang = @lang
LIMIT 1;

-- name: GetAllCategory :many
SELECT *
FROM category
WHERE lang = @lang
ORDER BY created_at desc
LIMIT @limit_ OFFSET @offset_;

-- name: CreateCategory :one
INSERT INTO category (
        id,
        question,
        answer,
        slug,
        lang,
        active,
        created_at,
        updated_at
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: UpdateCategory :exec
UPDATE category
SET question = COALESCE(NULLIF(@question, ''), question),
    answer = COALESCE(NULLIF(@answer, ''), answer),
    active = COALESCE(@active, active),
    updated_at = COALESCE(@updated_at, updated_at)
WHERE slug = @slug and lang = @lang;

-- name: DeleteCategory :exec
DELETE FROM category
WHERE slug = $1;
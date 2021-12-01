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
        parent_id,
        name,
        image,
        active,
        slug,
        lang,
        information
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: UpdateCategory :exec
UPDATE category
SET parent_id = COALESCE(NULLIF(@parent_id, ''), parent_id),
    image = COALESCE(NULLIF(@image, ''), image),
    active = COALESCE(@active, active),
    name = COALESCE(@name, name),
    information = COALESCE(@information, information),
    created_at = COALESCE(@created_at, created_at),
    updated_at = COALESCE(@updated_at, updated_at)
WHERE slug = @slug and lang = @lang;

-- name: DeleteCategory :exec
DELETE FROM category
WHERE slug = $1;

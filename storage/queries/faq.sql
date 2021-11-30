-- name: GetFaq :one
SELECT *
FROM faq
WHERE slug = @slug and lang = @lang
LIMIT 1;

-- name: GetAllFaq :many
SELECT *
FROM faq
WHERE lang = @lang
ORDER BY created_at desc
LIMIT @limit_ OFFSET @offset_;

-- name: CreateFaq :one
INSERT INTO faq (
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

-- name: UpdateFaq :exec
UPDATE faq
SET question = COALESCE(@question, question),
    answer = COALESCE(@answer, answer),
    active = COALESCE(@active, active),
    updated_at = COALESCE(@updated_at, updated_at)
WHERE slug = @slug and lang = @lang;

-- name: DeleteFaq :exec
DELETE FROM faq
WHERE id = $1;
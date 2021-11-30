-- name: GeeFaq :one
SELECT *
FROM faq
WHERE id = $1
LIMIT 1;

-- name: GetUserByPhoneNeFaq :one
SELECT * from faq WHERE phone_number = $1 LIMIT 1;

-- name: GeteFaq :many
SELECT *
FROM faq
ORDER BY created_at desc OFFSET $1
LIMIT $2;

-- name: CreateFaq :one
INSERT INTO faq (
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
RETURNING *;

-- name: UpdateFaq :exec
UPDATE faq
SET first_name = COALESCE(@first_name, first_name),
    last_name = COALESCE(@second_name, last_name),
    phone_number = COALESCE(NULLIF(@phone_number, ''), phone_number),
    is_verified = COALESCE(@is_verified, is_verified),
    long = COALESCE(@long, long),
    lat = COALESCE(@lat, lat),
    user_type = COALESCE(@user_type, user_type),
    updated_at = COALESCE(@updated_at, updated_at)
WHERE id = @id;

-- name: DeleteFaq :exec
DELETE FROM faq
WHERE id = $1;
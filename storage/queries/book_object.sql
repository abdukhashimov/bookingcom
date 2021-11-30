-- name: GetBookObject :one
SELECT *
FROM book_object
WHERE id = $1
LIMIT 1;

-- name: GetAllBookObject :many
SELECT *
FROM book_object
ORDER BY created_at desc OFFSET $1
LIMIT $2;

-- name: CreateBookObject :one
INSERT INTO book_object (
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

-- name: UpdateBookObject :exec
UPDATE book_object
SET first_name = COALESCE(@first_name, first_name),
    last_name = COALESCE(@second_name, last_name),
    phone_number = COALESCE(NULLIF(@phone_number, ''), phone_number),
    is_verified = COALESCE(@is_verified, is_verified),
    long = COALESCE(@long, long),
    lat = COALESCE(@lat, lat),
    user_type = COALESCE(@user_type, user_type),
    updated_at = COALESCE(@updated_at, updated_at)
WHERE id = @id;

-- name: DeleteBookObject :exec
DELETE FROM book_object
WHERE id = $1;